package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xyproto/minimap"
)

func main() {

	filename := ""
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if filename == "" {
		log.Println("please provide a filename as the first argument")
		os.Exit(1)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	s := minimap.Block(string(data), 40, 10)
	fmt.Println(s)
}
