package pages

import (
	tea "charm.land/bubbletea/v2"
)

// Projects is the projects page.
type Projects struct{}

// NewProjects returns a new projects page.
func NewProjects() Projects { return Projects{} }

// Init is a no-op placeholder.
func (p Projects) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (p Projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

// View renders the projects page placeholder.
func (p Projects) View() tea.View { return tea.NewView("Projects") }
