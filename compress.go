package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/klauspost/pgzip"
)

func compress(volumePath string, w io.Writer) error {
	gzipWriter, gzipWriterErr := pgzip.NewWriterLevel(w, pgzip.BestSpeed)
	if gzipWriterErr != nil {
		return fmt.Errorf("could not create gzip writer: %v", gzipWriterErr)
	}
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

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

			header, err := tar.FileInfoHeader(info, relPath)
			if err != nil {
				return fmt.Errorf("could not create tar header for directory: %v", err)
			}
			header.Name = relPath

			if err := tarWriter.WriteHeader(header); err != nil {
				return fmt.Errorf("could not write tar directory header: %v", err)
			}

			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("could not open file: %v", err)
		}
		defer file.Close()

		header, err := tar.FileInfoHeader(info, relPath)
		if err != nil {
			return fmt.Errorf("could not create tar header: %v", err)
		}
		header.Name = relPath

		if err := tarWriter.WriteHeader(header); err != nil {
			return fmt.Errorf("could not write tar header: %v", err)
		}

		if _, err := io.Copy(tarWriter, file); err != nil {
			return fmt.Errorf("could not copy file contents: %v", err)
		}

		filesProcessed++
		fmt.Printf("\rFile processed %d/%d", filesProcessed, totalFiles)

		return nil
	})

	if walkErr != nil {
		return fmt.Errorf("error walking the file tree: %v", walkErr)
	}

	fmt.Println("\nTar file streamed successfully")

	return nil
}
