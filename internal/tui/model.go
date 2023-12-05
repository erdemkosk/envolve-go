package tui

import (
	"github.com/erdemkosk/envolve-go/internal/config"

	"github.com/charmbracelet/bubbles/filepicker"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// Bubble represents the state of the UI.
type Bubble struct {
	keys         keyMap
	help         help.Model
	loader       spinner.Model
	Filepicker   filepicker.Model
	SelectedFile string
	viewport     viewport.Model
	appConfig    config.Config
	ready        bool
}

// NewBubble creates an instance of the UI.
func NewBubble(cfg config.Config) Bubble {
	keys := getDefaultKeyMap()

	l := spinner.New()
	l.Spinner = spinner.Dot

	h := help.New()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	return Bubble{
		keys:         keys,
		help:         h,
		loader:       l,
		viewport:     viewport.Model{},
		Filepicker:   filepicker.Model{},
		SelectedFile: "",
		appConfig:    cfg,
		ready:        false,
	}
}

func CreateBubble(filePicker filepicker.Model) Bubble {
	keys := getDefaultKeyMap()

	l := spinner.New()
	l.Spinner = spinner.Dot

	h := help.New()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	return Bubble{
		keys:         keys,
		help:         h,
		loader:       l,
		viewport:     viewport.Model{},
		Filepicker:   filePicker,
		SelectedFile: "",
		ready:        false,
	}
}
