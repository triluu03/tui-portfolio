package components

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// Header renders the top navigation bar with one tab per entry in tabs. The
// tab at index active is highlighted. The returned string is a single line
// exactly width columns wide, or empty when width <= 0.
func Header(width int, active int, tabs []string) string {
	if width <= 0 || len(tabs) == 0 {
		return ""
	}

	tabBaseStyle := lipgloss.NewStyle().
		Height(1).
		AlignVertical(lipgloss.Center).
		Padding(0, 1)
	activeStyle := tabBaseStyle.
		Foreground(lipgloss.Color(style.ColorBackground)).
		Background(lipgloss.Color(style.ColorAccent)).
		Bold(true)
	inactiveStyle := tabBaseStyle.
		Foreground(lipgloss.Color(style.ColorDim))

	parts := make([]string, len(tabs))
	for i, t := range tabs {
		label := fmt.Sprintf("%d %s", i+1, t)
		if i == active {
			parts[i] = activeStyle.Render(label)
		} else {
			parts[i] = inactiveStyle.Render(label)
		}
	}

	return lipgloss.NewStyle().
		Width(width).
		Render(lipgloss.JoinHorizontal(lipgloss.Center, parts...))
}
