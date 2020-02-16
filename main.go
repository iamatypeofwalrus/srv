package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port string
var dir string

func init() {
	flag.StringVar(&port, "port", "8080", "static file server port")
	flag.StringVar(&dir, "dir", ".", "directory to server")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, `srv: a simple static file server

Usage:`)

		flag.PrintDefaults()
	}

	flag.Parse()

	fileServer := http.FileServer(http.Dir(dir))

	fmt.Printf("serving %s on port %s\n", dir, port)

	log.Fatal(
		http.ListenAndServe(":"+port, fileServer),
	)
}
