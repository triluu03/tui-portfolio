package pages

import (
	tea "charm.land/bubbletea/v2"
)

// Contact is the contact page.
type Contact struct{}

// Init is a no-op placeholder.
func (c Contact) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return c, nil }

// View renders the contact page placeholder.
func (c Contact) View() tea.View { return tea.NewView("Contact") }
