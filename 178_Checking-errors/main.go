package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	
	r := strings.NewReader("Mak'e")
	
	io.Copy(f, r)
	
	o, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	
	bs, err := ioutil.ReadAll(o)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}