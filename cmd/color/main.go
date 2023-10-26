package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

	highlightIndex := 10
	if len(os.Args) > 2 {
		highlightString := os.Args[2]
		if i, err := strconv.Atoi(highlightString); err == nil {
			highlightIndex = i
		}
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

	s, err := minimap.ColorMinimap(string(data), width, height, fileMode, highlightIndex)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(s)
}
