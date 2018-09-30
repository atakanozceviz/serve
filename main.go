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
	port := flag.String("p", "4242", "port on which the server will listen")
	directory := flag.String("s", ".", "filesystem path to read file(s) relative from")
	flag.Parse()

	args := os.Args[1:]
	if len(args) == 1 {
		*directory = args[0]
	}

	ok, err := isDirectory(*directory)
	switch {
	case err != nil:
		log.Fatalf("[FATAL] %v\n", err)
	case ok:
		http.HandleFunc("/", fileServer(*directory))
	case !ok:
		http.HandleFunc("/", serveFile(*directory))
	}

	http.HandleFunc("/u", uploadFile(*directory))

	ip, err := outboundIP()
	if err != nil {
		server.Addr = "127.0.0.1" + ":" + *port
		fmt.Printf("[WARN] Cannot get outbound IP: %v\n", err)
		fmt.Printf("[INFO] Serving \"%s\" on: http://%s\n", path.Base(*directory), server.Addr)
		fmt.Printf("[INFO] Use http://%s/u to upload files\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("[FATAL] %v\n", err)
		}
	} else {
		server.Addr = ip.String() + ":" + *port
		fmt.Printf("[INFO] Serving \"%s\" on: http://%s\n", path.Base(*directory), server.Addr)
		fmt.Printf("[INFO] Use http://%s/u to upload files\n", server.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("[FATAL] %v\n", err)
		}
	}
}
