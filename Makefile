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
	@go run -mod=mod main.go -create $(year) $(day)
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
