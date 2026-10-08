package pages

import (
	tea "charm.land/bubbletea/v2"
)

// Project is a single entry in the projects list.
type ExperienceRecord struct {
	CompanyNameArt  string
	CompanyName     string
	Position        string
	Location        string
	StartTime       string
	EndTime         string
	Description     string   // one or two short sentences
	Highlights      []string // bullet points in the detail panel
	Skills          []string // bracketed skills tags in the detail panel
	RelatedProjects []string // related projects URLs
}

// projects is the seed data. Adding an entry is a single slice literal; the
// list panel title count is derived from len(projects), never hard-coded.
var experienceRecords = []ExperienceRecord{
	{
		CompanyNameArt:  "Sellforte",
		CompanyName:     "Sellforte Solutions Oy",
		Position:        "Senior Data Scientist",
		StartTime:       "Oct 2025",
		EndTime:         "Now",
		Description:     "Placeholder description",
		Highlights:      []string{"Placeholder 1", "Placeholder 2"},
		Skills:          []string{"Polars", "JavaScript"},
		RelatedProjects: []string{},
	},
	{
		CompanyNameArt:  "Sellforte",
		CompanyName:     "Sellforte Solutions Oy",
		Position:        "Data Scientist",
		StartTime:       "Oct 2024",
		EndTime:         "Oct 2025",
		Description:     "Placeholder description",
		Highlights:      []string{"Placeholder 1", "Placeholder 2"},
		Skills:          []string{"Polars", "JavaScript"},
		RelatedProjects: []string{},
	},
	{
		CompanyNameArt:  "Uni of Hel",
		CompanyName:     "University of Helsinki",
		Position:        "Master's Program in Data Science",
		StartTime:       "Oct 2022",
		EndTime:         "Oct 2024",
		Description:     "Placeholder description",
		Highlights:      []string{"Placeholder 1", "Placeholder 2"},
		Skills:          []string{"Polars", "JavaScript"},
		RelatedProjects: []string{},
	},
}

var nExperienceRecords = len(experienceRecords)

// Experience is the experience page.
type Experience struct{}

// Init is a no-op placeholder.
func (e Experience) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return e, nil }

// View renders the experience page placeholder.
func (e Experience) View() tea.View { return tea.NewView("Experience") }
