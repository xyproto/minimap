package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xyproto/minimap"
	"github.com/xyproto/mode"
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

	// Use mode.SimpleDetect to detect the mode based on the file contents
	fileMode := mode.SimpleDetectBytes(data)

	// Generate colored minimap
	const width = 20
	const height = 40

	s, err := minimap.ColorMinimap(string(data), width, height, fileMode, 5)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
