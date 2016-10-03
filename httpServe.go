package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/koding/websocketproxy"
)

func main() {

	cwd, err := os.Getwd() // Get the current working directory
	// quit if something goes wrong
	// and, don't do anything to resolve it
	// just quit for now

	if err != nil {
		log.Fatal(err)
	}

	version := "0.0.2"

	// Process the comman line options flag
	versionFlag := flag.Bool("version", false, "show version")
	directory := flag.String("dir", cwd, "directory to serve")
	port := flag.String("port", "8000", "port")
	flagBackend := flag.String("backend", "ws://localhost:8000", "Backend URL for proxying")

	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		os.Exit(0)
	}

	u, err := url.Parse(*flagBackend)
	if err != nil {
		log.Fatal(err)
	}

	// Set a startup message
	log.Println("Serving files from ", *directory)

	// serve the directory
	fs := http.FileServer(http.Dir(*directory))
	http.Handle("/", fs)

	log.Println("Serving on port", *port)
	http.ListenAndServe(":"+*port, websocketproxy.NewProxy(u))

}
