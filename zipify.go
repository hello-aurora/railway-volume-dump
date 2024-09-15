package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func zipify(volumePath, outputPath string) error {
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
