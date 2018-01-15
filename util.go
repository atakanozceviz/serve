package main

import (
	"net"
	"net/http"
	"os"
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
