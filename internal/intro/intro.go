package intro

import (
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/app"
	"github.com/triluu03/tui-portfolio/internal/style"
)

const (
	frameInterval = 50 * time.Millisecond
	totalFrames   = 30
	barWidth      = 30

	filledRune = "█"
	emptyRune  = "░"
)

type tickMsg struct{}

// Model animates a centered loading progress bar before handing off to the
// real application.
type Model struct {
	frame      int
	termWidth  int
	termHeight int
}

// New returns the intro model with its frame counter reset.
func New() Model {
	return Model{}
}

// Init schedules the first animation frame.
func (m Model) Init() tea.Cmd {
	return tea.Every(frameInterval, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

// Update advances the animation and hands off to the app once complete.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		return m, nil
	case tickMsg:
		m.frame++
		if m.frame >= totalFrames {
			next := app.New()
			return next, tea.Batch(
				next.Init(),
				func() tea.Msg { return tea.WindowSizeMsg{Width: m.termWidth, Height: m.termHeight} },
			)
		}
		return m, tea.Every(frameInterval, func(time.Time) tea.Msg {
			return tickMsg{}
		})
	default:
		return m, nil
	}
}

// View renders the centered loading progress bar.
func (m Model) View() tea.View {
	pct := m.frame * barWidth / totalFrames
	bar := strings.Repeat(filledRune, pct) + strings.Repeat(emptyRune, barWidth-pct)
	bar = lipgloss.NewStyle().Render(bar)

	label := lipgloss.NewStyle().
		Width(barWidth).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color(style.ColorDim)).
		Render("Loading portfolio...")

	content := lipgloss.JoinVertical(lipgloss.Left, bar, label)

	v := tea.NewView(lipgloss.Place(m.termWidth, m.termHeight, lipgloss.Center, lipgloss.Center, content))
	v.AltScreen = true
	v.BackgroundColor = lipgloss.Color(style.ColorBackground)
	v.ForegroundColor = lipgloss.Color(style.ColorForeground)
	return v
}
