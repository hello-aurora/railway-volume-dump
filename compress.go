package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/compress/zip"
)

func compress(volumePath, outputPath string) error {
	outFile, outFileErr := os.Create(outputPath)
	if outFileErr != nil {
		return fmt.Errorf("could not create zip file: %v", outFileErr)
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
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

	outFileErr = filepath.Walk(volumePath, func(path string, info os.FileInfo, err error) error {
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
				return fmt.Errorf("could not create directory in zip: %v", err)
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
			return fmt.Errorf("could not create zip header: %v", err)
		}
		zipHeader.Name = relPath
		zipHeader.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(zipHeader)
		if err != nil {
			return fmt.Errorf("could not create zip writer: %v", err)
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			return fmt.Errorf("could not copy file contents: %v", err)
		}

		filesProcessed++
		fmt.Printf("\rFile processed %d/%d", filesProcessed, totalFiles)

		return nil
	})

	if outFileErr != nil {
		return fmt.Errorf("error walking the file tree: %v", outFileErr)
	}

	outFileErr = zipWriter.Close()
	if outFileErr != nil {
		return fmt.Errorf("could not finalize zip file: %v", outFileErr)
	}

	fmt.Println("\nZip file created successfully")

	return nil
}
