package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	const filename = "./input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open the file %s: %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 0

	startPos := 50
	zeroCount := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		direction := line[0] // 'L' or 'R'

		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			fmt.Printf("Couldn't parse distance %s as integer %v\n", distanceStr, err)
			continue
		}

		for i := 0; i < distance; i++ {
			if direction == 'L' {
				startPos--
			} else if direction == 'R' {
				startPos++
			}

			if startPos%100 == 0 {
				zeroCount++
			}
		}
	}
	fmt.Println("Total zero count: ", zeroCount)

	if err := scanner.Err(); err != nil {
		fmt.Errorf("file scanning failed: %w", err)
	}
}
