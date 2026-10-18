package pages

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

// Descriptor is the single source of truth for a page's navigation metadata and
// model. Header tabs and the footer breadcrumb are derived from these entries.
type Descriptor struct {
	Title string                                // label shown in the header
	Path  string                                // label shown in the footer
	New   func(focus, editMode *bool) tea.Model // page model constructor
}

// Pages lists the application's pages in navigation order.
var Pages = []Descriptor{
	{
		Title: "About",
		Path:  "~/portfolio",
		New:   func(focus, _ *bool) tea.Model { return About{focus: focus} },
	},
	{
		Title: "Projects",
		Path:  "~/portfolio/projects",
		New:   func(focus, _ *bool) tea.Model { return Projects{focus: focus} },
	},
	{
		Title: "Experience",
		Path:  "~/portfolio/experience",
		New:   func(focus, _ *bool) tea.Model { return Experience{focus: focus} },
	},
	{
		Title: "Contact",
		Path:  "~/portfolio/contact",
		New:   func(focus, editMode *bool) tea.Model { return Contact{focus: focus, editMode: editMode} },
	},
}

// hyperlink wraps text in an OSC 8 escape sequence so the user's terminal
// renders it as a clickable link. Clicking opens the URL in the client's
// browser, which works locally and over SSH.
func hyperlink(url, text string) string {
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}
