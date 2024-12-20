package template

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
)

//go:embed templates/*
var templates embed.FS

func CreateNewDay(year, day string) error {
	dir := filepath.Join(year, fmt.Sprintf("day%s", day))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	files := []string{"main.go", "main_test.go"}
	for _, f := range files {
		template, err := templates.ReadFile(filepath.Join("templates", f+".tmpl"))
		if err != nil {
			return err
		}

		outPath := filepath.Join(dir, f)
		if _, err := os.Stat(outPath); err == nil {
			// File exists, skip
			continue
		}

		if err := os.WriteFile(outPath, template, 0644); err != nil {
			return err
		}
	}

	return nil
}
