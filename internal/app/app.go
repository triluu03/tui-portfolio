package app

import (
	tea "charm.land/bubbletea/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/pages"
)

// Model is the root application model. It owns the list of pages and tracks
// which page is currently active.
type Model struct {
	pages   []tea.Model
	current int
}

// New returns a root model with the pages wired up.
func New() Model {
	return Model{
		pages: []tea.Model{
			pages.NewAbout(),
			pages.NewProjects(),
			pages.NewExperience(),
			pages.NewContact(),
		},
	}
}

// Init returns the initial commands from every page.
func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	for _, p := range m.pages {
		cmds = append(cmds, p.Init())
	}
	return tea.Batch(cmds...)
}

// Update handles global key bindings and delegates messages to the active page.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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

// View composes the header, active page, and footer into a single screen.
func (m Model) View() tea.View {
	v := tea.NewView(
		components.Header() + "\n" +
			m.pages[m.current].View().Content + "\n" +
			components.Footer(),
	)
	v.AltScreen = true
	return v
}
