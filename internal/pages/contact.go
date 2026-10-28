package pages

import (
	"context"
	"fmt"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	"github.com/mymmrac/telego"

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
	{label: "cv", value: "cv.pdf", url: "https://www.triluu.dev/files/cv.pdf"},
}

var nContactLinks = len(contactLinks)

const nEditingFields = 2

// Compose field indices for [2]string{"email", "message"}.
const (
	emailField = iota
	messageField
)

const (
	// maxEmailRunes bounds the untrusted email field so it stays within one line.
	maxEmailRunes = style.ContactComposeWidth - 4

	statusSent = "sent"
)

// Contact is the contact page.
type Contact struct {
	focus    *bool
	editMode *bool

	selectedContactLink int

	editingField int

	// composeFields are: "email" and "message" (in this order)
	composeFields [2]string

	// sending is true while an async Telegram send is in flight.
	sending bool
	// status is shown under the send hint ("sent", an error, or empty).
	status string
	// copied is true while the transient "copied" confirmation is shown on
	// the links panel.
	copied bool
}

// Init is a no-op placeholder.
func (c Contact) Init() tea.Cmd { return nil }

// Update handles key presses and the result of an async Telegram send.
func (c Contact) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(copyFeedbackMsg); ok {
		c.copied = false
		return c, nil
	}

	if result, ok := msg.(SendResultMsg); ok {
		c.sending = false
		if result.err != nil {
			c.status = result.err.Error()
			return c, nil
		}
		c.status = statusSent
		c.composeFields = [2]string{"", ""}
		c.editingField = 0
		*c.editMode = false
		return c, nil
	}

	key, ok := msg.(tea.KeyPressMsg)
	if !ok {
		return c, nil
	}

	if *c.focus {
		switch key.String() {
		case "up", "k":
			c.selectedContactLink = (c.selectedContactLink + nContactLinks - 1) % nContactLinks
			c.copied = false
		case "down", "j":
			c.selectedContactLink = (c.selectedContactLink + 1) % nContactLinks
			c.copied = false
		case "enter":
			c.copied = true
			return c, copyCmd(contactLinks[c.selectedContactLink].url)
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
		if c.sending {
			return c, nil
		}

		switch key.String() {
		case "ctrl+enter", "alt+enter":
			fields, err := validateComposeFields(c.composeFields)
			if err != nil {
				c.status = err.Error()
				return c, nil
			}
			c.composeFields = fields
			c.sending = true
			c.status = ""

			text := buildTelegramMessage(fields)
			return c, func() tea.Msg {
				return SendResultMsg{err: sendTelegramMessage(text)}
			}
		case "up":
			c.editingField = (c.editingField + nEditingFields - 1) % nEditingFields
			c.status = ""
		case "tab", "down":
			c.editingField = (c.editingField + 1) % nEditingFields
			c.status = ""
		case "backspace":
			c.composeFields[c.editingField] = deleteLastChar(c.composeFields[c.editingField])
			c.status = ""
		case "enter":
			if c.editingField == messageField {
				c.composeFields[c.editingField] += "\n"
				c.status = ""
			}
		default:
			c.composeFields[c.editingField] += key.Text
			c.status = ""
		}
		return c, nil
	}

	return c, nil

}

// View renders the contact page.
func (c Contact) View() tea.View { return tea.NewView(c.renderContact()) }

// renderContact builds the two-column contact page padded to the pinned frame size.
func (c Contact) renderContact() string {
	compose := contactComposePanel(style.ContactComposeWidth, style.ContentHeight, c)
	links := contactLinksPanel(style.ContactLinksWidth, style.ContentHeight, c.selectedContactLink, *c.focus, c.copied)
	row := lipgloss.JoinHorizontal(lipgloss.Top, compose, strings.Repeat(" ", style.ColumnGap), links)
	return lipgloss.NewStyle().Padding(0, style.ContentPadX).Render(row)
}

// contactComposePanel renders the mail compose panel.
func contactComposePanel(width, height int, c Contact) string {
	innerWidth := width - lipgloss.Width("│  │")
	promptLine := accentStyle.Render("$ mail " + recipientEmail)

	body := []string{"", promptLine, ""}
	for i, fieldName := range [2]string{"email", "message"} {
		cursor, fieldStyle, focus := "", dimStyle, false
		if *c.editMode && i == c.editingField {
			fieldStyle = accentStyle
			cursor = "█"
			focus = true
		}

		innerHeight := 1
		if fieldName == "message" {
			innerHeight = 16
		}

		body = append(body,
			fieldStyle.Render(fieldName),
			composeFieldBox(
				innerWidth, innerHeight,
				fieldStyle.Render(c.composeFields[i]+cursor),
				focus,
			),
			"",
		)
	}
	body = append(body, accentStyle.Render("[ send ⏎ ]")+" "+dimStyle.Render("ctrl+enter or alt+enter to send"))

	if c.sending {
		body = append(body, dimStyle.Render("sending..."))
	} else if c.status != "" {
		statusStyle := errorStyle
		if c.status == statusSent {
			statusStyle = accentStyle
		}
		body = append(body, statusStyle.Render(c.status))
	}

	if *c.editMode {
		body = append(body, warmStyle.Render("[EDITING MODE] (esc to exit)"))
	}

	return components.Panel(
		"mail --compose",
		width,
		height,
		strings.Join(body, "\n"),
		!*c.focus,
	)
}

// contactLinksPanel renders the links panel with the signature pinned to the
// bottom: filler lines = panel inner height - links - signature lines.
func contactLinksPanel(width, height int, selected int, focused bool, copied bool) string {
	body := make([]string, 0, height-2)
	kv := make([]keyValueRow, len(contactLinks))
	for i, l := range contactLinks {
		decoration := " ↗"
		if copied {
			decoration = " ✓"
		}
		kv[i] = keyValueRow{label: l.label, value: l.value, url: l.url, decoration: decoration}
	}
	body = append(body, keyValueBody(kv, selected)...)

	signature := []string{
		warmStyle.Render(signatureName),
		dimStyle.Render(signatureRole + " · " + signatureCity),
	}
	if n := height - 2 - len(signature) - len(body); n > 0 {
		body = append(body, make([]string, n)...)
	}
	body = append(body, signature...)

	return components.Panel("links (enter to copy)", width, height, strings.Join(body, "\n"), focused)
}

// composeFieldBox renders a bordered box sized to width visible columns,
// wrapping text onto additional lines. The border uses the faint panel color
// and the text is dim.
func composeFieldBox(width, height int, text string, focused bool) string {
	borderStyle := style.ColorFaint
	if focused {
		borderStyle = style.ColorAccent
	}

	// The border consumes 2 columns; wrap so long lines flow onto additional
	// lines instead of overflowing the box.
	text = lipgloss.Wrap(text, width-2, " ")

	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color(borderStyle)).
		Width(width).
		Height(height).
		Render(text)
}

func deleteLastChar(s string) string {
	r := []rune(s)
	if len(r) == 0 {
		return s
	}
	return string(r[:len(r)-1])
}

// SendResultMsg reports the outcome of an async Telegram send. It is exported
// so the app router can deliver it to the contact page even when another page
// is active.
type SendResultMsg struct {
	err error
}

// validateComposeFields trims and validates untrusted compose input and
// returns the trimmed fields so the UI reflects what will be sent.
func validateComposeFields(fields [2]string) ([2]string, error) {
	for i := range fields {
		fields[i] = strings.TrimSpace(fields[i])
	}
	switch {
	case fields[emailField] == "":
		return fields, fmt.Errorf("email is required")
	case fields[messageField] == "":
		return fields, fmt.Errorf("message is required")
	case len([]rune(fields[emailField])) > maxEmailRunes:
		return fields, fmt.Errorf("email is too long")
	case !isValidEmail(fields[emailField]):
		return fields, fmt.Errorf("email is invalid")
	}
	return fields, nil
}

// isValidEmail reports whether s is a single, reply-able email address.
func isValidEmail(s string) bool {
	addr, err := mail.ParseAddress(s)
	if err != nil || addr.Address != s {
		return false
	}
	at := strings.IndexByte(s, '@')
	return at > 0 && strings.IndexByte(s[at+1:], '.') > 0
}

// buildTelegramMessage formats the message, truncating the body so the final
// text stays under Telegram's limit with margin.
func buildTelegramMessage(fields [2]string) string {
	header := "New contact message\n\nEmail: " + fields[emailField] + "\n\n"
	message := fields[messageField]
	return header + message
}

// sendTelegramMessage posts text to the configured Telegram chat via telego. Both
// TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID must be set; missing configuration
// is reported as an error that the UI can surface.
func sendTelegramMessage(text string) error {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")
	if token == "" || chatID == "" {
		return fmt.Errorf("send not configured: set TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID")
	}
	id, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		return fmt.Errorf("send not configured: TELEGRAM_CHAT_ID must be a numeric chat ID")
	}

	bot, err := telego.NewBot(token, telego.WithDiscardLogger())
	if err != nil {
		return fmt.Errorf("failed to send message")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = bot.SendMessage(ctx, &telego.SendMessageParams{
		ChatID: telego.ChatID{ID: id},
		Text:   text,
	})
	if err != nil {
		return fmt.Errorf("failed to send message")
	}
	return nil
}
