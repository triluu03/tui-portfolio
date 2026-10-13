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
	pages   []tea.Model
	current int

	termWidth  int
	termHeight int

	// Global states shared with all pages
	focusRight *bool
	editMode   *bool
}

// Return a new model with the pages wired up into the states.
func New() Model {
	focus := new(bool)
	editMode := new(bool)

	var m Model
	m.focusRight = focus
	m.editMode = editMode

	for _, d := range pages.Pages {
		m.pages = append(m.pages, d.New(focus, editMode))
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
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if *m.editMode {
			switch msg.String() {
			case "esc":
				*m.editMode = false
				return m, nil
			}
		}

		if !*m.editMode {
			switch msg.String() {
			case "q":
				return m, tea.Quit
			case "left", "h", "esc":
				*m.focusRight = false
				return m, nil
			case "right", "l", "enter":
				*m.focusRight = true
				return m, nil
			case "1", "2", "3", "4":
				m.current = int(msg.String()[0] - '1')
				*m.focusRight = false
				return m, nil
			}
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
	frame := components.Header(style.FrameWidth, m.current, tabs) + "\n" +
		m.pages[m.current].View().Content + "\n" +
		components.Footer(style.FrameWidth, active.Path)
	frame = lipgloss.NewStyle().
		Width(style.FrameWidth).
		MaxWidth(style.FrameWidth).
		Height(style.FrameHeight).
		MaxHeight(style.FrameHeight).
		Render(frame)

	if m.termWidth < style.FrameWidth || m.termHeight < style.FrameHeight {
		frame = fmt.Sprintf("terminal too small: need %dx%d", style.FrameWidth, style.FrameHeight)
	}

	v := tea.NewView(lipgloss.Place(m.termWidth, m.termHeight, lipgloss.Center, lipgloss.Center, frame))
	v.AltScreen = true
	v.BackgroundColor = lipgloss.Color(style.ColorBackground)
	v.ForegroundColor = lipgloss.Color(style.ColorForeground)
	return v
}
