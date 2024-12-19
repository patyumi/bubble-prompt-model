package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView
	bodyView
)

type model struct {
	state     uint
	store     *Store
	notes     []Note
	currNote  Note
	listIndex int
	textArea  textarea.Model
	textInput textinput.Model

	// textarea.Model
	// ...
}

func NewModel(store *Store) model {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatalf("model: can't get notes %v", err)
	}

	return model{
		state:     listView,
		store:     store,
		notes:     notes,
		textArea:  textarea.New(),
		textInput: textinput.New(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.textArea, cmd = m.textArea.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String() // up, down, ctrl+c
		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit

			case "n":
				m.textInput.SetValue("")
				m.textInput.Focus()
				m.currNote = Note{}
				m.state = titleView
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex < len(m.notes)-1 {
					m.listIndex++
				}
			case "enter":
				m.currNote = m.notes[m.listIndex]
				m.textArea.SetValue(m.currNote.Body)
				m.textArea.Focus()
				m.textArea.CursorEnd()
				m.state = bodyView
				// show textArea
			}
		}
	}

	return m, tea.Batch(cmds...)
}
