package app

import (
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func TestViewSetsTerminalStyle(t *testing.T) {
	v := New().View()
	if !v.AltScreen {
		t.Error("View().AltScreen = false, want true")
	}
	if v.BackgroundColor == nil {
		t.Error("View().BackgroundColor is not set")
	}
	if v.ForegroundColor == nil {
		t.Error("View().ForegroundColor is not set")
	}
}

func TestViewPinsAndCentersFrame(t *testing.T) {
	m := New()
	next, _ := m.Update(tea.WindowSizeMsg{Width: 160, Height: 50})
	lines := strings.Split(next.(Model).View().Content, "\n")

	if len(lines) != 50 {
		t.Fatalf("view height = %d lines, want 50", len(lines))
	}
	for i, l := range lines {
		if w := lipgloss.Width(l); w != 160 {
			t.Fatalf("line %d width = %d, want 160", i, w)
		}
	}

	// The appWidth x appHeight frame is centered, so it starts at column
	// (160-appWidth)/2 and row (50-appHeight)/2.
	left := strings.Repeat(" ", (160-appWidth)/2)
	if !strings.HasPrefix(lines[(50-appHeight)/2], left) {
		t.Errorf("frame is not horizontally centered on row %d", (50-appHeight)/2)
	}
}

func TestViewTooSmallShowsHint(t *testing.T) {
	m := New()
	next, _ := m.Update(tea.WindowSizeMsg{Width: 60, Height: 20})
	got := next.(Model).View().Content
	if !strings.Contains(got, "too small") {
		t.Errorf("small terminal view = %q, want a \"too small\" hint", got)
	}
	lines := strings.Split(got, "\n")
	if len(lines) != 20 {
		t.Errorf("view height = %d lines, want 20", len(lines))
	}
	for i, l := range lines {
		if w := lipgloss.Width(l); w != 60 {
			t.Errorf("line %d width = %d, want 60", i, w)
		}
	}
}
