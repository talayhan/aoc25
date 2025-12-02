package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func isRepeatedSequence(n int) bool {
	s := strconv.Itoa(n)
	totalLen := len(s)

	for i := 1; i <= totalLen/2; i++ {
		if totalLen%i == 0 {
			pattern := s[:i]
			repeats := totalLen / i
			reconstructed := strings.Repeat(pattern, repeats)

			if reconstructed == s {
				return true
			}
		}
	}

	return false
}

func main() {
	const filename = "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open the file %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	rangesRaw := strings.Split(line, ",")

	// @NOTE: debug logs
	fmt.Println("line: ", line)
	fmt.Println("rangeRaw: ", rangesRaw)

	var parsedRanges []Range

	sumInvalidIds := 0

	for _, r := range rangesRaw {

		before, after, found := strings.Cut(r, "-")
		if !found {
			fmt.Printf("Invalid format for range: %s\n", r)
			continue
		}

		startNum, err1 := strconv.Atoi(before)
		endNum, err2 := strconv.Atoi(after)

		if err1 != nil || err2 != nil {
			fmt.Println("Error converting numbers in range %s", r)
			continue
		}

		currentRange := Range{
			Start: startNum,
			End:   endNum,
		}

		parsedRanges = append(parsedRanges, currentRange)
		fmt.Printf("Processed: Start=%d, End=%d\n", startNum, endNum)

		for i := startNum; i <= endNum; i++ {
			if isRepeatedSequence(i) {
				fmt.Printf("InvalidId:%d\n", i)
				sumInvalidIds += i
			}
		}
	}
	fmt.Printf("Sum invalid IDs:%d\n", sumInvalidIds)
}
