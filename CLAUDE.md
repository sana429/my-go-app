# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
# Run the application
go run main.go

# Build
go build -o my-go-app .

# Run tests
go test ./...

# Run a single test
go test ./internal/ -run TestName
```

## Architecture

This is a CLI menu application. `main.go` presents an interactive menu and dispatches to handlers based on user input. Business logic lives in `internal/` and is imported via the `my-go-app/internal` package path.

- **`main.go`**: Menu loop and top-level routing
- **`internal/greet.go`**: `GreetUser()` — prompts for a name and prints a greeting
- **`internal/datetime.go`**: `ShowDateTime()` — prints the current date and time using Go's `time` package

New menu options should follow the same pattern: add a case in `main.go` and implement the logic as a function in `internal/`.
