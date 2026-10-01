package pages

import (
	tea "charm.land/bubbletea/v2"
)

// Experience is the experience page.
type Experience struct{}

// NewExperience returns a new experience page.
func NewExperience() Experience { return Experience{} }

// Init is a no-op placeholder.
func (e Experience) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return e, nil }

// View renders the experience page placeholder.
func (e Experience) View() tea.View { return tea.NewView("Experience") }
