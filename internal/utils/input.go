package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func DownloadInput(year, day string) error {
	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return fmt.Errorf("AOC_SESSION environment variable not set")
	}

	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad response: %s", resp.Status)
	}

	dir := filepath.Join(year, fmt.Sprintf("day%s", day))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(dir, "input.txt"))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func ReadInput(year, day string) ([]string, error) {
	file, err := os.Open(filepath.Join(year, fmt.Sprintf("day%s", day), "input.txt"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
