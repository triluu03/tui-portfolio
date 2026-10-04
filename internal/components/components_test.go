package components

import (
	"strings"
	"testing"

	lipgloss "charm.land/lipgloss/v2"
)

func TestHeaderWidth(t *testing.T) {
	tabs := []string{"About", "Projects", "Experience", "Contact"}
	for _, width := range []int{120, 60, 40, 20, 0} {
		got := Header(width, 1, tabs)
		for _, line := range strings.Split(got, "\n") {
			if w := lipgloss.Width(line); w != width {
				t.Errorf("Header(%d) line width = %d, want %d (%q)", width, w, width, line)
			}
		}
	}
}

func TestFooterWidth(t *testing.T) {
	for _, width := range []int{120, 60, 40, 20, 0} {
		got := Footer(width, "About")
		for _, line := range strings.Split(got, "\n") {
			if w := lipgloss.Width(line); w != width {
				t.Errorf("Footer(%d) line width = %d, want %d (%q)", width, w, width, line)
			}
		}
	}
}

func TestPalette(t *testing.T) {
	colors := map[string]string{
		"ColorBackground": ColorBackground,
		"ColorForeground": ColorForeground,
		"ColorDim":        ColorDim,
		"ColorAccent":     ColorAccent,
	}
	seen := map[string]bool{}
	for name, c := range colors {
		if !strings.HasPrefix(c, "#") || len(c) != 7 {
			t.Errorf("%s = %q, want a #RRGGBB hex color", name, c)
		}
		if seen[c] {
			t.Errorf("%s duplicates another palette color (%s)", name, c)
		}
		seen[c] = true
	}
}
