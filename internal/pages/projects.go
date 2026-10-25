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
		Title:       "Terminal UI Portfolio",
		StartTime:   "Oct 2026",
		EndTime:     "Now",
		Description: "A terminal portfolio built in Go and Bubble Tea.",
		Highlights: []string{
			"Built an {{interactive terminal portfolio}} in {{Go}} with {{Bubble Tea}}," + "\n" +
				"which anyone can open with a single command (`{{ssh hello@triluu.dev}}`)," + "\n" +
				"with no installations, browser, or account needed.",
			"Deployed it as a {{self-hosted service on a Linux server}}, using" + "\n" +
				"{{OpenSSH}} `{{ForceCommand}}` to launch the app for every connection" + "\n" +
				"on a dedicated shell-less user.",
		},
		Stack:  []string{"Go", "Bubble Tea", "OpenSSH configuration", "Linux server administration"},
		Source: "https://github.com/triluu03/tui-portfolio",
	},
	{
		Name:        "pennysheet",
		Title:       "Pennysheet",
		StartTime:   "May 2026",
		EndTime:     "Sep 2026",
		Description: "A personal finance tracking app based on event-sourcing.",
		Highlights: []string{
			"Designed and built {{Pennysheet}}, a full-stack personal finance app" + "\n" +
				"that syncs bank transactions via the Enable Banking API.",
			"Implemented an {{event-sourced backend in Rust (Axum)}} with an" + "\n" +
				"{{append-only}} event table in {{PostgreSQL}} and projections for" + "\n" +
				"read models, ensuring the full transaction and budget tracibility.",
			"Built a {{Model Context Protocol (MCP)}} server that gives AI" + "\n" +
				"assistants {{read-only access}} to the projections containing" + "\n" +
				"personal financial data, allowing users to analyze personal" + "\n" +
				"finances in natural language.",
			"Automated {{daily bank synchronization}} and Telegram notifications" + "\n" +
				"with background jobs at 8:30 every morning, giving an overview of" + "\n" +
				"current personal budgets {{without opening the app.}}",
		},
		Stack:  []string{"Rust", "Axum", "TypeScript", "React", "REST API", "MCP", "PostgreSQL"},
		Source: "https://github.com/triluu03/pennysheet",
	},
	{
		Name:        "sf-cli",
		Title:       "Sellforte Command Line Interface (sf-cli)",
		StartTime:   "Jun 2026",
		EndTime:     "Now",
		Description: "An internal productivity tool for Sellforte data scientists.",
		Highlights: []string{
			"Designed and built {{sf-cli}}, an internal command-line tool in" + "\n" +
				"{{Rust}} that streamlines day-to-day workflows for Sellforte's" + "\n" +
				"data scientists.",
			"Enabled data scientists to {{run pipeline tasks locally}} for" + "\n" +
				"any chosen customer with a single command, removing manual" + "\n" +
				"environment setup and speeding up development and debugging.",
			"Built a terminal client for {{PostgreSQL databases on AWS RDS}}," + "\n" +
				"letting users connect and run SQL queries directly from the" + "\n" +
				"command line without other separate database tooling.",
		},
		Stack: []string{"Rust", "Tokio", "PostgreSQL"},
	},
	{
		Name:        "foodmap",
		Title:       "Foodmap",
		StartTime:   "Mar 2026",
		EndTime:     "Apr 2026",
		Description: "A small map project built with React and SpacetimeDB.",
		Highlights: []string{
			"Designed and built {{Foodmap}}, a map-based social platform" + "\n" +
				"where users share and rate food places, as part of a one-month" + "\n" +
				"challenge in which a friend and I built the same product using" + "\n" +
				"different tech stacks.",
			"Implemented the full stack with {{React (TypeScript)}} and {{Maplibre}}" + "\n" +
				"for the interactive map, and {{SpacetimeDB (Rust)}} as a combined" + "\n" +
				"backend and database.",
			"Compared my build with the {{Phoenix and Ash (Elixir)}} implementation" + "\n" +
				"of the same feature set, evaluating the development speed," + "\n" +
				"architecture, tooling, and operations, and presented the" + "\n" +
				"findings in a retrospective.",
		},
		Stack:  []string{"Rust", "SpacetimeDB", "TypeScript", "React", "Maplibre GL"},
		Source: "https://github.com/triluu03/foodmap",
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
	for _, row := range figure.NewFigure(p.Name, "thin", true).Slicify() {
		lines = append(lines, warmStyle.Render(row))
	}
	lines = append(lines,
		"",
		accentStyle.Render(p.Title),
		dimStyle.Render(p.Description),
		"",
		strings.Join(tags, " "),
		"",
		warmStyle.Render("## Highlights"),
	)
	for _, h := range p.Highlights {
		for i, part := range strings.Split(h, "\n") {
			if i == 0 {
				lines = append(lines, accentStyle.Render("- ")+renderEmphasis(part))
			} else {
				lines = append(lines, "  "+renderEmphasis(part))
			}
		}
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
