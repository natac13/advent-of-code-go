package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1(input []string) int {
	sum := 0

	re, err := regexp.Compile(`mul\((\d{1,3}),(\d{1,3})\)`)
	if err != nil {
		panic(err)
	}

	for _, line := range input {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, err := strconv.Atoi(match[1])
			if err != nil {
				continue
			}
			b, err := strconv.Atoi(match[2])
			if err != nil {
				continue
			}
			sum += a * b
		}

	}

	return sum
}

func Part2(input []string) int {
	sum := 0
	enabled := true

	mulRe := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	doNotRe := regexp.MustCompile(`don't\(\)`)

	for _, line := range input {
		mulMatches := mulRe.FindAllStringSubmatchIndex(line, -1)
		doMatches := doRe.FindAllStringSubmatchIndex(line, -1)
		doNotMatches := doNotRe.FindAllStringSubmatchIndex(line, -1)

		pos := 0
		for pos < len(line) {

			if len(doMatches) > 0 && doMatches[0][0] == pos {
				enabled = true
				pos = doMatches[0][1]     // move pos to end of match
				doMatches = doMatches[1:] // remove the match from the list
				continue
			}

			if len(doNotMatches) > 0 && doNotMatches[0][0] == pos {
				enabled = false
				pos = doNotMatches[0][1]        // move pos to end of match
				doNotMatches = doNotMatches[1:] // remove the match from the list
				continue
			}

			if len(mulMatches) > 0 && mulMatches[0][0] == pos {
				if enabled {
					num1, err := strconv.Atoi(line[mulMatches[0][2]:mulMatches[0][3]])
					if err != nil {
						panic(err)
					}
					num2, err := strconv.Atoi(line[mulMatches[0][4]:mulMatches[0][5]])
					if err != nil {
						panic(err)
					}
					sum += num1 * num2
				}
				pos = mulMatches[0][1]      // move pos to end of match
				mulMatches = mulMatches[1:] // remove the match from the list
				continue
			}
			pos++
		}
	}

	return sum
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
