package pages

import (
	tea "charm.land/bubbletea/v2"
)

// Descriptor is the single source of truth for a page's navigation metadata and
// model. Header tabs and the footer breadcrumb are derived from these entries.
type Descriptor struct {
	Title string                      // label shown in the header
	Path  string                      // label shown in the footer
	New   func(focus *bool) tea.Model // page model constructor
}

// Pages lists the application's pages in navigation order.
var Pages = []Descriptor{
	{Title: "About", Path: "~/portfolio", New: func(focus *bool) tea.Model { return About{focus: focus} }},
	{Title: "Projects", Path: "~/portfolio/projects", New: func(focus *bool) tea.Model { return Projects{focus: focus} }},
	{Title: "Experience", Path: "~/portfolio/experience", New: func(focus *bool) tea.Model { return Experience{focus: focus} }},
	{Title: "Contact", Path: "~/portfolio/contact", New: func(focus *bool) tea.Model { return Contact{focus: focus} }},
}
