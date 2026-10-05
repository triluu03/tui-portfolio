package pages

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
)

// About is the about page.
type About struct{}

// Init is a no-op placeholder.
func (a About) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (a About) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return a, nil }

const (
	// role is the placeholder role line.
	role = "Senior Data Scientist · Espoo, Finland"
	// tagline is the placeholder tagline.
	tagline = "Just a normal tech guy."
	// nameArt is the 5-row block rendering of the literal "[TRI LUU]".
	// Every line is exactly 49 visible columns wide.
	nameArt = `█▀▀ █████ ████  █████       █     █   █ █   █ ▀▀█
█     █   █   █   █         █     █   █ █   █   █
█     █   ████    █         █     █   █ █   █   █
█     █   █  █    █         █     █   █ █   █   █
█▄▄   █   █   █ █████       █████  ███   ███  ▄▄█`
)

// View renders the about page whoami panel.
func (a About) View() tea.View { return tea.NewView(a.render()) }

// render builds the whoami panel padded to the pinned frame size.
func (a About) render() string {
	accent := lipgloss.NewStyle().Foreground(lipgloss.Color(components.ColorAccent))
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color(components.ColorDim))

	promptLine := accent.Render("$") + dim.Render(" whoami")
	roleLine := lipgloss.NewStyle().
		Foreground(lipgloss.Color(components.ColorWarm)).
		Render(role)
	taglineLine := dim.Render(tagline)
	nameArtLine := strings.Split(nameArt, "\n")

	body := append([]string{"", promptLine, ""}, nameArtLine...)
	body = append(body, "", roleLine, taglineLine)

	panel := components.Panel("whoami", components.FrameWidth-4, components.ContentHeight, strings.Join(body, "\n"))

	// A FrameWidth-4 panel plus 2 columns of margin on each side = FrameWidth.
	return "  " + strings.ReplaceAll(panel, "\n", "  \n  ") + "  "
}
