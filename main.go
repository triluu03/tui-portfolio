package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/triluu03/tui-portfolio/internal/app"
)

func main() {
	p := tea.NewProgram(app.New())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
