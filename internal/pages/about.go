package pages

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// About is the about page.
type About struct{}

// Init is a no-op placeholder.
func (a About) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (a About) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return a, nil }

// View renders the about page with the whoami and sysinfo panels.
func (a About) View() tea.View { return tea.NewView(render()) }

const (
	role    = "Senior Data Scientist @ Sellfote Solutions Oy"
	tagline = "Just a normal tech guy."
	// nameArt is the 5-row block rendering of the literal "[TRI LUU]".
	// Every line is exactly 49 visible columns wide.
	nameArt = `█▀▀ █████ ████  █████       █     █   █ █   █ ▀▀█
█     █   █   █   █         █     █   █ █   █   █
█     █   ████    █         █     █   █ █   █   █
█     █   █  █    █         █     █   █ █   █   █
█▄▄   █   █   █ █████       █████  ███   ███  ▄▄█`
)

// nameArtLines is nameArt split into its rows, computed once.
var nameArtLines = strings.Split(nameArt, "\n")

// sysInfoRows are the key/value rows shown in the sysinfo panel.
var sysInfoRows = [][2]string{
	{"based", "Espoo, Finland"},
	{"from", "Hanoi, Vietnam"},
	{"timezone", "GMT+3"},
	{"building", "a technical stuff"},
	{"reading", "an intriguing book"},
	{"watching", "a marvelous movie"},
}

// swatchColors are the palette colours shown at the bottom of the sysinfo panel.
var swatchColors = []string{
	style.ColorWarm,
	style.ColorAccent,
	style.ColorForeground,
	style.ColorDim,
	style.ColorFaint,
}

// Panel styles, declared once and treated as immutable. lipgloss styles are
// value types, so sharing these is safe.
var (
	accentStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorAccent))
	dimStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorDim))
	warmStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorWarm))
)

// render builds the two-column about page padded to the pinned frame size.
func render() string {
	whoami := whoamiPanel(style.AboutWhoAmIWidth, style.ContentHeight)
	sysinfo := sysInfoPanel(style.AboutSysInfoWidth, style.ContentHeight, sysInfoRows)
	row := lipgloss.JoinHorizontal(lipgloss.Top, whoami, strings.Repeat(" ", style.ColumnGap), sysinfo)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(row)
}

// whoamiPanel renders the whoami panel.
func whoamiPanel(width, height int) string {
	promptLine := accentStyle.Render("$") + dimStyle.Render(" whoami")
	roleLine := warmStyle.Render(role)
	taglineLine := dimStyle.Render(tagline)

	body := append([]string{"", promptLine, ""}, nameArtLines...)
	body = append(body, "", roleLine, taglineLine)

	return components.Panel("whoami", width, height, strings.Join(body, "\n"))
}

// sysInfoPanel renders the sysinfo panel: key/value rows at the top and a palette
// swatch row pinned to the bottom.
func sysInfoPanel(width, height int, rows [][2]string) string {
	body := make([]string, 0, height-2)
	body = append(body, "")
	for _, r := range rows {
		body = append(body, dimStyle.Render(r[0]), r[1], "")
	}
	if n := height - 3 - len(body); n > 0 {
		body = append(body, make([]string, n)...)
	}
	body = append(body, swatchRow(swatchColors))
	return components.Panel("sysinfo", width, height, strings.Join(body, "\n"))
}

// swatchRow renders each colour as a two-cell block joined by a space.
func swatchRow(colors []string) string {
	blocks := make([]string, len(colors))
	for i, c := range colors {
		blocks[i] = lipgloss.NewStyle().Foreground(lipgloss.Color(c)).Render("███")
	}
	return strings.Join(blocks, " ")
}
