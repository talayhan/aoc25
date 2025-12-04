package main

import (
	"bufio"
	"fmt"
	"os"
)

const gridSize = 138

type Matrix [][]rune

func readMatrixFromFile(filepath string) (Matrix, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	matrix := make(Matrix, 0, gridSize)
	scanner := bufio.NewScanner(file)
	rowIndex := 0

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != gridSize {
			return nil, fmt.Errorf("line %d has length %d, expected %d", rowIndex, len(line), gridSize)
		}

		matrix = append(matrix, []rune(line))
		rowIndex++
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	if rowIndex != gridSize {
		return nil, fmt.Errorf("file has %d rows, expected %d", rowIndex, gridSize)
	}

	return matrix, nil
}

func printMatrix(matrix Matrix) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func countAtNeighbors(matrix Matrix, r, c int) int {
	rows := len(matrix)
	if rows == 0 {
		return 0
	}
	cols := len(matrix[0])
	count := 0

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}

			nr, nc := r+dr, c+dc
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				if matrix[nr][nc] == '@' {
					count++
				}
			}
		}
	}
	return count
}

func transformMatrix(original Matrix) (Matrix, int) {
	xCount := 0
	rows := len(original)
	if rows == 0 {
		return nil, 0
	}
	cols := len(original[0])

	newMatrix := make(Matrix, rows)
	for r := range newMatrix {
		newMatrix[r] = make([]rune, cols)
		copy(newMatrix[r], original[r])
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if original[r][c] == '@' {
				atCount := countAtNeighbors(original, r, c)

				if atCount < 4 {
					newMatrix[r][c] = 'X'
					xCount++
				}
			}
		}
	}

	return newMatrix, xCount
}

func main() {
	filepath := "input.txt"

	originalMatrix, err := readMatrixFromFile(filepath)
	if err != nil {
		fmt.Println("failed to read file:", err)
		return
	}

	printMatrix(originalMatrix)
	transformedMatrix, xCount := transformMatrix(originalMatrix)
	fmt.Println("xCount: ", xCount)
	printMatrix(transformedMatrix)
}
