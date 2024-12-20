package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
)

func Part1(input []string) int {
	left := make([]int, len(input))
	right := make([]int, len(input))

	for i, line := range input {
		fmt.Sscanf(line, "%d   %d", &left[i], &right[i])
	}

	slices.Sort(left)
	slices.Sort(right)

	total := 0

	for i := 0; i < len(left); i++ {
		lv := left[i]
		rv := right[i]
		diff := math.Abs(float64(lv - rv))
		total += int(diff)
	}

	return total
}

func Part2(input []string) int {
	left := make([]int, len(input))
	rightFreq := make(map[int]int)

	// Parse input
	for i, line := range input {
		var r int
		fmt.Sscanf(line, "%d   %d", &left[i], &r)
		rightFreq[r]++
	}

	total := 0
	// For each number in left list, multiply it by its frequency in right list
	for _, num := range left {
		total += num * rightFreq[num]
	}

	return total
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
