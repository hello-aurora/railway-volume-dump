package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zip"
)

func compress(volumePath string, w io.Writer) error {
	zipWriter := zip.NewWriter(w)
	defer zipWriter.Close()

	totalFiles := 0
	filesProcessed := 0

	_ = filepath.Walk(volumePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			totalFiles++
		}

		return nil
	})

	walkErr := filepath.Walk(volumePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(volumePath, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			relPath += "/"

			_, err := zipWriter.Create(relPath)
			if err != nil {
				return fmt.Errorf("could not create directory in ZIP: %v", err)
			}

			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("could not open file: %v", err)
		}
		defer file.Close()

		zipHeader, err := zip.FileInfoHeader(info)
		if err != nil {
			return fmt.Errorf("could not create ZIP header: %v", err)
		}
		zipHeader.Name = relPath
		zipHeader.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(zipHeader)
		if err != nil {
			return fmt.Errorf("could not create ZIP writer: %v", err)
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			return fmt.Errorf("could not copy file contents: %v", err)
		}

		filesProcessed++
		fmt.Printf("\rFile processed %d/%d", filesProcessed, totalFiles)

		return nil
	})

	if walkErr != nil {
		return fmt.Errorf("error walking the file tree: %v", walkErr)
	}

	fmt.Println("\nZIP file streamed successfully")

	return nil
}
