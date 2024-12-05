package main

// Input file: day-1/input.txt

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Read input file
func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

// Part 1
func part1(input []string) int {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}

	count := 0
	rows, cols := len(grid), len(grid[0])

	// Directions: right, down, diagonal down-right, diagonal down-left
	// Also search backwards in each direction
	directions := [][]int{
		{0, 1},  // right
		{1, 0},  // down
		{1, 1},  // diagonal down-right
		{1, -1}, // diagonal down-left
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				// Forward search
				if checkXMAS(grid, r, c, dir[0], dir[1]) {
					count++
				}
				// Backward search
				if checkXMAS(grid, r, c, -dir[0], -dir[1]) {
					count++
				}
			}
		}
	}

	return count
}

func checkXMAS(grid [][]rune, startR, startC, deltaR, deltaC int) bool {
	xmas := []rune("XMAS")
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < len(xmas); i++ {
		r := startR + i*deltaR
		c := startC + i*deltaC

		if r < 0 || r >= rows || c < 0 || c >= cols {
			return false
		}

		if grid[r][c] != xmas[i] {
			return false
		}
	}

	return true
}

// Part 2
func part2(input []string) int {
	grid := make([][]rune, len(input))
	for i, line := range input {
		grid[i] = []rune(line)
	}

	count := 0
	rows, cols := len(grid), len(grid[0])

	// For each position that could be the center of an X
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] != 'A' {
				continue
			}

			// Check all possible combinations of MAS in X pattern
			// Top-left to bottom-right and top-right to bottom-left
			tlbr := checkDiagonal(grid, r, c, -1, -1, 1, 1)
			trbl := checkDiagonal(grid, r, c, -1, 1, 1, -1)

			if tlbr && trbl {
				count++
			}
		}
	}

	return count
}

// checkDiagonal checks if there's a MAS sequence (forward or backward) along a diagonal
func checkDiagonal(grid [][]rune, r, c, dr1, dc1, dr2, dc2 int) bool {
	// Check forward MAS
	if (grid[r+dr1][c+dc1] == 'M' && grid[r+dr2][c+dc2] == 'S') ||
		(grid[r+dr1][c+dc1] == 'S' && grid[r+dr2][c+dc2] == 'M') {
		return true
	}
	return false
}

func main() {
	input := readInput()
	fmt.Println("Part 1:", part1(input))
	fmt.Println("Part 2:", part2(input))
}
