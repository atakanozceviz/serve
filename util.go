package main

import (
	"fmt"
	"io"
	"log"
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

var form = `<!DOCTYPE html><html lang="en"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"> <title>Upload</title></head><body><h5>%s</h5> <form enctype="multipart/form-data" action="/u" method="post"> <input type="file" name="files" multiple/> <input type="submit" value="upload"/> </form></body></html>`

func uploadFile(directory string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintf(w, form, "")
		} else {
			err := r.ParseMultipartForm(32 << 20)
			if err != nil {
				fmt.Fprintf(w, "%v\n", err)
				log.Printf("[WARN] [%s] %v\n", r.RemoteAddr, err)
				return
			}
			files := r.MultipartForm.File["files"]
			for _, file := range files {
				src, err := file.Open()
				if err != nil {
					fmt.Fprintf(w, "%v\n", err)
					log.Printf("[WARN] [%s] %v\n", r.RemoteAddr, err)
					return
				}
				defer src.Close()

				dst, err := os.Create(path.Join(directory, file.Filename))
				if err != nil {
					fmt.Fprintf(w, "%v\n", err)
					log.Printf("[WARN] [%s] %v\n", r.RemoteAddr, err)
					return
				}
				defer dst.Close()
				if _, err = io.Copy(dst, src); err != nil {
					fmt.Fprintf(w, "%v\n", err)
					log.Printf("[WARN] [%s] %v\n", r.RemoteAddr, err)
					return
				}
			}
			fmt.Fprintf(w, form, fmt.Sprintf("Uploaded successfully %d files", len(files)))
		}
	}
}
