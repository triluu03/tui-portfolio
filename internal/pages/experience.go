package pages

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	figure "github.com/common-nighthawk/go-figure"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
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
type Experience struct{ selected int }

// Init is a no-op placeholder.
func (e Experience) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return e, nil }

// View renders the experience page placeholder.
func (e Experience) View() tea.View { return tea.NewView("Experience") }

// renderProjects builds the two-column projects page padded to the pinned frame size.
func (p Experience) renderExperience() string {
	selectedProject := projects[p.selected]

	projectList := components.Panel(fmt.Sprintf("projects (%d)", nProjects), style.ProjectListWidth, style.ContentHeight, experienceListBody(p.selected))
	projectDetail := components.Panel("README.md", style.ProjectDetailWidth, style.ContentHeight, experienceDetailBody(selectedProject))

	body := lipgloss.JoinHorizontal(lipgloss.Top, projectList, strings.Repeat(" ", style.ColumnGap), projectDetail)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(body)
}

// experienceListBody renders the projects list, marking the selected row.
func experienceListBody(selected int) string {
	innerWidth := style.ProjectListWidth - lipgloss.Width("│  │")
	rows := make([]string, nProjects+1) // rows[0] is the blank spacer
	for i, p := range projects {
		marker, nameStyle := "  ", dimStyle
		if i == selected {
			marker, nameStyle = "> ", accentStyle
		}
		name := nameStyle.Render(marker + p.Name)
		year := dimStyle.Render(p.Year)
		gap := max(0, innerWidth-lipgloss.Width(name)-lipgloss.Width(year))
		rows[i+1] = name + strings.Repeat(" ", gap) + year
	}
	return strings.Join(rows, "\n")
}

// experienceDetailBody renders the README-style detail panel for one project.
func experienceDetailBody(p Project) string {
	tags := make([]string, len(p.Stack))
	for i, s := range p.Stack {
		tags[i] = accentStyle.Render("[" + s + "]")
	}
	lines := []string{""}
	for _, row := range figure.NewFigure(p.Name, "rectangles", true).Slicify() {
		lines = append(lines, warmStyle.Render(row))
	}
	lines = append(lines,
		"",
		dimStyle.Render(p.Title),
		strings.Join(tags, " "),
		"",
		p.Description,
		"",
		warmStyle.Render("## Highlights"),
	)
	for _, h := range p.Highlights {
		lines = append(lines, accentStyle.Render("- ")+h)
	}
	var links []string
	links = append(links, accentStyle.Render("→ source: ")+dimStyle.Render(p.Source))

	lines = append(lines, "")
	lines = append(lines, links...)
	return strings.Join(lines, "\n")
}
