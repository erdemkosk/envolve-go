package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	var currentView string

	if !b.ready {
		return fmt.Sprintf("%s%s", b.loader.View(), "loading...")
	}

	if b.help.ShowAll {
		currentView = b.help.View(b.keys)
	} else {
		var s strings.Builder
		s.WriteString("\n  ")
		if b.SelectedFile == "" {
			s.WriteString("Pick a file:")
		} else {
			s.WriteString("Selected file: " + b.Filepicker.Styles.Selected.Render(b.SelectedFile))
		}
		s.WriteString("\n\n" + b.Filepicker.View() + "\n")
		return s.String()
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true).
		Italic(true).
		Render(currentView)
}
