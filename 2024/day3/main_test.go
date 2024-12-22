package main

import (
	"strings"
	"testing"
)

var example = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestPart1(t *testing.T) {
	want := 161
	got := Part1(parseInput(example))
	if got != want {
		t.Errorf("Part1() = %v, want %v", got, want)
	}
}

var example2 = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestPart2(t *testing.T) {
	want := 48
	got := Part2(parseInput(example2))
	if got != want {
		t.Errorf("Part2() = %v, want %v", got, want)
	}
}

func parseInput(input string) []string {
	return strings.Split(strings.TrimSpace(input), "\n")
}
