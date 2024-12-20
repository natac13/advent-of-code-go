package main

import (
	"strings"
	"testing"
)

var example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestPart1(t *testing.T) {
	want := 2
	got := Part1(parseInput(example))
	if got != want {
		t.Errorf("got Part1() = %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	want := 4
	got := Part2(parseInput(example))
	if got != want {
		t.Errorf("got Part2() = %v, want %v", got, want)
	}
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
