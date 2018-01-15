package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	port := flag.String("p", "8080", "port to serve on")
	directory := flag.String("d", ".", "the directory or file to host")
	flag.Parse()

	args := os.Args[1:]
	if len(args) == 1 {
		*directory = args[0]
	}

	ok, err := isDirectory(*directory)
	switch {
	case err != nil:
		log.Fatal(err)
	case ok:
		http.Handle("/", http.FileServer(http.Dir(*directory)))
	case !ok:
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, *directory)
		})
	}

	ip, err := outboundIP()
	if err != nil {
		server.Addr = "127.0.0.1" + ":" + *port
		fmt.Printf("[WARN] Cannot get outbound IP: %v\n", err)
		fmt.Printf("[INFO] Serving \"%s\" on: http://%s\n", path.Base(*directory), server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("[FATAL] %v", err)
		}
	} else {
		server.Addr = ip.String() + ":" + *port
		fmt.Printf("[INFO] Serving \"%s\" on: http://%s\n", path.Base(*directory), server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("[FATAL] %v", err)
		}
	}
}
