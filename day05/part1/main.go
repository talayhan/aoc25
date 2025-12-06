package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "input.txt"

type NumberRange struct {
	Start int64
	End   int64
}

func main() {
	lines, err := readLines(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	rangeLines, availableIDLines, err := separateData(lines)
	if err != nil {
		fmt.Printf("Error separating data: %v\n", err)
		return
	}

	ranges, err := parseRanges(rangeLines)
	if err != nil {
		fmt.Printf("Error parsing ranges: %v\n", err)
		return
	}

	count := countGoodAvailableIDs(availableIDLines, ranges)

	fmt.Printf("Total number of available good IDs: %d\n", count)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func separateData(lines []string) ([]string, []string, error) {
	separatorIndex := -1
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			separatorIndex = i
			break
		}
	}

	if separatorIndex == -1 {
		return nil, nil, fmt.Errorf("file format error: missing blank line separator")
	}

	rangeLines := make([]string, 0, separatorIndex)
	for i := 0; i < separatorIndex; i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed != "" {
			rangeLines = append(rangeLines, trimmed)
		}
	}

	availableIDLines := make([]string, 0, len(lines)-separatorIndex-1)
	for i := separatorIndex + 1; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed != "" {
			availableIDLines = append(availableIDLines, trimmed)
		}
	}

	return rangeLines, availableIDLines, nil
}

func parseRanges(lines []string) ([]NumberRange, error) {
	var ranges []NumberRange

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid range format: %s", line)
		}

		start, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			return nil, err
		}

		end, err := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
		if err != nil {
			return nil, err
		}

		ranges = append(ranges, NumberRange{Start: start, End: end})
	}

	return ranges, nil
}

func countGoodAvailableIDs(idLines []string, ranges []NumberRange) int {
	goodCount := 0

	for _, line := range idLines {
		id, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			fmt.Printf("Warning: Skipping invalid ID '%s'\n", line)
			continue
		}

		if isIDInRange(id, ranges) {
			goodCount++
		}
	}

	return goodCount
}

func isIDInRange(id int64, ranges []NumberRange) bool {
	for _, r := range ranges {
		if id >= r.Start && id <= r.End {
			return true
		}
	}
	return false
}
