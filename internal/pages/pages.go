package pages

import (
	tea "charm.land/bubbletea/v2"
)

// Descriptor is the single source of truth for a page's navigation metadata and
// model. Header tabs and the footer breadcrumb are derived from these entries.
type Descriptor struct {
	Title string    // label shown in the header
	Path  string    // label shown in the footer
	Model tea.Model // page model
}

// Pages lists the application's pages in navigation order.
var Pages = []Descriptor{
	{Title: "About", Path: "~/portfolio", Model: About{}},
	{Title: "Projects", Path: "~/portfolio/projects", Model: Projects{}},
	{Title: "Experience", Path: "~/portfolio/experience", Model: Experience{}},
	{Title: "Contact", Path: "~/portfolio/contact", Model: Contact{}},
}
