package ls

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/erdemkosk/envolve-go/internal/handler"
)

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)

type Model struct {
	Filepicker   filepicker.Model
	Viewport     viewport.Model
	SelectedFile string
	quitting     bool
	err          error
	ready        bool
	folderName   string
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
	var (
		cmd tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "j", "down":
			m.Viewport.YOffset = m.Viewport.YOffset + 1

		case "k", "up":
			if m.Viewport.YOffset > 0 {
				m.Viewport.YOffset = m.Viewport.YOffset - 1
			}
		}

	case clearErrorMsg:
		m.err = nil

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.Viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.Viewport.YPosition = headerHeight
			m.ready = true
			m.Viewport.YPosition = headerHeight + 1
		} else {
			m.Viewport.Width = msg.Width
			m.Viewport.Height = msg.Height - verticalMarginHeight
			m.Viewport, cmd = m.Viewport.Update(msg)

		}
	}

	m.Filepicker, cmd = m.Filepicker.Update(msg)

	if didSelect, path := m.Filepicker.DidSelectFile(msg); didSelect {
		m.folderName = handler.GetFoldername(path)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			m.err = err
			m.SelectedFile = ""
			return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
		}

		m.SelectedFile = stylizeLines(string(content))
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

	if m.err != nil {
		s.WriteString(m.Filepicker.Styles.DisabledFile.Render(m.err.Error()))
	}
	if m.SelectedFile != "" {
		s.WriteString("Selected file: " + m.Filepicker.Styles.Selected.Render(m.SelectedFile))
		m.Viewport.SetContent(m.SelectedFile)
		//s.WriteString("\n\n" + m.Viewport.View() + "\n")

		return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.Viewport.View(), m.footerView())
	} else {
		s.WriteString("Pick a file:")
		s.WriteString("\n\n" + m.Filepicker.View() + "\n")
	}

	return s.String()
}

func NewModel(filePicker filepicker.Model) Model {

	return Model{
		Filepicker: filePicker,
	}
}

//STYLE//

func (m Model) headerView() string {
	styled := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(m.folderName)
	title := titleStyle.Render(styled + "/.env")
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.Viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.Viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func stylizeLines(content string) string {
	lines := strings.Split(content, "\n")
	var sb strings.Builder

	for _, line := range lines {
		if strings.Contains(line, "=") {
			idx := strings.Index(line, "=")
			firstPart := line[:idx]
			secondPart := line[idx:]

			styledFirstPart := lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Render(firstPart)

			line = styledFirstPart + secondPart
		}
		sb.WriteString(line + "\n")
	}

	return sb.String()
}

//
