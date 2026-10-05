package pages

import (
	"strings"
	"testing"

	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
)

func TestAboutViewSize(t *testing.T) {
	lines := strings.Split(About{}.View().Content, "\n")
	if len(lines) != components.ContentHeight {
		t.Errorf("About view height = %d lines, want %d", len(lines), components.ContentHeight)
	}
	for i, line := range lines {
		if w := lipgloss.Width(line); w != components.FrameWidth {
			t.Errorf("About view line %d width = %d, want %d", i, w, components.FrameWidth)
		}
	}
}

func TestNameArt(t *testing.T) {
	lines := strings.Split(nameArt, "\n")
	if len(lines) != 5 {
		t.Errorf("nameArt lines = %d, want 5", len(lines))
	}
	want := lipgloss.Width(lines[0])
	for i, line := range lines {
		if w := lipgloss.Width(line); w != want {
			t.Errorf("nameArt line %d width = %d, want %d (%q)", i, w, want, line)
		}
	}
}

func TestPanelSize(t *testing.T) {
	body := strings.TrimSuffix(strings.Repeat("x\n", 40), "\n")
	lines := strings.Split(components.Panel("whoami", 116, 34, body), "\n")
	if len(lines) != 34 {
		t.Errorf("Panel height = %d lines, want 34", len(lines))
	}
	for i, line := range lines {
		if w := lipgloss.Width(line); w != 116 {
			t.Errorf("Panel line %d width = %d, want 116 (%q)", i, w, line)
		}
	}
	if got := components.Panel("whoami", 0, 34, "x"); got != "" {
		t.Errorf("Panel(width 0) = %q, want empty", got)
	}
}

func TestPanelTruncatesOverwide(t *testing.T) {
	title := strings.Repeat("T", 200)
	body := strings.Repeat("x", 200)
	lines := strings.Split(components.Panel(title, 40, 4, body), "\n")
	if len(lines) != 4 {
		t.Fatalf("Panel height = %d lines, want 4", len(lines))
	}
	for i, line := range lines {
		if w := lipgloss.Width(line); w != 40 {
			t.Errorf("Panel line %d width = %d, want 40 (%q)", i, w, line)
		}
	}
}
