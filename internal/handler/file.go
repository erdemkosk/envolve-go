package handler

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func getHomePath() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetEnvolveHomePath() string {
	homePath := getHomePath()
	envolvePath := filepath.Join(homePath, ".envolve-go")

	return envolvePath
}

func GetCurrentPathAndFolder() (string, string) {
	path, _ := os.Getwd()
	folder := filepath.Base(path)
	return path, folder
}

func CreateFolderIfDoesNotExist(homePath string) error {
	err := os.Mkdir(homePath, 0755)

	return err
}

func Symlink(source string, target string) error {
	err := os.Symlink(source, target)

	return err
}

func CopyFile(sourceFilePath string, targetFilePath string) {
	sourceFile, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Println("Source file problem", err)
		return
	}
	defer sourceFile.Close()

	targetFile, err := os.Create(targetFilePath)
	if err != nil {
		fmt.Println("Target file problem", err)
		return
	}
	defer targetFile.Close()

	_, err = sourceFile.Seek(0, 0)
	if err != nil {
		fmt.Println("Seek error", err)
		return
	}

	_, err = io.Copy(targetFile, sourceFile)
	if err != nil {
		fmt.Println("Dosya kopyalanamadı:", err)
		return
	}
}

func DeleteFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Remove problem", err)
		return
	}

}
