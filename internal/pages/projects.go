package pages

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	figure "github.com/common-nighthawk/go-figure"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// Project is a single entry in the projects list.
type Project struct {
	Name        string   // list label
	Title       string   // detail heading
	Year        string   // shown right-aligned in the list
	Description string   // one or two short sentences
	Highlights  []string // bullet points in the detail panel
	Stack       []string // bracketed tags in the detail panel
	Source      string   // source URL
}

// projects is the seed data. Adding an entry is a single slice literal; the
// list panel title count is derived from len(projects), never hard-coded.
var projects = []Project{
	{
		Name:        "tui-portfolio",
		Title:       "TUI Portfolio",
		Year:        "Oct 2026",
		Description: "A terminal portfolio rendered in a pinned 120x36 frame.",
		Highlights:  []string{"Shared palette and spacing tokens", "Multi-page navigation with q/1-4 keys"},
		Stack:       []string{"Go", "Bubble Tea", "Lip Gloss"},
		Source:      "https://github.com/triluu03/tui-portfolio",
	},
	{
		Name:        "pennysheet",
		Title:       "Pennysheet",
		Year:        "2024",
		Description: "XGBoost classifier scoring monthly churn risk for subscriptions.",
		Highlights:  []string{"AUC 0.91 on held-out cohort", "SHAP explanations per account"},
		Stack:       []string{"Python", "scikit-learn", "XGBoost"},
		Source:      "https://github.com/triluu03/churn-model",
	},
	{
		Name:        "etl-pipeline",
		Title:       "Warehouse ETL Pipeline",
		Year:        "2023",
		Description: "Airflow DAGs and dbt models loading nightly product telemetry.",
		Highlights:  []string{"Backfilled three years of history", "Downstream dashboards refresh on time"},
		Stack:       []string{"Airflow", "dbt", "PostgreSQL"},
		Source:      "https://github.com/triluu03/etl-pipeline",
	},
	{
		Name:        "viz-dashboard",
		Title:       "Metrics Dashboard",
		Year:        "2022",
		Description: "Dashboards generated from warehouse queries and served as flat files.",
		Highlights:  []string{"Renders in under a second per page", "No backend needed after build"},
		Stack:       []string{"SQL", "Go", "HTML"},
		Source:      "https://github.com/triluu03/viz-dashboard",
	},
}

var nProjects = len(projects)

// Projects is the projects page.
type Projects struct{ selected int }

// Init is a no-op placeholder.
func (p Projects) Init() tea.Cmd { return nil }

// Update handles list navigation and ignores all other messages.
func (p Projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok {
		switch key.String() {
		case "up", "k":
			p.selected = (p.selected + nProjects - 1) % nProjects
		case "down", "j":
			p.selected = (p.selected + 1) % nProjects
		case "enter":
			return p, openURL(projects[p.selected].Source)
		}
	}
	return p, nil
}

// View renders the projects list and detail panels.
func (p Projects) View() tea.View { return tea.NewView(p.renderProjects()) }

// renderProjects builds the two-column projects page padded to the pinned frame size.
func (p Projects) renderProjects() string {
	selectedProject := projects[p.selected]

	projectList := components.Panel(fmt.Sprintf("projects (%d)", nProjects), style.ProjectListWidth, style.ContentHeight, listBody(p.selected))
	projectDetail := components.Panel("README.md", style.ProjectDetailWidth, style.ContentHeight, detailBody(selectedProject))

	body := lipgloss.JoinHorizontal(lipgloss.Top, projectList, strings.Repeat(" ", style.ColumnGap), projectDetail)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(body)
}

// listBody renders the projects list, marking the selected row.
func listBody(selected int) string {
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

// detailBody renders the README-style detail panel for one project.
func detailBody(p Project) string {
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

// openURL launches url in the user's default browser without blocking the UI.
func openURL(url string) tea.Cmd {
	if url == "" {
		return nil
	}
	return func() tea.Msg {
		var cmd *exec.Cmd
		switch runtime.GOOS {
		case "darwin":
			cmd = exec.Command("open", url)
		case "windows":
			cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
		default:
			cmd = exec.Command("xdg-open", url)
		}
		_ = cmd.Start()
		return nil
	}
}
