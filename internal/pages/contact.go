package pages

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
)

const (
	recipientEmail = "ductriluu.work@gmail.com"
	signatureName  = "Tri Luu"
	signatureRole  = "Data Scientist"
	signatureCity  = "Espoo, Finland"
)

// contactLink is a single entry in the links panel.
type contactLink struct {
	label string
	value string
	url   string
}

// contactLinks is the seed data. Entries are rendered in order; the url is
// reserved for a future interactive pass and is not rendered yet.
var contactLinks = []contactLink{
	{label: "email", value: "ductriluu.work@gmail.com", url: "mailto:ductriluu.work@gmail.com"},
	{label: "github", value: "triluu03", url: "https://github.com/triluu03"},
	{label: "linkedin", value: "triluu03", url: "https://linkedin.com/in/triluu03"},
	{label: "cv", value: "cv.pdf", url: "cv.pdf"},
}

var nContactLinks = len(contactLinks)

var nEditingFields = 3

// Contact is the contact page.
type Contact struct {
	focus    *bool
	editMode *bool

	selectedContactLink int

	editingField int

	// composeFields are: "name", "email", and "message" (in this order)
	composeFields [3]string
}

// Init is a no-op placeholder.
func (c Contact) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return c, nil
	}

	if *c.focus {
		switch key.String() {
		case "up", "k":
			c.selectedContactLink = (c.selectedContactLink + nContactLinks - 1) % nContactLinks
		case "down", "j":
			c.selectedContactLink = (c.selectedContactLink + 1) % nContactLinks
		case "enter":
			return c, openURL(contactLinks[c.selectedContactLink].url)
		}
	}

	if !*c.editMode && !*c.focus {
		switch key.String() {
		case "tab", "down", "j":
			*c.editMode = true
			c.editingField = 0
		}
		return c, nil
	}

	if *c.editMode {
		switch key.String() {
		case "ctrl+enter", "alt+enter":
			// clear all the fields without sending anything for now
			c.composeFields = [3]string{"", "", ""}
			c.editingField = 0
			*c.editMode = false
		case "up":
			c.editingField = (c.editingField + nEditingFields - 1) % nEditingFields
		case "tab", "down":
			c.editingField = (c.editingField + 1) % nEditingFields
		case "backspace":
			c.composeFields[c.editingField] = deleteLastChar(c.composeFields[c.editingField])
		default:
			c.composeFields[c.editingField] += key.Text
		}
		return c, nil
	}

	return c, nil

}

// View renders the contact page.
func (c Contact) View() tea.View { return tea.NewView(c.renderContact()) }

// renderContact builds the two-column contact page padded to the pinned frame size.
func (c Contact) renderContact() string {
	compose := contactComposePanel(
		style.ContactComposeWidth,
		style.ContentHeight,
		*c.editMode,
		c.editingField,
		c.composeFields,
		!*c.focus,
	)
	links := contactLinksPanel(style.ContactLinksWidth, style.ContentHeight, c.selectedContactLink, *c.focus)
	row := lipgloss.JoinHorizontal(lipgloss.Top, compose, strings.Repeat(" ", style.ColumnGap), links)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(row)
}

// contactComposePanel renders the mail compose panel.
func contactComposePanel(width, height int, editing bool, editingField int, composeFields [3]string, focused bool) string {
	innerWidth := width - lipgloss.Width("│  │")
	promptLine := accentStyle.Render("$ mail " + recipientEmail)

	body := []string{"", promptLine, ""}
	for i, fieldName := range [3]string{"name", "email", "message"} {
		cursor, fieldStyle, focus := "", dimStyle, false
		if editing && i == editingField {
			fieldStyle = accentStyle
			cursor = "█"
			focus = true
		}

		body = append(body,
			fieldStyle.Render(fieldName),
			composeFieldBox(innerWidth, fieldStyle.Render(composeFields[i]+cursor), focus),
			"",
		)
	}
	body = append(body, accentStyle.Render("[ send ⏎ ]")+" "+dimStyle.Render("ctrl+enter to send"))

	return components.Panel(
		"mail --compose",
		width,
		height,
		strings.Join(body, "\n"),
		focused,
	)
}

// contactLinksPanel renders the links panel with the signature pinned to the
// bottom: filler lines = panel inner height - links - signature lines.
func contactLinksPanel(width, height int, selected int, focused bool) string {
	body := make([]string, 0, height-2)
	body = append(body, "")
	for i, l := range contactLinks {
		if i == selected {
			body = append(body,
				dimStyle.Render(l.label),
				accentStyle.Render(l.value+" ↗"),
				"",
			)
		} else {
			body = append(body,
				dimStyle.Render(l.label),
				l.value,
				"",
			)
		}
	}

	signature := []string{
		warmStyle.Render(signatureName),
		dimStyle.Render(signatureRole + " · " + signatureCity),
	}
	if n := height - 2 - len(signature) - len(body); n > 0 {
		body = append(body, make([]string, n)...)
	}
	body = append(body, signature...)

	return components.Panel("links", width, height, strings.Join(body, "\n"), focused)
}

// composeFieldBox renders a single-line bordered box sized to width visible
// columns. The border uses the faint panel color and the text is dim.
func composeFieldBox(width int, text string, focused bool) string {
	borderStyle := style.ColorFaint
	if focused {
		borderStyle = style.ColorAccent
	}

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(borderStyle)).
		Width(width).
		Render(text)
}

func deleteLastChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}
