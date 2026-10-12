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

// Contact is the contact page.
type Contact struct {
	selectedContactLink int
	focus               *bool
}

// Init is a no-op placeholder.
func (c Contact) Init() tea.Cmd { return nil }

// Update is a no-op placeholder.
func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && (*c.focus) {
		switch key.String() {
		case "up", "k":
			c.selectedContactLink = (c.selectedContactLink + nContactLinks - 1) % nContactLinks
		case "down", "j":
			c.selectedContactLink = (c.selectedContactLink + 1) % nContactLinks
		case "enter":
			return c, openURL(contactLinks[c.selectedContactLink].url)
		}
	}

	return c, nil
}

// View renders the contact page.
func (c Contact) View() tea.View { return tea.NewView(c.renderContact()) }

// renderContact builds the two-column contact page padded to the pinned frame size.
func (c Contact) renderContact() string {
	compose := contactComposePanel(style.ContactComposeWidth, style.ContentHeight, !*c.focus)
	links := contactLinksPanel(style.ContactLinksWidth, style.ContentHeight, c.selectedContactLink, *c.focus)
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

// placeholderBox renders a single-line bordered box sized to width visible
// columns. The border uses the faint panel color and the text is dim.
func placeholderBox(width int, text string) string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(style.ColorFaint)).
		Width(width).
		Render(dimStyle.Render(text))
}
