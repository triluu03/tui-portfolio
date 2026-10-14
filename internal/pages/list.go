package pages

import "strings"

// listEntry is one selectable item in a marker list. Lines[0] is the title,
// the rest are indented sub-lines.
type listEntry struct {
	lines []string
}

// listBody renders marker-list entries separated by blank lines. The selected
// entry gets "> " + accent; the rest get "  " + dim.
func listBody(entries []listEntry, selected int) string {
	rows := []string{""}
	for i, e := range entries {
		marker, lineStyle := "  ", dimStyle
		if i == selected {
			marker, lineStyle = "> ", accentStyle
		}
		rows = append(rows, lineStyle.Render(marker+e.lines[0]))
		for _, l := range e.lines[1:] {
			rows = append(rows, lineStyle.Render("  "+l))
		}
		if i < len(entries)-1 {
			rows = append(rows, "")
		}
	}
	return strings.Join(rows, "\n")
}

// keyValueRow is one label/value row. decoration is appended to the value
// only when the row is selected.
type keyValueRow struct {
	label      string
	value      string
	decoration string
}

// keyValueBody returns the body as a line slice (callers append a footer and
// fill to the panel height, so they need the slice, not a joined string).
func keyValueBody(rows []keyValueRow, selected int) []string {
	body := []string{""}
	for i, r := range rows {
		if i == selected {
			body = append(body, dimStyle.Render(r.label), accentStyle.Render(r.value+r.decoration), "")
		} else {
			body = append(body, dimStyle.Render(r.label), r.value, "")
		}
	}
	return body
}
