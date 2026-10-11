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
		StartTime:   "Jun 2026",
		EndTime:     "Aug 2026",
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

	projectList := components.Panel(
		fmt.Sprintf("projects (%d)", nProjects),
		style.ProjectListWidth,
		style.ContentHeight,
		projectListBody(p.selected),
		true,
	)
	projectDetail := components.Panel(
		"README.md",
		style.ProjectDetailWidth,
		style.ContentHeight,
		projectDetailBody(selectedProject),
		false,
	)

	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		projectList,
		strings.Repeat(" ", style.ColumnGap),
		projectDetail,
	)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(body)
}

// projectListBody renders the projects list, marking the selected row.
func projectListBody(selected int) string {
	rows := []string{""}
	for i, p := range projects {
		marker, lineStyle := "  ", dimStyle
		if i == selected {
			marker, lineStyle = "> ", accentStyle
		}
		rows = append(rows,
			lineStyle.Render(marker+p.Name),
			lineStyle.Render("  "+p.StartTime+" - "+p.EndTime),
		)
		if i < nProjects-1 {
			rows = append(rows, "")
		}

	}
	return strings.Join(rows, "\n")
}

// projectDetailBody renders the README-style detail panel for one project.
func projectDetailBody(p Project) string {
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

	if p.Source != "" {
		links = append(links, accentStyle.Render("→ source: ")+dimStyle.Render(p.Source))
	}

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
