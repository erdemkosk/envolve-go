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
	Short: "Backs up your current project's .env file",
	Long:  `Backs up your current project's .env file, restores the variables from a global .env file, and creates a symbolic link to the latest environment settings.`,
	Run: func(cmd *cobra.Command, args []string) {
		envolvePath := handler.GetEnvolveHomePath()
		currentPath, currentFolderName := handler.GetCurrentPathAndFolder()
		targetPath := filepath.Join(envolvePath, currentFolderName)
		currentEnvFilePath := filepath.Join(currentPath, "/.env")
		targetEnvFilePath := filepath.Join(targetPath, "/.env")

		handler.CreateFolderIfDoesNotExist(targetPath)
		handler.CopyFile(currentEnvFilePath, targetEnvFilePath)
		handler.DeleteFile(currentEnvFilePath)
		handler.Symlink(targetEnvFilePath, currentEnvFilePath)

		fmt.Println("Sync work with success!")

		os.Exit(0)
	},
}
