package main

import (
	"strings"
	"testing"
)

var example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestPart1(t *testing.T) {
	want := 18
	got := Part1(parseInput(example))
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 9
	got := Part2(parseInput(example))
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
