# Advent of Code Solutions

> My solutions for the Advent of Code programming puzzles.

## About Advent of Code

[Advent of Code](https://adventofcode.com/) is an annual event featuring programming puzzles released
daily from December 1st to ~~25th~~ 12th. You can read more info in the about section of the site.

This repository contains my solutions. I'm participating casually and always
looking to improve my skills — feel free to suggest optimizations or better approaches!

## Project Structure

- `cmd/year*/day*/`: Solutions for puzzles
- `templates/`: Code templates for generating new day scaffolding
- Each day directory contains:
  - `cmd.go`: Main solution implementation
  - `cmd_test.go`: Unit tests, as given by the puzzle description
  - `test1.txt`, `test2.txt`: Unit test input files
  - `input1.txt`, `input2.txt`: Actual input files

## Setup

Ensure you have Go installed. Clone the repository and navigate to the project directory.

## Usage

This project uses [Cobra](https://github.com/spf13/cobra) for CLI commands.

### Running Tests

```bash
go test ./cmd/year2024/day1/...
```

### Running Solutions

```bash
go run main.go 2024 day1
```

### Bootstrapping New Days

```bash
go run main.go bootstrap 2024 2 && go fmt ./...
```
