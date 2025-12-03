package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func findBiggestDigit(s string) (int, int) {
	var maxDigit rune
	maxIndex := -1

	for index, char := range s {
		if unicode.IsDigit(char) {
			if char > maxDigit {
				maxDigit = char
				maxIndex = index
			}
		}
	}

	return int(maxDigit - '0'), maxIndex
}

func main() {
	const filename = "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open the file %s: %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	firstDigit := 0
	secondDigit := 0
	voltage := 0
	totalVoltage := 0

	for scanner.Scan() {
		line := scanner.Text()

		length := len(line)
		fmt.Println("line:", line)
		fmt.Println("len:", length)

		digit1, index1 := findBiggestDigit(line)
		fmt.Printf("digit1:%d, index1:%d\n", digit1, index1)

		if index1 == 0 {
			digit2, index2 := findBiggestDigit(line[1:])
			fmt.Printf("digit2:%d, index2:%d\n", digit2, index2)

			firstDigit = digit1
			secondDigit = digit2
		} else if index1 == length-1 {
			digit2, index2 := findBiggestDigit(line[:length-1])
			fmt.Printf("digit2:%d, index2:%d\n", digit2, index2)

			firstDigit = digit2
			secondDigit = digit1
		} else {
			digit2, index2 := findBiggestDigit(line[:index1])
			fmt.Println("line[:index1]: ", line[:index1])
			fmt.Printf("digit2:%d, index2:%d\n", digit2, index2)

			digit3, index3 := findBiggestDigit(line[index1+1:])
			fmt.Println("line[index1+1:]: ", line[index1+1:])
			fmt.Printf("digit3:%d, index3:%d\n", digit3, index3)

			firstDigit = digit1
			if digit1 == digit2 {
				secondDigit = digit1
			} else {
				secondDigit = digit3
			}
		}
		fmt.Println("----------------------------------------")
		voltage = firstDigit*10 + secondDigit
		totalVoltage += voltage
		fmt.Println("voltage: ", voltage)
		fmt.Println("----------------------------------------")

	}
	fmt.Println("Total voltage: ", totalVoltage)

	if err := scanner.Err(); err != nil {
		fmt.Errorf("file scanning failed: %w", err)
	}
}
