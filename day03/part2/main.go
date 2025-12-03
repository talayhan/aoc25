package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Some part of this function was written by Gemini 3 - Fast
func findBiggestSubsequence(s string) string {
	n := len(s)
	targetLength := 12

	if n <= targetLength {
		return s
	}

	// K is the number of digits we must remove
	k := n - targetLength

	// Use a slice of runes (stack) to build the result subsequence
	result := make([]rune, 0, targetLength)

	for _, digit := range s {
		// Greedy check: While the result has digits, and we still have removals (k > 0),
		// AND the current digit is greater than the last digit in the result:
		for len(result) > 0 && k > 0 && digit > result[len(result)-1] {
			// Remove the smaller preceding digit
			result = result[:len(result)-1] // Pop
			k--
		}
		// Add the current digit
		result = append(result, digit) // Push
	}

	// If k > 0 remaining removals (happens when digits are ascending, e.g., "12345")
	// Remove the remaining k digits from the end
	if k > 0 {
		result = result[:len(result)-k]
	}

	// The result slice might be longer than 12 if the original string was longer than 12
	// but less than the necessary removals. Ensure we return exactly the target length.
	return string(result[:targetLength])
}

func main() {
	const filename = "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open the file %s: %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalVoltage := 0

	for scanner.Scan() {
		line := scanner.Text()

		length := len(line)
		fmt.Println("line:", line)
		fmt.Println("len:", length)

		resultStr := findBiggestSubsequence(line)
		fmt.Printf("Input: %s -> Output: %s\n", line, resultStr)
		result, _ := strconv.Atoi(resultStr)
		totalVoltage += result
	}

	fmt.Println("Total voltage: ", totalVoltage)

	if err := scanner.Err(); err != nil {
		fmt.Errorf("file scanning failed: %w", err)
	}
}
