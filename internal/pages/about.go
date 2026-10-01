package pages

import (
	tea "charm.land/bubbletea/v2"
)

// About is the about page.
type About struct{}

// NewAbout returns a new about page.
func NewAbout() About { return About{} }

// Init is a no-op placeholder.
func (a About) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (a About) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return a, nil }

// View renders the about page placeholder.
func (a About) View() tea.View { return tea.NewView("About") }
