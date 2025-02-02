package deploy

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dream11/odin/pkg/constant"
	"github.com/dream11/odin/pkg/ui"
	"github.com/dream11/odin/pkg/util"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/dream11/odin/internal/service"
	"github.com/dream11/odin/pkg/config"
	serviceDto "github.com/dream11/odin/proto/gen/go/dream11/od/dto/v1"
	serviceProto "github.com/dream11/odin/proto/gen/go/dream11/od/service/v1"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var env string
var definitionFile string
var provisioningFile string
var serviceName string
var serviceVersion string
var serviceClient = service.Service{}
var labels string
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Deploy service",
	Args: func(cmd *cobra.Command, args []string) error {
		return cobra.NoArgs(cmd, args)
	},
	Long: "Deploy service using files or service name",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd)
	},
}

func init() {
	serviceCmd.Flags().StringVar(&env, "env", "", "environment for deploying the service")
	serviceCmd.Flags().StringVar(&definitionFile, "file", "", "path to the service definition file")
	serviceCmd.Flags().StringVar(&provisioningFile, "provisioning", "", "path to the provisioning file")
	serviceCmd.Flags().StringVar(&serviceName, "name", "", "released service name")
	serviceCmd.Flags().StringVar(&serviceVersion, "version", "", "released service version")
	serviceCmd.Flags().StringVar(&labels, "labels", "", "comma separated labels for the service version ex key1=value1,key2=value2")

	deployCmd.AddCommand(serviceCmd)
}

func execute(cmd *cobra.Command) {
	env = config.EnsureEnvPresent(env)
	// Add program in context
	ctx := cmd.Context()
	program := tea.NewProgram(
		&Model{
			ServiceView: ServiceView{
				Name: "Initiating service deployment",
			},
		},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support, so we can track the mouse wheel
	)

	go func() {
		if (serviceName == "" && serviceVersion == "" && labels == "") && (definitionFile != "" && provisioningFile != "") {
			deployUsingFiles(ctx, program)
		} else if (serviceName != "" && serviceVersion != "" && labels == "") && (definitionFile == "" && provisioningFile == "") {
			deployUsingServiceNameAndVersion(ctx)
		} else if (serviceName != "" && labels != "" && serviceVersion == "") && (definitionFile == "" && provisioningFile == "") {
			if err := validateLabels(labels); err != nil {
				log.Fatal("Invalid labels format: ", err)
			}
			deployUsingServiceNameAndLabels(ctx)
		} else {
			log.Fatal("Invalid combination of flags. Use either (service name and version) or (service name and labels) or (definitionFile and provisioningFile).")
		}
	}()

	if _, err := program.Run(); err != nil {
		os.Exit(1)
	}
}

func (m *Model) Init() tea.Cmd {
	m.ServiceDisplayMeta.Cursor = 0
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle updates
	case Model:
		m.ServiceView = msg.ServiceView
		if len(m.ServiceView.ComponentsView) > len(m.ServiceDisplayMeta.ComponentDisplayMeta) {
			for i := len(m.ServiceDisplayMeta.ComponentDisplayMeta); i < len(m.ServiceView.ComponentsView); i++ {
				m.ServiceDisplayMeta.ComponentDisplayMeta = append(m.ServiceDisplayMeta.ComponentDisplayMeta, ComponentDisplayMeta{
					LogViewPort: viewport.Model{
						Width:  m.ServiceDisplayMeta.Width,
						Height: 10,
					},
				})
			}
		}
	// Handle key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.ServiceDisplayMeta.Cursor > 0 {
				m.ServiceDisplayMeta.Cursor--
			}
			break

		case "down":
			if m.ServiceDisplayMeta.Cursor < len(m.ServiceDisplayMeta.ComponentDisplayMeta)-1 {
				m.ServiceDisplayMeta.Cursor++
			}
			break

		case "enter", " ":
			if m.ServiceDisplayMeta.Cursor < len(m.ServiceDisplayMeta.ComponentDisplayMeta) {
				m.ServiceDisplayMeta.ComponentDisplayMeta[m.ServiceDisplayMeta.Cursor].Toggle =
					!m.ServiceDisplayMeta.ComponentDisplayMeta[m.ServiceDisplayMeta.Cursor].Toggle
			}
			break

		case "q":
			return m, tea.Quit
		}

	// Handle window resizes
	case tea.WindowSizeMsg:
		m.ServiceDisplayMeta.Height = msg.Height
		m.ServiceDisplayMeta.Width = msg.Width
		if !m.ServiceDisplayMeta.Ready {
			m.ServiceDisplayMeta.Ready = true
		} else {
			for i := range m.ServiceDisplayMeta.ComponentDisplayMeta {
				m.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.Width = msg.Width
				m.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.Height =
					min(msg.Height-len(m.ServiceDisplayMeta.ComponentDisplayMeta)*(lipgloss.Height(m.ServiceView.ComponentsView[i].Name)),
						10)
			}
		}
	}
	// Handle keyboard and mouse events in the viewport
	viewPortCmd := m.updateViewPort(msg)

	return m, tea.Batch(viewPortCmd...)
}

func (m *Model) updateViewPort(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	for i := range m.ServiceDisplayMeta.ComponentDisplayMeta {
		m.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort, cmd =
			m.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.Update(msg)
		cmds = append(cmds, cmd)
	}
	return cmds
}

func (m *Model) View() string {
	if !m.ServiceDisplayMeta.Ready {
		return "\n  Initializing..."
	}
	var builder strings.Builder

	// Build Service View
	serviceHeader := util.GetHeaderText(m.ServiceView.Name, m.ServiceView.Action, m.ServiceView.Status, "Service")
	builder.WriteString(ui.H1Style.Render(serviceHeader))

	for i, componentView := range m.ServiceView.ComponentsView {
		componentHeader := util.GetHeaderText(componentView.Name, componentView.Action, componentView.Status, "Component")
		var componentHeaderText string
		if i == m.ServiceDisplayMeta.Cursor {
			componentHeaderText = ui.SelectedStyle(ui.H2Style).Render(componentHeader)
		} else {
			componentHeaderText = ui.H2Style.Render(componentHeader)
		}

		if !m.ServiceDisplayMeta.ComponentDisplayMeta[i].Toggle {
			builder.WriteString(fmt.Sprintf("%s \n", componentHeaderText))
		} else {
			builder.WriteString(fmt.Sprintf("%s \n", componentHeaderText))
			// Render logs
			logsText := strings.Split(componentView.Content, "\\n")
			m.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.SetContent(ui.InfoStyle.Render(strings.Join(logsText, "\n")))
			builder.WriteString(fmt.Sprintf("%s \n", m.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.View()))
		}
	}

	return builder.String()
}

func deployUsingFiles(ctx context.Context, program *tea.Program) {
	definitionData, err := os.ReadFile(definitionFile)
	if err != nil {
		log.Fatal("Error while reading definition file ", err)
	}
	var definitionProto serviceDto.ServiceDefinition
	if err := json.Unmarshal(definitionData, &definitionProto); err != nil {
		log.Fatalf("Error unmarshalling definition file: %v", err)
	}

	provisioningData, err := os.ReadFile(provisioningFile)
	if err != nil {
		log.Fatal("Error while reading provisioning file ", err)
	}
	var compProvConfigs []*serviceDto.ComponentProvisioningConfig
	if err := json.Unmarshal(provisioningData, &compProvConfigs); err != nil {
		log.Fatalf("Error unmarshalling provisioning file: %v", err)
	}
	provisioningProto := &serviceDto.ProvisioningConfig{
		ComponentProvisioningConfig: compProvConfigs,
	}

	stream, err := serviceClient.DeployService(&ctx, &serviceProto.DeployServiceRequest{
		EnvName:            env,
		ServiceDefinition:  &definitionProto,
		ProvisioningConfig: provisioningProto,
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}
	for {
		response, err := stream.Recv()
		if err != nil {
			if errors.Is(err, context.Canceled) || err == io.EOF {
				break
			}
			log.Errorf("TraceID: %s", ctx.Value(constant.TraceIDKey))
			log.Fatal("Failed to deploy service ", err)
			program.Quit()
		}
		program.Send(GetServiceDeployModel(response))
	}
}

func deployUsingServiceNameAndVersion(ctx context.Context) {
	log.Info("deploying service :", serviceName, ":", serviceVersion, " in env :", env)
	err := serviceClient.DeployReleasedService(&ctx, &serviceProto.DeployReleasedServiceRequest{
		EnvName: env,
		ServiceIdentifier: &serviceProto.ServiceIdentifier{
			ServiceName:    serviceName,
			ServiceVersion: serviceVersion,
		},
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}
}

func deployUsingServiceNameAndLabels(ctx context.Context) {
	log.Info("deploying service :", serviceName, " with labels: ", labels, " in env :", env)
	err := serviceClient.DeployReleasedService(&ctx, &serviceProto.DeployReleasedServiceRequest{
		EnvName: env,
		ServiceIdentifier: &serviceProto.ServiceIdentifier{
			ServiceName: serviceName,
			Tags:        labels,
		},
	})

	if err != nil {
		log.Fatal("Failed to deploy service ", err)
	}
}

func validateLabels(labels string) error {
	labelPattern := `^(\w+=\w+)(,\w+=\w+)*$`
	matched, err := regexp.MatchString(labelPattern, labels)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("labels must be in format key1=value1,key2=value2")
	}
	return nil
}
