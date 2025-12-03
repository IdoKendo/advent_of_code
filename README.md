# Advent Of Code

> My solutions for the puzzles of Advent of Code

# Advent Of Code

[Advent Of Code](https://adventofcode.com/2023) is an event of programming puzzles,
more info can be found in the about section of the site!

I am participating in the event casually. Let me know if I did something
inefficient, or I could've done something better, I want to learn!

# How to run

I created this project using [Cobra](https://github.com/spf13/cobra)
so that it's easy to create new days without having to work hard to prepare it
all in advance and also easy to run, e.g.:

Run the test:
```sh
go test ./cmd/year2024/day1/...
```

Run the actual file:
```sh
go run main.go 2024 day1
```

Add scaffolding for solving a new day:
```sh
go run main.go bootstrap 2024 2 && go fmt ./...
```
