package main

import (
	"io/ioutil"
	"log"
)

func main() {
	log.Println("Hello world")
	err := ioutil.WriteFile("filename.txt", []byte("Hello world"), 0755)
	if err != nil {
		log.Printf("Unable to write file: %v\n", err)
	}
}
