package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define command-line flags
	var inputDir string
	flag.StringVar(&inputDir, "dir", ".", "Input directory to scan for audio files")
	flag.Parse()

	// Supported audio file extensions
	audioExtensions := []string{".mp3", ".wav", ".flac", ".ogg", ".aac", ".m4a"}

	// Validate the input directory
	if _, err := os.Stat(inputDir); os.IsNotExist(err) {
		fmt.Printf("Error: Directory '%s' does not exist.\n", inputDir)
		return
	}

	// Generate M3U files for the first-level directories
	if err := generateM3UForFirstLevelDirs(inputDir, audioExtensions); err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("M3U files generated successfully.")
	}
}

// generateM3UForFirstLevelDirs scans a directory and generates M3U files for first-level directories.
func generateM3UForFirstLevelDirs(rootDir string, extensions []string) error {
	dirs, err := os.ReadDir(rootDir)
	if err != nil {
		return fmt.Errorf("failed to read root directory: %w", err)
	}

	for _, entry := range dirs {
		if entry.IsDir() {
			subDirPath := filepath.Join(rootDir, entry.Name())
			m3uFilePath := filepath.Join(rootDir, entry.Name(), entry.Name()+".m3u")

			files, err := listFilesWithExtensions(subDirPath, extensions)
			if err != nil {
				return fmt.Errorf("failed to list files in %s: %w", subDirPath, err)
			}

			if len(files) > 0 {
				fmt.Printf("Generating M3U file for %s file count %v m3uFilePath %s...\n", subDirPath, len(files), m3uFilePath)
				if err := writeM3UFile(m3uFilePath, files); err != nil {
					return fmt.Errorf("failed to write M3U file %s: %w", m3uFilePath, err)
				}
				fmt.Printf("M3U file generated: %s\n", m3uFilePath)
			}
		}
	}

	return nil
}

// listFilesWithExtensions recursively lists files with specified extensions.
func listFilesWithExtensions(dir string, extensions []string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && hasValidExtension(path, extensions) {
			relPath, _ := filepath.Rel(dir, path) // Get relative path
			files = append(files, relPath)
		}
		return nil
	})

	return files, err
}

// hasValidExtension checks if a file has one of the specified extensions.
func hasValidExtension(filePath string, extensions []string) bool {
	ext := strings.ToLower(filepath.Ext(filePath))
	for _, validExt := range extensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

// writeM3UFile writes a list of files to an M3U file.
func writeM3UFile(m3uFilePath string, files []string) error {
	file, err := os.Create(m3uFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	file.WriteString("#EXTM3U\n")

	for _, f := range files {
		if _, err := file.WriteString(f + "\n"); err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}
	}

	return nil
}
