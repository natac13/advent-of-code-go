package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/natac13/advent-of-code-go/internal/template"
	"github.com/natac13/advent-of-code-go/internal/utils"
)

func main() {
	downloadFlag := flag.Bool("download", false, "Download input file only")
	partFlag := flag.Int("part", 0, "Run specific part (1 or 2). If not specified, runs both")
	createFlag := flag.Bool("create", false, "Create new day")
	flag.Parse()

	if *createFlag {
		if flag.NArg() != 2 {
			fmt.Println("Usage: go run main.go -create YEAR DAY")
			os.Exit(1)
		}
		year, day := flag.Arg(0), flag.Arg(1)
		if err := template.CreateNewDay(year, day); err != nil {
			fmt.Printf("Error creating new day: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go run main.go [-part 1|2] [-download] <year> <day>")
		os.Exit(1)
	}

	year := flag.Args()[0]
	day := flag.Args()[1]

	if *downloadFlag {
		err := utils.DownloadInput(year, day)
		if err != nil {
			fmt.Printf("Error downloading input: %v\n", err)
			os.Exit(1)
		}
		return
	}

	input, err := utils.ReadInput(year, day)
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}

	dayPkg := fmt.Sprintf("./%s/day%s", year, day)

	if *partFlag == 0 || *partFlag == 1 {
		cmd := exec.Command("go", "run", filepath.Join(dayPkg, "main.go"), "1")
		cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running part 1: %v\n", err)
		}
	}

	if *partFlag == 0 || *partFlag == 2 {
		cmd := exec.Command("go", "run", filepath.Join(dayPkg, "main.go"), "2")
		cmd.Stdin = strings.NewReader(strings.Join(input, "\n"))
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error running part 2: %v\n", err)
		}
	}
}
