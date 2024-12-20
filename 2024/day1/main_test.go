package main

import (
	"strings"
	"testing"
)

var example = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	want := 11
	got := Part1(parseInput(example))
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 31
	got := Part2(parseInput(example))
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
