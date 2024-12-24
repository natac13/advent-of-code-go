package main

import (
	"bufio"
	"fmt"
	"os"
)

var directions = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

func checkWord(grid []string, word string, row, col, dr, dc int) bool {
	steps := len(word)
	maxRow := len(grid)
	maxCol := len(grid[0])

	endRow := row + (steps-1)*dr
	endCol := col + (steps-1)*dc

	if endRow < 0 || endRow >= maxRow || endCol < 0 || endCol >= maxCol {
		return false
	}

	for i := 0; i < steps; i++ {
		r := row + i*dr
		c := col + i*dc
		if grid[r][c] != word[i] {
			return false
		}
	}

	return true
}

func countOccurrences(grid []string, word string) int {
	result := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			for _, d := range directions {
				if checkWord(grid, word, r, c, d[0], d[1]) {
					result++
				}
			}
		}
	}
	return result
}

func Part1(input []string) int {
	return countOccurrences(input, "XMAS")
}

type Pairing struct {
	pair [2][2]int
}

func checkMAS(grid []string, row, col int, p [2][2]int) bool {
	maxRow := len(grid)
	maxCol := len(grid[0])

	x1 := p[0]
	x2 := p[1]

	if row+x1[0] < 0 || row+x1[0] >= maxRow || col+x1[1] < 0 || col+x1[1] >= maxCol {
		return false
	}

	if row+x2[0] < 0 || row+x2[0] >= maxRow || col+x2[1] < 0 || col+x2[1] >= maxCol {
		return false
	}

	hasX1 := false
	hasX2 := false

	if grid[row+x1[0]][col+x1[1]] == 'M' && grid[row+x2[0]][col+x2[1]] == 'S' {
		hasX1 = true
	} else if grid[row+x1[0]][col+x1[1]] == 'S' && grid[row+x2[0]][col+x2[1]] == 'M' {
		hasX1 = true
	}

	if grid[row+x1[0]][col+x1[1]] == 'S' && grid[row+x2[0]][col+x2[1]] == 'M' {
		hasX2 = true
	} else if grid[row+x1[0]][col+x1[1]] == 'M' && grid[row+x2[0]][col+x2[1]] == 'S' {
		hasX2 = true
	}

	return hasX1 && hasX2

}

func Part2(input []string) int {
	result := 0

	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			if r == 0 || r == len(input)-1 {
				continue
			}

			if input[r][c] == 'A' {
				x1 := checkMAS(input, r, c, [2][2]int{{-1, -1}, {1, 1}})
				x2 := checkMAS(input, r, c, [2][2]int{{-1, 1}, {1, -1}})
				if x1 && x2 {
					result++
				}
			}
		}
	}

	return result
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify which part to run (1 or 2)")
		os.Exit(1)
	}

	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	switch os.Args[1] {
	case "1":
		fmt.Printf("Part 1: %d\n", Part1(input))
	case "2":
		fmt.Printf("Part 2: %d\n", Part2(input))
	default:
		fmt.Printf("Invalid part specified: %s\n", os.Args[1])
		os.Exit(1)
	}
}
