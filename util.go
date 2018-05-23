package main

import (
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

var server = &http.Server{
	ReadTimeout:  60 * time.Second,
	WriteTimeout: 60 * time.Second,
	IdleTimeout:  60 * time.Second,
}

// Get preferred outbound ip of this machine
func outboundIP() (net.IP, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP, nil
}

// Check if path is a directory
func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return fileInfo.IsDir(), err
}

func fileServer(directory string) http.HandlerFunc {
	fs := http.FileServer(http.Dir(directory))
	return func(w http.ResponseWriter, r *http.Request) {
		// Set some header.
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		// Serve with the actual handler.
		fs.ServeHTTP(w, r)
	}
}

func serveFile(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate") // HTTP 1.1.
		w.Header().Set("Pragma", "no-cache")                                   // HTTP 1.0.
		w.Header().Set("Expires", "0")                                         // Proxies.
		w.Header().Set("Content-Disposition", "attachment; filename=\""+path.Base(directory)+"")
		http.ServeFile(w, r, directory)
	}
}
