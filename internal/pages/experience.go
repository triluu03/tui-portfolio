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

// ExperienceRecord is a single entry in the experience list.
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

// experienceRecords is the seed data. Adding an entry is a single slice
// literal; the list panel title count is derived from
// len(experienceRecords), never hard-coded.
var experienceRecords = []ExperienceRecord{
	{
		CompanyNameArt:  "Sellforte",
		CompanyName:     "Sellforte Solutions Oy",
		Position:        "Senior Data Scientist & Developer",
		Location:        "Espoo, Finland",
		StartTime:       "Jul 2026",
		EndTime:         "Now",
		Description:     "Main developer of three high-stake customer-facing modules.",
		Highlights:      []string{"Performance Recommendations", "Media Optimizer", "MTA"},
		Skills:          []string{"Python", "Polars", "TypeScript", "VueJS"},
		RelatedProjects: []string{},
	},
	{
		CompanyNameArt:  "Sellforte",
		CompanyName:     "Sellforte Solutions Oy",
		Position:        "Data Scientist",
		Location:        "Espoo, Finland",
		StartTime:       "Aug 2024",
		EndTime:         "Jul 2026",
		Description:     "Main developer of the Media Optimizer module.",
		Highlights:      []string{"", ""},
		Skills:          []string{"Python", "Polars", "JavaScript", "VueJS"},
		RelatedProjects: []string{},
	},
	{
		CompanyNameArt:  "Sellforte",
		CompanyName:     "Sellforte Solutions Oy",
		Position:        "Junior Data Scientist",
		Location:        "Espoo, Finland",
		StartTime:       "Jun 2024",
		EndTime:         "Aug 2026",
		Description:     "Joined as an intern for the summer.",
		Highlights:      []string{"Developed data integration and modeling pipelines\nfor company's weather analytics solution used by 7+ customers.", "Onboarded 2 customers to the company's analytics solutions\nand ensured weekly results delivery."},
		Skills:          []string{"Python", "Pandas", "Matplotlib"},
		RelatedProjects: []string{},
	},
	{
		CompanyNameArt:  "Uni of Hel",
		CompanyName:     "University of Helsinki",
		Position:        "Master in Data Science",
		Location:        "Helsinki, Finland",
		StartTime:       "Aug 2024",
		EndTime:         "Jun 2026",
		Description:     "Focused on Machine Learning and Bayesian analytics.",
		Highlights:      []string{"GPA: 4.77/5.0", "Thesis: Bayesian Experimental Design in Marketing Mix Modeling"},
		Skills:          []string{"Placeholder 1"},
		RelatedProjects: []string{"https://helda.helsinki.fi/items/48e5ae2f-8bf0-4c87-b3a2-6f47b34d91bf"},
	},
	{
		CompanyNameArt:  "Uni of Hel",
		CompanyName:     "University of Helsinki",
		Position:        "Bachelor in Science",
		Location:        "Helsinki, Finland",
		StartTime:       "Aug 2021",
		EndTime:         "Jun 2024",
		Description:     "Major: Computer Science and Data Science.",
		Highlights:      []string{"GPA: 4.93/5.0", "Thesis: Application of Deep Learning in Drug Response Prediction"},
		Skills:          []string{"Placeholder 1"},
		RelatedProjects: []string{},
	},
}

var nExperienceRecords = len(experienceRecords)

// Experience is the experience page.
type Experience struct{ selected int }

// Init is a no-op placeholder.
func (e Experience) Init() tea.Cmd { return nil }

// Update handles list navigation and ignores all other messages.
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "up", "k":
			e.selected = (e.selected + nExperienceRecords - 1) % nExperienceRecords
		case "down", "j":
			e.selected = (e.selected + 1) % nExperienceRecords
		}
	}
	return e, nil
}

// View renders the experience list and detail panels.
func (e Experience) View() tea.View { return tea.NewView(e.renderExperience()) }

// renderExperience builds the two-column experience page padded to the pinned frame size.
func (e Experience) renderExperience() string {
	selectedRecord := experienceRecords[e.selected]

	experienceList := components.Panel(fmt.Sprintf("experience (%d)", nExperienceRecords), style.ProjectListWidth, style.ContentHeight, experienceListBody(e.selected))
	experienceDetail := components.Panel("README.md", style.ProjectDetailWidth, style.ContentHeight, experienceDetailBody(selectedRecord))

	body := lipgloss.JoinHorizontal(lipgloss.Top, experienceList, strings.Repeat(" ", style.ColumnGap), experienceDetail)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(body)
}

// experienceListBody renders the experience list, marking the selected record.
func experienceListBody(selected int) string {
	rows := []string{""} // leading blank spacer
	for i, r := range experienceRecords {
		marker, lineStyle := "  ", dimStyle
		if i == selected {
			marker, lineStyle = "> ", accentStyle
		}
		rows = append(rows,
			lineStyle.Render(marker+r.CompanyName),
			lineStyle.Render("  "+r.Position),
			lineStyle.Render("  "+r.StartTime+" - "+r.EndTime),
		)
		if i < nExperienceRecords-1 {
			rows = append(rows, "")
		}
	}
	return strings.Join(rows, "\n")
}

// experienceDetailBody renders the README-style detail panel for one record.
func experienceDetailBody(r ExperienceRecord) string {
	tags := make([]string, len(r.Skills))
	for i, s := range r.Skills {
		tags[i] = accentStyle.Render("[" + s + "]")
	}
	lines := []string{""}
	for _, row := range figure.NewFigure(r.CompanyNameArt, "rectangles", true).Slicify() {
		lines = append(lines, warmStyle.Render(row))
	}
	lines = append(lines,
		"",
		dimStyle.Render(r.CompanyName),
		dimStyle.Render(r.Position),
		dimStyle.Render(r.Location),
		"",
		strings.Join(tags, " "),
		"",
		r.Description,
		"",
		warmStyle.Render("## Highlights"),
	)
	for _, h := range r.Highlights {
		for i, part := range strings.Split(h, "\n") {
			if i == 0 {
				lines = append(lines, accentStyle.Render("- ")+part)
			} else {
				lines = append(lines, "  "+part)
			}
		}
	}
	lines = append(lines, "")
	for _, url := range r.RelatedProjects {
		lines = append(lines, accentStyle.Render("→ related: ")+dimStyle.Render(url))
	}
	return strings.Join(lines, "\n")
}
