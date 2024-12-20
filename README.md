# Advent of Code - Go Solutions

## Setup

1. Create a `.env` file with your Advent of Code session token:

```
AOC_SESSION=your_session_token_here
```

To get your session token, login to Advent of Code, inspect your cookies, and copy the value of the 'session' cookie.

2. Create a new day's solution:

```bash
make new year=2023 day=1
```

3. Run a specific solution:

```bash
make run year=2023 day=1
```

## Project Structure

```
.
├── internal/
│   └── utils/        # Common utilities
├── yearXXXX/         # Year folders
│   └── dayXX/        # Day folders
│       ├── main.go   # Solution code
│       └── input.txt # Day's input
└── main.go           # Entry point
```
