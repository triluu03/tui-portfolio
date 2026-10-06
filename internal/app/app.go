package app

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/pages"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// The root application Bubble Tea's Model.
type Model struct {
	pages      []tea.Model
	current    int
	termWidth  int
	termHeight int
}

// Return a new model with the pages wired up into the states.
func New() Model {
	var m Model
	for _, d := range pages.Pages {
		m.pages = append(m.pages, d.Model)
	}
	return m
}

// Init returns the initial commands from every page.
func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, p := range m.pages {
		cmds = append(cmds, p.Init())
	}
	return tea.Batch(cmds...)
}

// Update handles global key bindings, terminal resizes, and delegates other
// messages to the active page.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termWidth = msg.Width
		m.termHeight = msg.Height
		return m, nil
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1", "2", "3", "4":
			m.current = int(msg.String()[0] - '1')
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.pages[m.current], cmd = m.pages[m.current].Update(msg)
	return m, cmd
}

// View renders the pinned FrameWidth x FrameHeight frame centered in the terminal.
// When the terminal is smaller than the frame, it shows a hint instead.
func (m Model) View() tea.View {
	tabs := make([]string, len(pages.Pages))
	for i, d := range pages.Pages {
		tabs[i] = d.Title
	}

	active := pages.Pages[m.current]
	frame := components.Header(components.FrameWidth, m.current, tabs) + "\n" +
		m.pages[m.current].View().Content + "\n" +
		components.Footer(components.FrameWidth, active.Path)
	frame = lipgloss.NewStyle().
		Width(components.FrameWidth).
		MaxWidth(components.FrameWidth).
		Height(components.FrameHeight).
		MaxHeight(components.FrameHeight).
		Render(frame)

	if m.termWidth < components.FrameWidth || m.termHeight < components.FrameHeight {
		frame = fmt.Sprintf("terminal too small: need %dx%d", components.FrameWidth, components.FrameHeight)
	}

	v := tea.NewView(lipgloss.Place(m.termWidth, m.termHeight, lipgloss.Center, lipgloss.Center, frame))
	v.AltScreen = true
	v.BackgroundColor = lipgloss.Color(style.ColorBackground)
	v.ForegroundColor = lipgloss.Color(style.ColorForeground)
	return v
}
