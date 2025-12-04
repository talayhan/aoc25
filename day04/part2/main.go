package main

import (
	"bufio"
	"fmt"
	"os"
)

const gridSize = 138

type Matrix [][]rune

func countNeighbors(m Matrix, r, c int) int {
	count := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}

			nr, nc := r+dr, c+dc

			if nr >= 0 && nr < gridSize && nc >= 0 && nc < gridSize {
				if m[nr][nc] == '@' {
					count++
				}
			}
		}
	}
	return count
}

func runSinglePass(current Matrix) (Matrix, int) {
	next := make(Matrix, gridSize)
	changes := 0

	for r := 0; r < gridSize; r++ {
		next[r] = make([]rune, gridSize)
		copy(next[r], current[r])
	}

	for r := 0; r < gridSize; r++ {
		for c := 0; c < gridSize; c++ {
			if current[r][c] == '@' {
				neighbors := countNeighbors(current, r, c)

				if neighbors < 4 {
					changes++
				}
			}
		}
	}

	return next, changes
}

func stabilizeMatrix(initial Matrix) {
	current := initial
	totalRemoved := 0
	pass := 1

	for {
		next, removedCount := runSinglePass(current)

		if removedCount == 0 {
			break
		}

		fmt.Printf("Pass %d: Removed %d chars\n", pass, removedCount)

		current = next
		totalRemoved += removedCount
		pass++
	}

	fmt.Printf("Stabilization complete. Total '@' chars removed: %d\n", totalRemoved)
	fmt.Println("Final Matrix State:")
	printMatrix(current)
}

func readMatrix(filename string) (Matrix, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix Matrix
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == gridSize {
			matrix = append(matrix, []rune(line))
		}
	}

	if len(matrix) != gridSize {
		return nil, fmt.Errorf("invalid matrix size: got %d rows, expected %d", len(matrix), gridSize)
	}

	return matrix, nil
}

func printMatrix(m Matrix) {
	for _, row := range m {
		fmt.Println(string(row))
	}
}

func main() {
	filename := "input.txt"

	matrix, err := readMatrix(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Initial State:")
	printMatrix(matrix)
	fmt.Println("Starting stabilization...")
	stabilizeMatrix(matrix)
}
