package ui

import "github.com/charmbracelet/lipgloss"

var H1Style = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#BC88D3")).
		Margin(1).
		Align(lipgloss.Left)
}()

var H2Style = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8D2F2")).
		Margin(1, 1, 0, 5).
		Align(lipgloss.Left)
}()

var TextStyle = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8D2F2")).
		Italic(true).
		Align(lipgloss.Left)
}()

var FooterStyle = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8D2F2")).
		Bold(false).
		Italic(true).
		Faint(true).
		Align(lipgloss.Left)
}()

var InfoStyle = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E8D2F2")).
		MarginLeft(5).
		BorderLeft(true).
		BorderStyle(lipgloss.RoundedBorder()).
		PaddingLeft(2).
		Align(lipgloss.Left)
}()

var SelectedStyle = func(style lipgloss.Style) lipgloss.Style {
	return style.
		Underline(true).
		Bold(true)
}

var SpinnerStyle = func() lipgloss.Style {
	return lipgloss.NewStyle().
		PaddingBottom(1).
		Foreground(lipgloss.Color("#D9D9D9"))
}()

var ProgressBarStyle = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#E88D4C"))
}()
