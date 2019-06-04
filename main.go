package main

import (
	"io/ioutil"
	"log"
)

const (
	filename = "scratchfile.txt"
)

func main() {
	log.Printf("Creating file: %q\n", filename)
	err := ioutil.WriteFile(filename, []byte("Hello world"), 0755)
	if err != nil {
		log.Printf("Unable to write file: %v\n", err)
	}
	log.Printf("Created file: %q\n", filename)
}
