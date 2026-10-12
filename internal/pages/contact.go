package pages

import (
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"

	"github.com/triluu03/tui-portfolio/internal/components"
	"github.com/triluu03/tui-portfolio/internal/style"
)

// Contact is the contact page.
type Contact struct{ focus *bool }

// Init is a no-op placeholder.
func (c Contact) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return c, nil }

// View renders the contact page.
func (c Contact) View() tea.View { return tea.NewView(renderContact(*c.focus)) }

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

// renderContact builds the two-column contact page padded to the pinned frame size.
func renderContact(focusRight bool) string {
	compose := contactComposePanel(style.ContactComposeWidth, style.ContentHeight, !focusRight)
	links := contactLinksPanel(style.ContactLinksWidth, style.ContentHeight, focusRight)
	row := lipgloss.JoinHorizontal(lipgloss.Top, compose, strings.Repeat(" ", style.ColumnGap), links)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(row)
}

// contactComposePanel renders the mail compose panel.
func contactComposePanel(width, height int, focused bool) string {
	innerWidth := width - lipgloss.Width("│  │")
	promptLine := accentStyle.Render("$ mail " + recipientEmail)

	body := []string{"", promptLine, ""}
	for _, f := range [][2]string{
		{"name:", "Your Name"},
		{"email:", "you@example.com"},
		{"message:", "Hi! I'd like to talk about..."},
	} {
		body = append(body, dimStyle.Render(f[0]), placeholderBox(innerWidth, f[1]), "")
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
func contactLinksPanel(width, height int, focused bool) string {
	body := make([]string, 0, height-2)
	body = append(body, "")
	for _, l := range contactLinks {
		body = append(body,
			accentStyle.Render(l.label),
			dimStyle.Render(l.value+" ↗"),
			"",
		)
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

// placeholderBox renders a single-line bordered box sized to width visible
// columns. The border uses the faint panel color and the text is dim.
func placeholderBox(width int, text string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(style.ColorFaint)).
		Width(width).
		Render(dimStyle.Render(text))
}
