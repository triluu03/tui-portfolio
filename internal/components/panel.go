package components

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// titleOpen and titleClose bracket the title on the top border. The width
// available to the title is derived from them rather than hard-coded.
const (
	titleOpen  = " { "
	titleClose = " } "
)

// Panel styles, declared once and treated as immutable. lipgloss styles are
// value types, so sharing these is safe.
var (
	panelBorderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(style.ColorFaint)).
				Inline(true)
	panelSelectedBorderStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color(style.ColorBoatYellow1)).
					Inline(true)
	panelTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(style.ColorWarm)).
			Bold(true).
			Inline(true)
)

// Panel draws a square box of the given outer size with title embedded in the
// top border: ┌ { whoami } ─────────────┐. Every emitted line is exactly width
// visible columns; over-wide titles and body lines are truncated to fit.
func Panel(title string, width, height int, body string, selected bool) string {
	if width <= 0 || height <= 0 {
		return ""
	}

	borderStyle := panelBorderStyle
	if selected {
		borderStyle = panelSelectedBorderStyle
	}

	titleMaxWidth := width - lipgloss.Width("┌"+titleOpen+titleClose+"┐")
	text := ansi.Truncate(title, max(0, titleMaxWidth), "")
	dashes := max(0, titleMaxWidth-lipgloss.Width(text))
	innerWidth := width - lipgloss.Width("│  │")

	top := borderStyle.Render("┌") +
		panelTitleStyle.Render(titleOpen+text+titleClose) +
		borderStyle.Render(strings.Repeat("─", dashes)+"┐")

	lines := make([]string, max(0, height-2))
	copy(lines, strings.Split(body, "\n"))

	rows := make([]string, 0, height)
	rows = append(rows, top)
	left, right := borderStyle.Render("│ "), borderStyle.Render(" │")
	for _, line := range lines {
		rows = append(rows, left+padTo(line, innerWidth)+right)
	}
	rows = append(rows, borderStyle.Render("└"+strings.Repeat("─", width-2)+"┘"))

	return strings.Join(rows, "\n")
}

// padTo pads or truncates s to exactly width visible columns.
func padTo(s string, width int) string {
	if w := lipgloss.Width(s); w < width {
		return s + strings.Repeat(" ", width-w)
	}
	return ansi.Truncate(s, width, "")
}
