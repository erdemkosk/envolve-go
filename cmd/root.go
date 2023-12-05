package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/elewis787/boa"
	"github.com/erdemkosk/envolve-go/internal/config"
	"github.com/erdemkosk/envolve-go/internal/handler"
	"github.com/erdemkosk/envolve-go/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var b tui.Bubble

var rootCmd = &cobra.Command{
	Use:     "envolve",
	Short:   "envolve is a starting point for bubbletea apps",
	Version: "0.0.1",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ParseConfig()
		if err != nil {
			log.Fatal(err)
		}

		// If logging is enabled, logs will be output to debug.log.
		if cfg.Settings.EnableLogging {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			defer func() {
				if err = f.Close(); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}()
		}

		b := tui.NewBubble(cfg)
		var opts []tea.ProgramOption

		// Always append alt screen program option.
		opts = append(opts, tea.WithAltScreen())

		// If mousewheel is enabled, append it to the program options.
		if cfg.Settings.EnableMouseWheel {
			opts = append(opts, tea.WithMouseAllMotion())
		}

		// Initialize new app.
		p := tea.NewProgram(b, opts...)
		if err := p.Start(); err != nil {
			log.Fatal("Failed to start bubbletea-starter", err)
			os.Exit(1)
		}
	},
}

func loadInitialValues() {
	envolvePath := handler.GetEnvolveHomePath()
	handler.CreateFolderIfDoesNotExist(envolvePath)
}

// Execute executes the root command which starts the application.
func Execute() {
	loadInitialValues()
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(SyncCmd)
	rootCmd.AddCommand(LsCmd)

	rootCmd.SetUsageFunc(boa.UsageFunc)
	rootCmd.SetHelpFunc(boa.HelpFunc)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func RetriveBubbleModel() tui.Bubble {
	return b
}
