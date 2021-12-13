package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no.txt")
	if err != nil {
		log.Println(err)
	}
}