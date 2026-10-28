package pages

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/charmbracelet/x/ansi"
	figure "github.com/common-nighthawk/go-figure"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// ExperienceRecord is a single entry in the experience list.
type ExperienceRecord struct {
	CompanyNameArt string
	CompanyName    string
	Position       string
	Location       string
	StartTime      string
	EndTime        string
	Description    string
	Highlights     []string
	Skills         []string
	RelatedLink    string
}

// experienceRecords is the seed data. Adding an entry is a single slice
// literal; the list panel title count is derived from
// len(experienceRecords), never hard-coded.
var experienceRecords = []ExperienceRecord{
	{
		CompanyNameArt: "Sellforte",
		CompanyName:    "Sellforte Solutions Oy",
		Position:       "Senior Data Scientist",
		Location:       "Espoo, Finland",
		StartTime:      "Jul 2026",
		EndTime:        "Now",
		Description:    "Main developer of three high-stakes customer-facing products.",
		Highlights: []string{
			"Primary developer of the {{Performance product}}, used daily by {{all}} " + "\n" +
				"{{digital marketing customers (40+)}}, that reports historical campaign" + "\n" +
				"and ad-group performance, generates daily bidding parameter" + "\n" +
				"recommendations, and analyzes the impact of previously made changes.",
			"Continues to develop and maintain the Media Optimizer product," + "\n" +
				"which has become the {{backbone of other services}} provided by" + "\n" +
				"the company, such as {{Home}} or {{Sellforte AI}}.",
			"Owns the {{Pixel and MTA analytics solutions}}, including design and" + "\n" +
				"implementation of the ingestion pipeline and attribution module" + "\n" +
				"that delivers analyzed results daily to customers.",
		},
		Skills:      []string{"Python", "Polars", "TypeScript", "Vue.js", "Airflow", "AWS services"},
		RelatedLink: "https://demo.sellforte.com/",
	},
	{
		CompanyNameArt: "Sellforte",
		CompanyName:    "Sellforte Solutions Oy",
		Position:       "Data Scientist",
		Location:       "Espoo, Finland",
		StartTime:      "Aug 2024",
		EndTime:        "Jul 2026",
		Description:    "Main developer of the Media Optimizer product.",
		Highlights: []string{
			"Designed and built the {{media optimization module}}, a production" + "\n" +
				"service that provides {{real-time}} optimal investment allocations" + "\n" +
				"across sales channels for {{30+ customers}}, supporting marketing" + "\n" +
				"budget planning and execution.",
			"Improved the optimization runtime performance {{by 20%}} through" + "\n" +
				"migrating the data-processing codebase from {{Pandas to Polars}}.",
			"Built and maintained data integration and processing pipelines" + "\n" +
				"for multiple ad platforms, including {{Google Ads, Facebook Ads,}}" + "\n" +
				"{{and Pinterest Ads}}, improving the quality and reliability of" + "\n" +
				"the analytics solutions.",
			"Onboarded {{15+ customers}} and ensured reliable {{daily delivery}} of " + "\n" +
				"analytics results in production.",
		},
		Skills:      []string{"Python", "Polars", "JavaScript", "VueJS", "Airflow", "AWS services"},
		RelatedLink: "https://sellforte.com",
	},
	{
		CompanyNameArt: "Sellforte",
		CompanyName:    "Sellforte Solutions Oy",
		Position:       "Junior Data Scientist",
		Location:       "Espoo, Finland",
		StartTime:      "Jun 2024",
		EndTime:        "Aug 2024",
		Description:    "Joined as an intern for the summer.",
		Highlights: []string{
			"Designed and built data integration and modeling pipelines" + "\n" +
				"for the company's {{weather analytics solution}} used by {{7+ customers}}.",
			"Onboarded {{2 customers}} to the company's analytics solutions" + "\n" +
				"and ensured {{weekly results}} delivery.",
		},
		Skills:      []string{"Python", "Pandas", "Matplotlib", "Airflow", "AWS services"},
		RelatedLink: "https://sellforte.com",
	},
	{
		CompanyNameArt: "Uni of Hel",
		CompanyName:    "University of Helsinki",
		Position:       "Master of Science",
		Location:       "Helsinki, Finland",
		StartTime:      "Aug 2024",
		EndTime:        "Jun 2026",
		Description:    "Focused on Machine Learning and Bayesian analytics.",
		Highlights:     []string{"GPA: 4.77/5.0", "Thesis: Bayesian Experimental Design in Marketing Mix Modeling"},
		Skills:         []string{"Bayesian analytics", "ML system engineering"},
		RelatedLink:    "https://helda.helsinki.fi/items/48e5ae2f-8bf0-4c87-b3a2-6f47b34d91bf",
	},
	{
		CompanyNameArt: "Signant Health",
		CompanyName:    "Signant Health Oy",
		Position:       "Software Trainee, R&D",
		Location:       "Helsinki, Finland",
		StartTime:      "Jan 2023",
		EndTime:        "Aug 2023",
		Description:    "Trainee in the R&D department.",
		Highlights: []string{
			"Ensured software product quality by implementing automated tests" + "\n" +
				"using {{Python and Robot Framework}}.",
			"Updated the company's cloud services by migrating AWS Lambda" + "\n" +
				"functions to the {{JavaScript SDK version 3}}.",
		},
		Skills:      []string{"Test Automation", "Python", "Robot Framework"},
		RelatedLink: "https://signanthealth.com",
	},
	{
		CompanyNameArt: "Uni of Hel",
		CompanyName:    "University of Helsinki",
		Position:       "Bachelor of Science",
		Location:       "Helsinki, Finland",
		StartTime:      "Aug 2021",
		EndTime:        "Jun 2024",
		Description:    "Major: Computer Science and Data Science.",
		Highlights:     []string{"GPA: 4.93/5.0", "Thesis: Application of Deep Learning in Drug Response Prediction"},
		Skills:         []string{"Software engineering", "Data analytics"},
	},
}

var nExperienceRecords = len(experienceRecords)

// Experience is the experience page.
type Experience struct {
	selected int
	focus    *bool
	copied   bool
}

// Init is a no-op placeholder.
func (e Experience) Init() tea.Cmd { return nil }

// Update handles list navigation and ignores all other messages.
func (e Experience) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(copyFeedbackMsg); ok {
		e.copied = false
		return e, nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return e, nil
	}

	if e.focus == nil || !*e.focus {
		e.copied = false
		switch key.String() {
		case "up", "k":
			e.selected = (e.selected + nExperienceRecords - 1) % nExperienceRecords
		case "down", "j":
			e.selected = (e.selected + 1) % nExperienceRecords
		}
	}

	if *e.focus {
		switch key.String() {
		case "enter":
			e.copied = true
			return e, copyCmd(experienceRecords[e.selected].RelatedLink)
		}
	}

	return e, nil
}

// View renders the experience list and detail panels.
func (e Experience) View() tea.View { return tea.NewView(e.renderExperience()) }

// renderExperience builds the two-column experience page padded to the pinned frame size.
func (e Experience) renderExperience() string {
	selectedRecord := experienceRecords[e.selected]

	experienceList := experienceListPanel(e.selected, !*e.focus)
	experienceDetail := experienceDetailPanel(selectedRecord, *e.focus, e.copied)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		experienceList,
		strings.Repeat(" ", style.ColumnGap),
		experienceDetail,
	)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(body)
}

// experienceListPanel renders the experience list, marking the selected record.
func experienceListPanel(selected int, focused bool) string {
	entries := make([]listEntry, len(experienceRecords))
	for i, r := range experienceRecords {
		entries[i] = listEntry{lines: []string{r.CompanyName, r.Position, r.StartTime + " - " + r.EndTime}}
	}
	return components.Panel(
		fmt.Sprintf("experience (%d)", nExperienceRecords),
		style.ProjectListWidth,
		style.ContentHeight,
		listBody(entries, selected),
		focused,
	)
}

// experienceDetailPanel renders the README-style detail panel for one record.
func experienceDetailPanel(r ExperienceRecord, focused bool, copied bool) string {
	tags := make([]string, len(r.Skills))
	for i, s := range r.Skills {
		tags[i] = accentStyle.Render("[" + s + "]")
	}
	lines := []string{""}
	for _, row := range figure.NewFigure(r.CompanyNameArt, "thin", true).Slicify() {
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
				lines = append(lines, accentStyle.Render("- ")+renderEmphasis(part))
			} else {
				lines = append(lines, "  "+renderEmphasis(part))
			}
		}
	}

	lines = append(lines, "")
	if r.RelatedLink != "" {
		if focused {
			innerWidth := style.ProjectDetailWidth - lipgloss.Width("│  │")
			linkWidth := innerWidth - lipgloss.Width("→ related: ") - lipgloss.Width(" (enter to copy link ↗)")
			hint, hintStyle := " (enter to copy link ↗)", accentStyle
			if copied {
				hint, hintStyle = " (copied ✓)", successStyle
			}
			lines = append(lines,
				accentStyle.Render("→ related: ")+
					hyperlink(r.RelatedLink, accentStyle.Render(ansi.Truncate(r.RelatedLink, linkWidth, "…")))+
					hintStyle.Render(hint),
			)
		} else {
			lines = append(lines,
				accentStyle.Render("→ related: ")+
					hyperlink(r.RelatedLink, dimStyle.Render(r.RelatedLink)),
			)
		}
	}

	return components.Panel(
		"README.md",
		style.ProjectDetailWidth,
		style.ContentHeight,
		strings.Join(lines, "\n"),
		focused,
	)
}
