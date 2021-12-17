package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	t, err := os.Open("no.txt")
	if err != nil {
		log.Println(err)
	}
	defer t.Close()
}
