package pages

import (
	"fmt"
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

const (
	// role is the placeholder role line.
	role = "Senior Data Scientist @ Sellfote Solutions Oy"
	// tagline is the placeholder tagline.
	tagline = "Just a normal tech guy."
	// nameArt is the 5-row block rendering of the literal "[TRI LUU]".
	// Every line is exactly 49 visible columns wide.
	nameArt = `█▀▀ █████ ████  █████       █     █   █ █   █ ▀▀█
█     █   █   █   █         █     █   █ █   █   █
█     █   ████    █         █     █   █ █   █   █
█     █   █  █    █         █     █   █ █   █   █
█▄▄   █   █   █ █████       █████  ███   ███  ▄▄█`

	// padX is the horizontal padding on each side of the content area.
	padX = 2
	// whoamiWidth is the outer width of the whoami panel.
	whoamiWidth = 77
	// columnGap is the horizontal gap between the two panels.
	columnGap = 2
	// sysInfoWidth is the outer width of the sysinfo panel.
	sysInfoWidth = 37
	// keyWidth is the width reserved for sysinfo row keys.
	keyWidth = 9
)

// sysInfoRows are the key/value rows shown in the sysinfo panel.
var sysInfoRows = [][2]string{
	{"", ""},
	{"based", "Espoo, Finland"},
	{"timezone", "GMT+3"},
	{"building", "tui-portfolio"},
	{"reading", "The Trial - Kafka"},
}

// swatchColors are the palette colours shown at the bottom of the sysinfo panel.
var swatchColors = []string{
	style.ColorWarm,
	style.ColorAccent,
	style.ColorForeground,
	style.ColorDim,
	style.ColorFaint,
}

// View renders the about page with the whoami and sysinfo panels.
func (a About) View() tea.View { return tea.NewView(a.render()) }

// render builds the two-column about page padded to the pinned frame size.
func (a About) render() string {
	whoami := whoamiPanel(whoamiWidth, components.ContentHeight)
	sysinfo := sysInfoPanel(sysInfoWidth, components.ContentHeight, sysInfoRows)
	row := lipgloss.JoinHorizontal(lipgloss.Top, whoami, strings.Repeat(" ", columnGap), sysinfo)
	pad := strings.Repeat(" ", padX)
	return pad + strings.ReplaceAll(row, "\n", pad+"\n"+pad) + pad
}

// whoamiPanel renders the whoami panel.
func whoamiPanel(width, height int) string {
	accent := lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorAccent))
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorDim))
	warm := lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorWarm))

	promptLine := accent.Render("$") + dim.Render(" whoami")
	roleLine := warm.Render(role)
	taglineLine := dim.Render(tagline)

	body := append([]string{"", promptLine, ""}, strings.Split(nameArt, "\n")...)
	body = append(body, "", roleLine, taglineLine)

	return components.Panel("whoami", width, height, strings.Join(body, "\n"))
}

// sysInfoPanel renders the sysinfo panel: key/value rows at the top and a palette
// swatch row pinned to the bottom.
func sysInfoPanel(width, height int, rows [][2]string) string {
	key := lipgloss.NewStyle().Foreground(lipgloss.Color(style.ColorDim))

	body := make([]string, 0, height-2)
	for _, r := range rows {
		body = append(body, key.Render(fmt.Sprintf("%-*s", keyWidth, r[0]))+r[1])
	}
	for len(body) < height-3 {
		body = append(body, "")
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
