package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var archiveExtension = "zip"
var archiveMimeType = "application/zip"

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

	volumeName := os.Getenv("RAILWAY_VOLUME_NAME")
	fileName := fmt.Sprintf("%s.%s", volumeName, archiveExtension)

	fmt.Printf("Volume path: %s\n", mountPath)
	fmt.Printf("Volume name: %s\n", volumeName)

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	w.Header().Set("Content-Type", archiveMimeType)

	if err := compress(mountPath, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", download)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server running at http://localhost:%s\n", port)

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
