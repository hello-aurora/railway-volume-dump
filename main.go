package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func download(w http.ResponseWriter, r *http.Request) {
	mountPath := os.Getenv("RAILWAY_VOLUME_MOUNT_PATH")
	if mountPath == "" {
		http.Error(w, "No volume mounted to this service, please mount a volume first.", http.StatusInternalServerError)
		return
	}

	password := r.Header.Get("password")
	if password == "" || password != os.Getenv("PASSWORD") {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	dir, dirErr := os.Getwd()
	if dirErr != nil {
		http.Error(w, dirErr.Error(), http.StatusInternalServerError)
		return
	}

	volumeName := os.Getenv("RAILWAY_VOLUME_NAME")
	fileName := fmt.Sprintf("%s.zip", volumeName)
	zipPath := filepath.Join(dir, fileName)

	fmt.Printf("Volume path: %s\n", mountPath)
	fmt.Printf("Volume name: %s\n", volumeName)
	fmt.Printf("Zip path: %s\n", zipPath)

	fileErr := zipify(mountPath, zipPath)
	if fileErr != nil {
		http.Error(w, fileErr.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", "application/zip")

	http.ServeFile(w, r, zipPath)
}

func main() {
	//// TEMP
	// os.Setenv("RAILWAY_VOLUME_MOUNT_PATH", "./test-data")
	// os.Setenv("RAILWAY_VOLUME_NAME", "test-data")
	// os.Setenv("PASSWORD", "temp")
	//// TEMP

	http.HandleFunc("/", download)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server running at http://localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
