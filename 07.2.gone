package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file := "07.input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	start := regexp.MustCompile(`S`)
	startLocation := start.FindStringIndex(lines[0])

	currentLine := make([]int, len(lines[0]))
	currentLine[startLocation[0]] = 1

	for _, lineString := range lines {
		line := strings.Split(lineString, "")
		for i := range currentLine {
			if line[i] == "^" {
				currentLine[i-1] += currentLine[i]
				currentLine[i+1] += currentLine[i]
				currentLine[i] = 0
			}
		}
	}

	timelines := 0
	for _, timeline := range currentLine {
		timelines += timeline
	}
	fmt.Println("Timelines: ", timelines)

}
