package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	lines := strings.Split(string(data), "\n")

	outputLines := 10

	batchSize := len(lines) / outputLines
	batchCounter := 0

	lengthMap := make(map[int]float64)

	if batchSize <= 0 {
		for _, line := range lines {
			if prev, ok := lengthMap[0]; !ok {
				lengthMap[0] = float64(len(line))
			} else {
				lengthMap[0] = (prev + float64(len(line))) / 2.0
			}
		}
	} else {
		for i, line := range lines {
			if i%batchSize == 0 {
				batchCounter++
			}

			if prev, ok := lengthMap[batchCounter]; !ok {
				lengthMap[batchCounter] = float64(len(line))
			} else {
				lengthMap[batchCounter] = (prev + float64(len(line))) / 2.0
			}
		}
	}

	for _, lineLength := range lengthMap {
		fmt.Println(strings.Repeat("*", int(lineLength)))
	}
}
