package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/erdemkosk/envolve-go/internal/handler"

	"github.com/spf13/cobra"
)

var SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "sync bubbletea-starter to the latest version",
	Long:  `sync bubbletea-starter to the latest version.`,
	Run: func(cmd *cobra.Command, args []string) {
		envolvePath := handler.GetEnvolveHomePath()
		currentPath, currentFolderName := handler.GetCurrentPathAndFolder()
		targetPath := filepath.Join(envolvePath, currentFolderName)
		currentEnvFilePath := filepath.Join(currentPath, "/.env")
		targetEnvFilePath := filepath.Join(targetPath, "/.env")

		handler.CreateFolderIfDoesNotExist(targetPath)
		handler.CopyFile(currentEnvFilePath, targetEnvFilePath)
		handler.DeleteFile(currentEnvFilePath)
		val := handler.Symlink(targetEnvFilePath, currentEnvFilePath)
		if val != nil {
			fmt.Println("Sembolik link oluşturulamadı:", val)
			return
		}
		fmt.Println("Sync work with success!")

		os.Exit(0)
	},
}
