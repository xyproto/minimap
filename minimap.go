package minimap

import (
	"strings"
)

// Simple function creates a basic minimap of the contents.
func Simple(contents string, targetLineLength, targetOutputLines int) string {
	if targetOutputLines == 0 {
		return ""
	}

	lines := strings.Split(contents, "\n")
	lenLines := len(lines)
	if lenLines == 0 {
		return ""
	}

	batchSize := lenLines / targetOutputLines
	remainder := lenLines % targetOutputLines
	if batchSize == 0 {
		return strings.Repeat("\n", targetOutputLines-1) // only empty lines
	}

	lineSums := make([]float64, targetOutputLines)
	maxBatchAverage := 0.0

	for i, line := range lines {
		batchIndex := i / batchSize
		if batchIndex >= targetOutputLines {
			batchIndex = targetOutputLines - 1 // for remainders
		}

		lineSums[batchIndex] += float64(len(line))
	}

	for i, sum := range lineSums {
		divider := batchSize
		if i == targetOutputLines-1 && remainder != 0 {
			divider = remainder
		}
		average := sum / float64(divider)
		lineSums[i] = average

		if average > maxBatchAverage {
			maxBatchAverage = average
		}
	}

	scaleDown := float64(targetLineLength) / maxBatchAverage

	var sb strings.Builder
	for _, avg := range lineSums {
		sb.WriteString(strings.Repeat("*", int(avg*scaleDown)) + "\n")
	}
	return strings.TrimSpace(sb.String())
}

// Dual function creates a detailed minimap using the ".:'" technique.
func Dual(contents string, targetLineLength, targetOutputLines int) string {
	// If targetOutputLines is 0, return an empty string
	if targetOutputLines == 0 {
		return ""
	}

	intermediateMap := Simple(contents, targetLineLength, 2*targetOutputLines)
	lines := strings.Split(intermediateMap, "\n")

	var sb strings.Builder
	for i := 0; i < len(lines); i += 2 {
		upper := lines[i]
		lower := ""
		if i+1 < len(lines) {
			lower = lines[i+1]
		}

		for j := 0; j < targetLineLength; j++ {
			if j < len(upper) && upper[j] == '*' && (j >= len(lower) || lower[j] != '*') {
				sb.WriteByte('\'')
			} else if j < len(lower) && lower[j] == '*' && (j >= len(upper) || upper[j] != '*') {
				sb.WriteByte('.')
			} else if j < len(upper) && upper[j] == '*' && j < len(lower) && lower[j] == '*' {
				sb.WriteByte(':')
			} else {
				sb.WriteByte(' ')
			}
		}
		sb.WriteByte('\n')
	}
	return strings.TrimSpace(sb.String())
}
