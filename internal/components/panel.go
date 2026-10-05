package components

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
)

// The application frame is pinned to FrameWidth x FrameHeight. ContentHeight is
// FrameHeight minus the header row and footer row.
const (
	FrameWidth    = 120
	FrameHeight   = 36
	ContentHeight = FrameHeight - 2
)

// Panel draws a square box of the given outer size with title embedded in the
// top border: ┌ | whoami | ─────────────┐. Every emitted line is exactly width
// visible columns; over-wide titles and body lines are truncated to fit.
func Panel(title string, width, height int, body string) string {
	if width <= 0 || height <= 0 {
		return ""
	}

	faint := lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorFaint)).
		Inline(true)
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(ColorDim)).
		Bold(true).
		Inline(true)

	title = ansi.Truncate(title, max(0, width-8), "")
	dashes := max(0, width-8-lipgloss.Width(title))
	top := faint.Render("┌ | ") +
		titleStyle.Render(title) +
		faint.Render(" | "+strings.Repeat("─", dashes)+"┐")

	bodyLines := strings.Split(body, "\n")
	bodyRows := height - 2
	rows := make([]string, 0, height)
	rows = append(rows, top)
	for i := range bodyRows {
		var line string
		if i < len(bodyLines) {
			line = bodyLines[i]
		}
		rows = append(rows, faint.Render("│ ")+padTo(line, width-4)+faint.Render(" │"))
	}
	rows = append(rows, faint.Render("└"+strings.Repeat("─", width-2)+"┘"))

	return strings.Join(rows, "\n")
}

// padTo pads or truncates s to exactly width visible columns.
func padTo(s string, width int) string {
	if w := lipgloss.Width(s); w < width {
		return s + strings.Repeat(" ", width-w)
	}
	return ansi.Truncate(s, width, "")
}
