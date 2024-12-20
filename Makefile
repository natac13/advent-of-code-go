.PHONY: run new download-input test

run:
	@if [ -z "$(year)" ] || [ -z "$(day)" ]; then \
		echo "Usage: make run year=YYYY day=D [part=1|2]"; \
		exit 1; \
	fi
	@if [ -z "$(part)" ]; then \
		go run main.go $(year) $(day); \
	else \
		go run main.go -part $(part) $(year) $(day); \
	fi

new:
	@if [ -z "$(year)" ] || [ -z "$(day)" ]; then \
		echo "Usage: make new year=YYYY day=D"; \
		exit 1; \
	fi
	@mkdir -p $(year)/day$(day)
	@if [ ! -f $(year)/day$(day)/main.go ]; then \
		echo 'package main\n\nfunc Part1(input []string) int {\n\treturn 0\n}\n\nfunc Part2(input []string) int {\n\treturn 0\n}' > $(year)/day$(day)/main.go; \
	fi
	@if [ ! -f $(year)/day$(day)/main_test.go ]; then \
		echo 'package main\n\nimport (\n\t"strings"\n\t"testing"\n)\n\nvar example = `input here`\n\nfunc TestPart1(t *testing.T) {\n\twant := 0\n\tgot := Part1(parseInput(example))\n\tif got != want {\n\t\tt.Errorf("Part1() = %v, want %v", got, want)\n\t}\n}\n\nfunc TestPart2(t *testing.T) {\n\twant := 0\n\tgot := Part2(parseInput(example))\n\tif got != want {\n\t\tt.Errorf("Part2() = %v, want %v", got, want)\n\t}\n}\n\nfunc parseInput(input string) []string {\n\treturn strings.Split(strings.TrimSpace(input), "\n")\n}' > $(year)/day$(day)/main_test.go; \
	fi
	@make download-input year=$(year) day=$(day)

download-input:
	@if [ -z "$(year)" ] || [ -z "$(day)" ]; then \
		echo "Usage: make download-input year=YYYY day=D"; \
		exit 1; \
	fi
	@go run main.go -download $(year) $(day)

test:
	@if [ -z "$(year)" ] || [ -z "$(day)" ]; then \
		go test ./...; \
	else \
		go test ./$(year)/day$(day); \
	fi
