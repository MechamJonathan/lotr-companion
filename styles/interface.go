package styles

import "github.com/charmbracelet/lipgloss"

var PromptStyle = lipgloss.NewStyle().Foreground(Orange).Bold(true).Align(lipgloss.Center)
var ArrowSymbol = lipgloss.NewStyle().Foreground(Red).Bold(true).Align(lipgloss.Center)
var ErrorMessage = lipgloss.NewStyle().Foreground(Yellow).Faint(true).Align(lipgloss.Center).PaddingBottom(1)
