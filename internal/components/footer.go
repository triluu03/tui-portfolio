package components

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// Footer renders the bottom breadcrumb bar showing the active page's title and
// source filename. The returned string is a single line exactly width columns
// wide, or empty when width <= 0.
func Footer(width int, title string) string {
	if width <= 0 {
		return ""
	}

	footerTitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(style.ColorAccent)).
		Bold(true).
		Inline(true).
		Render(title)

	hints := lipgloss.NewStyle().
		Foreground(lipgloss.Color(style.ColorDim)).
		Inline(true).
		Render("1-4 switch · ↑↓ move · enter open · q quit")

	gaps := width - lipgloss.Width(footerTitle) - lipgloss.Width(hints)
	if gaps < 1 {
		return footerTitle
	}

	return footerTitle + strings.Repeat(" ", gaps) + hints

}
