package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	cwd, err := os.Getwd() // Get the current working directory

	// quit if something goes wrong
	// and, don't do anything to resolve it
	// just quit for now

	if err != nil {
		log.Fatal(err)
	}

	// Set a startup message
	log.Println("Serving files from ", cwd)

	// serve the directory
	fs := http.FileServer(http.Dir(cwd))
	http.Handle("/", fs)

	log.Println("Serving on port 3000...")
	http.ListenAndServe(":3000", nil)

}
