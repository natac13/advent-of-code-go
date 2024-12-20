package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// one report per line
// each report is a list fo number called levels separated by spaces

// valid reports are those which:
// levels are all increasing or decreasing
// any 2 adjacent levels differ by at least one and at most 3

func isValidSequence(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	isIncreasing := nums[1] > nums[0]

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 || math.Abs(float64(diff)) > 3 {
			return false
		}
		if i > 1 && ((isIncreasing && diff < 0) || (!isIncreasing && diff > 0)) {
			return false
		}
	}
	return true
}

// Convert string slice to int slice once
func prepareReport(report string) []int {
	raw := strings.Split(report, " ")
	levels := make([]int, len(raw))
	for i, level := range raw {
		levels[i], _ = strconv.Atoi(level)
	}
	return levels
}

func Part1(input []string) int {
	safeReports := 0

	for _, raw := range input {
		report := prepareReport(raw)
		if isValidSequence(report) {
			safeReports++
		}
	}

	return safeReports
}

func Part2(input []string) int {
	safeReports := 0

	for _, report := range input {
		reportLevels := prepareReport(report)

		// Try removing each number and check if sequence is valid
		for skip := 0; skip < len(reportLevels); skip++ {
			// Create temporary slice without current number
			temp := make([]int, 0, len(reportLevels)-1)
			temp = append(temp, reportLevels[:skip]...)
			temp = append(temp, reportLevels[skip+1:]...)

			if isValidSequence(temp) {
				safeReports++
				break
			}
		}
	}

	return safeReports
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
