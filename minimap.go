package minimap

import (
	"strings"
)

func Simple(contents string, targetLineLength, targetOutputLines int) string {
	var (
		lines        = strings.Split(strings.TrimSpace(contents), "\n")
		lenLines     = len(lines)
		batchCounter = 0
		lengthMap    = make(map[int]float64)
		maxLength    = 0
	)

	batchSize := 0
	if targetOutputLines > 1 {
		batchSize = lenLines / (targetOutputLines - 1)
	}

	if batchSize <= 0 {
		for i := 0; i < lenLines; i++ {
			line := lines[i]
			lenLine := len(line)

			if prev, ok := lengthMap[0]; !ok {
				lengthMap[0] = float64(lenLine)
			} else {
				lengthMap[0] = (prev + float64(lenLine)) / 2.0
			}

			if lenLine > maxLength {
				maxLength = lenLine
			}
		}
	} else {
		for i := 0; i < lenLines; i++ {
			line := lines[i]
			lenLine := len(line)

			if i%batchSize == 0 {
				batchCounter++
			}

			if prev, ok := lengthMap[batchCounter]; !ok {
				lengthMap[batchCounter] = float64(lenLine)
			} else {
				lengthMap[batchCounter] = (prev + float64(lenLine)) / 2.0
			}

			if lenLine > maxLength {
				maxLength = lenLine
			}
		}
	}

	// maxLength * ? = targetLineLength
	// ? = targetLineLength / maxLength
	scaleDown := float64(targetLineLength) / float64(maxLength)

	var sb strings.Builder
	for i := 0; i < len(lengthMap); i++ {
		lineLength := lengthMap[i]
		sb.WriteString(strings.Repeat("*", int(lineLength*scaleDown)) + "\n")
	}
	return strings.TrimSpace(sb.String())
}
