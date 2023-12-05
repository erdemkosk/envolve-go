package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/erdemkosk/envolve-go/internal/handler"
	"github.com/erdemkosk/envolve-go/internal/ui/ls"
	"github.com/spf13/cobra"
)

var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Backs up your current project's .env file",
	Long:  `Backs up your current project's .env file, restores the variables from a global .env file, and creates a symbolic link to the latest environment settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		fp := filepicker.New()
		fp.AllowedTypes = []string{".env"}
		fp.ShowHidden = true

		fp.CurrentDirectory = handler.GetEnvolveHomePath()

		m := ls.NewModel(fp)

		tm, _ := tea.NewProgram(&m, tea.WithOutput(os.Stderr)).Run()
		mm := tm.(ls.Model)

		fmt.Println("\n  You selected: " + m.Filepicker.Styles.Selected.Render(mm.SelectedFile) + "\n")

		os.Exit(0)
	},
}
