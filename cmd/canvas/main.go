package main

import (
	"log"
	"os"
	"strconv"

	"github.com/xyproto/minimap"
	"github.com/xyproto/mode"
	"github.com/xyproto/vt100"
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

	highlightIndex := 7
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

	// Initialize vt100 terminal settings
	vt100.Init()
	defer vt100.Close() // ensure that terminal settings are reset when we're done

	// Prepare a canvas
	c := vt100.NewCanvas()

	// Draw minimap onto the canvas
	var cw = int(c.Width())
	var ch = int(c.Height())

	const margin = 2

	const width = 20
	var height = ch - (margin * 2)

	// The x and y position for where the minimap should be drawn
	var xpos = cw - (width + margin)
	const ypos = margin

	err = minimap.DrawMinimap(c, string(data), xpos, ypos, width, height, fileMode, highlightIndex, vt100.DarkGray, vt100.Black, vt100.LightYellow, vt100.BackgroundBlack)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// Draw the contents of the canvas
	c.Draw()

	// Wait for a keypress
	vt100.WaitForKey()
}
