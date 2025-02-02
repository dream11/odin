package ui

import "github.com/charmbracelet/lipgloss"

var H1Style = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a556f4")).
		Border(lipgloss.RoundedBorder(), true, false).
		Padding(1, 2).
		Align(lipgloss.Center)
}()

var H2Style = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7d56f4")).
		MarginLeft(5).
		Border(lipgloss.NormalBorder(), true, false).
		Align(lipgloss.Center)
}()

var InfoStyle = func() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fafafa")).
		MarginLeft(10).
		Padding(1, 2).
		Align(lipgloss.Left)
}()

var SelectedStyle = func(style lipgloss.Style) lipgloss.Style {
	return style.
		Background(style.GetForeground()).
		Foreground(lipgloss.Color("#fafafa")).
		Bold(true)
}

//var GreenBorder = func(style lipgloss.Style) lipgloss.Style {
//	return style.
//		Border(lipgloss.DoubleBorder()).
//		BorderForeground(lipgloss.Color("#00ff00"))
//}
