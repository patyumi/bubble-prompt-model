package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)

	faintStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)

	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render("NOTES APP") + "\n\n"

	if m.state == listView {

		for i, n := range m.notes {
			prefix := " "
			if i == m.listIndex {
				prefix = ">"
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += enumeratorStyle.Render(prefix) + n.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}

		s += faintStyle.Render("n - new note, q - quit")
	}

	return s
}
