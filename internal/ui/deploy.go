package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dream11/odin/pkg/util"
	"strings"
)

func (s *ServiceDeployModel) Init() tea.Cmd {
	s.ServiceDisplayMeta.Progress = progress.New(progress.WithDefaultScaledGradient())
	s.ServiceDisplayMeta.Progress.PercentageStyle = ProgressBarStyle
	s.ServiceDisplayMeta.Progress.SetPercent(100)
	s.ServiceDisplayMeta.Cursor = 0
	return nil
}

func (s *ServiceDeployModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var updateCmds []tea.Cmd
	switch msg := msg.(type) {
	// Handle updates
	case ServiceDeployModel:
		s.ServiceDisplayMeta.Ready = true
		s.ServiceView = msg.ServiceView
		if len(s.ServiceView.ComponentsView) > len(s.ServiceDisplayMeta.ComponentDisplayMeta) {
			for i := len(s.ServiceDisplayMeta.ComponentDisplayMeta); i < len(s.ServiceView.ComponentsView); i++ {
				s.ServiceDisplayMeta.ComponentDisplayMeta = append(s.ServiceDisplayMeta.ComponentDisplayMeta, ComponentDisplayMeta{
					LogViewPort: viewport.Model{
						Width:  s.ServiceDisplayMeta.Width,
						Height: 10,
					},
				})
				s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner = spinner.New()
				s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner.Style = SpinnerStyle
				s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner.Spinner = spinner.Points
			}
		}
	// Handle key presses
	case tea.KeyMsg:
		anyToggled := false
		for _, component := range s.ServiceDisplayMeta.ComponentDisplayMeta {
			if component.Toggle {
				anyToggled = true
				break
			}
		}
		switch msg.String() {
		case "up":
			if s.ServiceDisplayMeta.Cursor > 0 && !anyToggled {
				s.ServiceDisplayMeta.Cursor--
			}
			break

		case "down":
			if s.ServiceDisplayMeta.Cursor < len(s.ServiceDisplayMeta.ComponentDisplayMeta)-1 && !anyToggled {
				s.ServiceDisplayMeta.Cursor++
			}
			break

		case "enter", " ":
			if s.ServiceDisplayMeta.Cursor < len(s.ServiceDisplayMeta.ComponentDisplayMeta) {
				s.ServiceDisplayMeta.ComponentDisplayMeta[s.ServiceDisplayMeta.Cursor].Toggle =
					!s.ServiceDisplayMeta.ComponentDisplayMeta[s.ServiceDisplayMeta.Cursor].Toggle
			}
			break

		case "q":
			return s, tea.Quit
		}

	// Handle window resizes
	case tea.WindowSizeMsg:
		s.ServiceDisplayMeta.Height = msg.Height
		s.ServiceDisplayMeta.Width = msg.Width
		if s.ServiceDisplayMeta.Ready {
			for i := range s.ServiceDisplayMeta.ComponentDisplayMeta {
				s.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.Width = msg.Width
				s.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.Height =
					max(util.GetAvailableViewPortHeight(msg.Height, lipgloss.Height("text"), len(s.ServiceDisplayMeta.ComponentDisplayMeta)), 10)
			}
		}
	case spinner.TickMsg:
		spinnerUpdateCmds := s.updateSpinners(msg)
		return s, tea.Batch(spinnerUpdateCmds...)
	}

	// Handle keyboard and mouse events in the viewport
	updateCmds = append(updateCmds, s.tickSpinners()...)
	updateCmds = append(updateCmds, s.updateViewPort(msg)...)

	return s, tea.Batch(updateCmds...)
}

func (s *ServiceDeployModel) View() string {
	if !s.ServiceDisplayMeta.Ready {
		return H1Style.Render("Initializing service deployment...")
	}
	var builder strings.Builder

	// Build Service View
	serviceHeader := util.GetHeaderText(s.ServiceView.Name, s.ServiceView.Action, s.ServiceView.Status, "Service")
	builder.WriteString(fmt.Sprintf("%s\n", H1Style.Render(serviceHeader)))
	s.ServiceDisplayMeta.Progress.Width = lipgloss.Width(serviceHeader) + 6 // to accommodate the percentage text
	builder.WriteString(fmt.Sprintf("%s\n", s.ServiceDisplayMeta.Progress.ViewAs(100.0)))
	builder.WriteString(fmt.Sprintf("Trace Id: %s\n", TextStyle.Render(s.ServiceView.TraceId)))

	for i, componentView := range s.ServiceView.ComponentsView {
		componentHeader := util.GetHeaderText(componentView.Name, componentView.Action, componentView.Status, "Component")
		var componentHeaderText string
		if i == s.ServiceDisplayMeta.Cursor {
			componentHeaderText = SelectedStyle(H2Style).Render(componentHeader)
		} else {
			componentHeaderText = H2Style.Render(componentHeader)
		}
		if s.ServiceView.ComponentsView[i].Status != "IN_PROGRESS" {
			s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner.Spinner = spinner.Pulse
		}
		builder.WriteString(fmt.Sprintf("%s %s\n", componentHeaderText, s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner.View()))
		if s.ServiceDisplayMeta.ComponentDisplayMeta[i].Toggle {
			// Render logs
			logsText := strings.Split(componentView.Content, "\\n")
			s.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.SetContent(InfoStyle.Render(strings.Join(logsText, "\n")))
			builder.WriteString(fmt.Sprintf("%s \n", s.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.View()))
		}
	}

	// Add Footer with operating instructions
	builder.WriteString("\n\n")
	builder.WriteString(FooterStyle.Render("Use ↑ and ↓ to navigate components, Enter to toggle logs, q to quit"))

	return builder.String()
}

func (s *ServiceDeployModel) updateViewPort(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	for i := range s.ServiceDisplayMeta.ComponentDisplayMeta {
		s.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort, cmd =
			s.ServiceDisplayMeta.ComponentDisplayMeta[i].LogViewPort.Update(msg)
		cmds = append(cmds, cmd)
	}
	return cmds
}

func (s *ServiceDeployModel) updateSpinners(msg tea.Msg) []tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	for i := range s.ServiceDisplayMeta.ComponentDisplayMeta {
		s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner, cmd =
			s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner.Update(msg)
		cmds = append(cmds, cmd)
	}
	return cmds
}

func (s *ServiceDeployModel) tickSpinners() []tea.Cmd {
	var cmds []tea.Cmd
	for i := range s.ServiceDisplayMeta.ComponentDisplayMeta {
		cmds = append(cmds, s.ServiceDisplayMeta.ComponentDisplayMeta[i].Spinner.Tick)
	}
	return cmds
}
