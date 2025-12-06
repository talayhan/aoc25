package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	ranges, err := parseRanges(lines)
	if err != nil {
		fmt.Printf("Error parsing ranges: %v\n", err)
		return
	}

	totalCount := countUniqueIDs(ranges)

	fmt.Printf("Total unique good number IDs: %d\n", totalCount)
}

func countUniqueIDs(ranges []NumberRange) uint64 {
	if len(ranges) == 0 {
		return 0
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	var totalCount uint64 = 0

	currentStart := ranges[0].Start
	currentEnd := ranges[0].End

	for i := 1; i < len(ranges); i++ {
		nextStart := ranges[i].Start
		nextEnd := ranges[i].End

		if nextStart <= currentEnd+1 {
			if nextEnd > currentEnd {
				currentEnd = nextEnd
			}
		} else {
			totalCount += uint64(currentEnd - currentStart + 1)

			currentStart = nextStart
			currentEnd = nextEnd
		}
	}

	totalCount += uint64(currentEnd - currentStart + 1)

	return totalCount
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
		text := strings.TrimSpace(scanner.Text())
		if text != "" {
			lines = append(lines, text)
		}
	}
	return lines, scanner.Err()
}

func parseRanges(lines []string) ([]NumberRange, error) {
	var ranges []NumberRange

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
		}

		start, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		end, err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)

		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid number format in line: %s", line)
		}

		if start > end {
			start, end = end, start
		}

		ranges = append(ranges, NumberRange{Start: start, End: end})
	}

	return ranges, nil
}
