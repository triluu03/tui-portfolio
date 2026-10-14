package pages

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
)

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

var nSysInfoRows = len(sysInfoRows)

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

// About is the about page.
type About struct {
	selectedSysInfo int
	focus           *bool
}

// Init is a no-op placeholder.
func (a About) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (a About) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && (*a.focus) {
		switch key.String() {
		case "up", "k":
			a.selectedSysInfo = (a.selectedSysInfo + nSysInfoRows - 1) % nSysInfoRows
		case "down", "j":
			a.selectedSysInfo = (a.selectedSysInfo + 1) % nSysInfoRows
		}
	}
	return a, nil
}

// View renders the about page with the whoami and sysinfo panels.
func (a About) View() tea.View { return tea.NewView(a.renderAbout()) }

// renderAbout builds the two-column about page padded to the pinned frame size.
func (a About) renderAbout() string {
	whoami := whoamiPanel(style.AboutWhoAmIWidth, style.ContentHeight, !*a.focus)
	sysinfo := sysInfoPanel(
		style.AboutSysInfoWidth,
		style.ContentHeight,
		sysInfoRows,
		a.selectedSysInfo,
		*a.focus,
	)

	row := lipgloss.JoinHorizontal(
		lipgloss.Top,
		whoami,
		strings.Repeat(" ", style.ColumnGap),
		sysinfo,
	)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(row)
}

// whoamiPanel renders the whoami panel.
func whoamiPanel(width, height int, focused bool) string {
	promptLine := accentStyle.Render("$") + dimStyle.Render(" whoami")
	roleLine := warmStyle.Render(role)
	taglineLine := dimStyle.Render(tagline)

	body := append([]string{"", promptLine, ""}, nameArtLines...)
	body = append(body, "", roleLine, taglineLine)

	return components.Panel("whoami", width, height, strings.Join(body, "\n"), focused)
}

// sysInfoPanel renders the sysinfo panel: key/value rows at the top and a palette
// swatch row pinned to the bottom.
func sysInfoPanel(width, height int, rows [][2]string, selected int, focused bool) string {
	body := make([]string, 0, height-2)
	kv := make([]keyValueRow, len(rows))
	for i, r := range rows {
		kv[i] = keyValueRow{label: r[0], value: r[1], decoration: " <"}
	}
	body = append(body, keyValueBody(kv, selected)...)
	if n := height - 3 - len(body); n > 0 {
		body = append(body, make([]string, n)...)
	}
	body = append(body, swatchRow(swatchColors))
	return components.Panel("sysinfo", width, height, strings.Join(body, "\n"), focused)
}

// swatchRow renders each colour as a two-cell block joined by a space.
func swatchRow(colors []string) string {
	blocks := make([]string, len(colors))
	for i, c := range colors {
		blocks[i] = lipgloss.NewStyle().Foreground(lipgloss.Color(c)).Render("███")
	}
	return strings.Join(blocks, " ")
}
