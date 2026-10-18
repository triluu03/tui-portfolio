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

// Project is a single entry in the projects list.
type Project struct {
	Name        string
	Title       string
	StartTime   string
	EndTime     string
	Description string
	Highlights  []string
	Stack       []string
	Source      string
}

// projects is the seed data. Adding an entry is a single slice literal; the
// list panel title count is derived from len(projects), never hard-coded.
var projects = []Project{
	{
		Name:        "tui-portfolio",
		Title:       "TUI Portfolio",
		StartTime:   "Oct 2026",
		EndTime:     "Now",
		Description: "A terminal portfolio built in Go.",
		Highlights:  []string{"Host a secured SSH server locally."},
		Stack:       []string{"Go", "Bubble Tea", "Lip Gloss"},
		Source:      "https://github.com/triluu03/tui-portfolio",
	},
	{
		Name:        "pennysheet",
		Title:       "Pennysheet",
		StartTime:   "May 2026",
		EndTime:     "Sep 2026",
		Description: "A personal finance tracking app based on event-sourcing.",
		Highlights:  []string{"Built a robust event-sourcing architecture.", "Integrated to Claude/Codex with a MCP server.", "Implemented real-time expenses tracking."},
		Stack:       []string{"Rust", "Axum", "TypeScript", "React", "REST API", "MCP", "PostgreSQL"},
		Source:      "https://github.com/triluu03/pennysheet",
	},
	{
		Name:        "sf-cli",
		Title:       "Sellforte Command Line Interface (sf-cli)",
		StartTime:   "Jun 2026",
		EndTime:     "Now",
		Description: "An internal productivity tool for Sellforte data scientists.",
		Highlights:  []string{"Automatically set up computer environments for data science work.", "Connect and run SQL queries to PostgreSQL databases."},
		Stack:       []string{"Rust", "Tokio", "PostgreSQL"},
	},
}

var nProjects = len(projects)

// Projects is the projects page.
type Projects struct {
	selected int
	focus    *bool
}

// Init is a no-op placeholder.
func (p Projects) Init() tea.Cmd { return nil }

// Update handles list navigation and ignores all other messages.
func (p Projects) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return p, nil
	}

	if p.focus == nil || !*p.focus {
		switch key.String() {
		case "up", "k":
			p.selected = (p.selected + nProjects - 1) % nProjects
		case "down", "j":
			p.selected = (p.selected + 1) % nProjects
		}
		return p, nil
	}

	if *p.focus {
		switch key.String() {
		case "enter":
			return p, tea.SetClipboard(projects[p.selected].Source)
		}
	}

	return p, nil
}

// View renders the projects list and detail panels.
func (p Projects) View() tea.View { return tea.NewView(p.renderProjects()) }

// renderProjects builds the two-column projects page padded to the pinned frame size.
func (p Projects) renderProjects() string {
	selectedProject := projects[p.selected]

	projectList := projectListBody(p.selected, !*p.focus)
	projectDetail := projectDetailBody(selectedProject, *p.focus)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		projectList,
		strings.Repeat(" ", style.ColumnGap),
		projectDetail,
	)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(body)
}

// projectListBody renders the projects list, marking the selected row.
func projectListBody(selected int, focused bool) string {
	entries := make([]listEntry, len(projects))
	for i, p := range projects {
		entries[i] = listEntry{lines: []string{p.Name, p.StartTime + " - " + p.EndTime}}
	}
	return components.Panel(
		fmt.Sprintf("projects (%d)", nProjects),
		style.ProjectListWidth,
		style.ContentHeight,
		listBody(entries, selected),
		focused,
	)
}

// projectDetailBody renders the README-style detail panel for one project.
func projectDetailBody(p Project, focused bool) string {
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

	lines = append(lines, "")
	if p.Source != "" {
		if focused {
			innerWidth := style.ProjectDetailWidth - lipgloss.Width("│  │")
			linkWidth := innerWidth - lipgloss.Width("→ source: ") - lipgloss.Width(" (enter to copy link ↗)")
			lines = append(lines,
				accentStyle.Render("→ source: ")+
					hyperlink(p.Source, accentStyle.Render(ansi.Truncate(p.Source, linkWidth, "…")))+
					accentStyle.Render(" (enter to copy link ↗)"),
			)
		} else {
			lines = append(lines,
				accentStyle.Render("→ source: ")+
					hyperlink(p.Source, dimStyle.Render(p.Source)),
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
