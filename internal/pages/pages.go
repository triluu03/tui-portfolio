package pages

import (
	"log"
	"os/exec"
	"runtime"

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
		if err := cmd.Start(); err != nil {
			log.Printf("openURL %s: %v", url, err)
		}
		return nil
	}
}
