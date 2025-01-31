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
		&ServiceView{
			ComponentsView: append(
				[]ComponentView{
					{
						Toggle: false,
					},
				},
			),
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

func (s *ServiceView) Init() tea.Cmd {
	return nil
}

func (s *ServiceView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Handle updates
	case ServiceView:
		s.Header = msg.Header
		s.Status = msg.Status
		s.Cursor = msg.Cursor
		prevCollapsed := make([]bool, len(s.ComponentsView))
		for i := range s.ComponentsView {
			if i < len(prevCollapsed) {
				prevCollapsed[i] = s.ComponentsView[i].Toggle
			}
		}

		for i := range s.ComponentsView {
			if i < len(prevCollapsed) {
				s.ComponentsView[i].Toggle = prevCollapsed[i]
				s.ComponentsView[i].Header = msg.ComponentsView[i].Header
				s.ComponentsView[i].Status = msg.ComponentsView[i].Status
				s.ComponentsView[i].LogView.Content = msg.ComponentsView[i].LogView.Content
			}
		}
	// Handle key presses
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if s.Cursor > 0 {
				s.Cursor--
			}
			break

		case "down":
			if s.Cursor < len(s.ComponentsView)-1 {
				s.Cursor++
			}
			break

		case "enter", " ":
			s.ComponentsView[s.Cursor].Toggle = !s.ComponentsView[s.Cursor].Toggle
			break

		case "q":
			return s, tea.Quit
		}

	// Handle window resizes
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(s.Header.Text)
		if !s.Ready {
			s.ComponentsView[s.Cursor].LogView.LogViewPort = viewport.New(msg.Width, msg.Height-headerHeight)
			s.ComponentsView[s.Cursor].LogView.LogViewPort.MouseWheelDelta = 1
			s.ComponentsView[s.Cursor].LogView.LogViewPort.YPosition = headerHeight
			s.Ready = true
		} else {
			s.ComponentsView[s.Cursor].LogView.LogViewPort.Width = msg.Width - 20
			s.ComponentsView[s.Cursor].LogView.LogViewPort.Height = msg.Height - headerHeight - 20
		}
	}

	// Handle keyboard and mouse events in the viewport
	s.ComponentsView[s.Cursor].LogView.LogViewPort.SetContent(s.ComponentsView[s.Cursor].LogView.Content)
	vpcmd := s.updateViewPort(msg)

	return s, tea.Batch(vpcmd...)
}

func (s *ServiceView) updateViewPort(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	for i := range s.ComponentsView {
		s.ComponentsView[i].LogView.LogViewPort, cmd = s.ComponentsView[i].LogView.LogViewPort.Update(msg)
		cmds = append(cmds, cmd)
	}
	return cmds
}

func (s *ServiceView) View() string {
	if !s.Ready {
		return "\n  Initializing..."
	}
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s\n", s.Header.Text))

	for i, componentView := range s.ComponentsView {
		cursor := " "
		if i == s.Cursor {
			if componentView.Toggle {
				cursor = ">"
			} else {
				cursor = "V"
			}
		}

		if !componentView.Toggle {
			builder.WriteString(fmt.Sprintf("%s %s \n", cursor, componentView.Header.Text))
		} else {
			builder.WriteString(fmt.Sprintf("%s\n%s\n%s\n", cursor, componentView.Header.Text, componentView.LogView.LogViewPort.View()))
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
		}
		program.Send(GetServiceView(response))
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
