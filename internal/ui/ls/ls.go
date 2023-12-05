package ls

import (
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Filepicker   filepicker.Model
	SelectedFile string
	quitting     bool
	err          error
}

type clearErrorMsg struct{}

func clearErrorAfter(t time.Duration) tea.Cmd {
	return tea.Tick(t, func(_ time.Time) tea.Msg {
		return clearErrorMsg{}
	})
}

func (m Model) Init() tea.Cmd {
	return m.Filepicker.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		}
	case clearErrorMsg:
		m.err = nil
	}

	var cmd tea.Cmd
	m.Filepicker, cmd = m.Filepicker.Update(msg)

	if didSelect, path := m.Filepicker.DidSelectFile(msg); didSelect {
		// Dosya seçildiğinde içeriğini oku ve Model içine koy
		content, err := ioutil.ReadFile(path)
		if err != nil {
			m.err = err
			m.SelectedFile = ""
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		}
		m.SelectedFile = string(content)

	}

	if didSelect, path := m.Filepicker.DidSelectDisabledFile(msg); didSelect {
		m.err = errors.New(path + " is not valid.")
		m.SelectedFile = ""
		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
	}

	return m, cmd
}

func (m Model) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder

	if m.SelectedFile != "" {
		s.WriteString(m.SelectedFile)
	} else {
		// Dosya seçilmediyse file picker'ı göster
		s.WriteString("\n\n" + m.Filepicker.View() + "\n")
	}

	s.WriteString("\n  ")
	if m.err != nil {
		s.WriteString(m.Filepicker.Styles.DisabledFile.Render(m.err.Error()))
	} else if m.SelectedFile == "" {
		s.WriteString("Pick a file:")
	} else {
		s.WriteString("Selected file: " + m.Filepicker.Styles.Selected.Render(m.SelectedFile))
	}

	return s.String()
}

func NewModel(filePicker filepicker.Model) Model {

	return Model{
		Filepicker: filePicker,
	}
}
