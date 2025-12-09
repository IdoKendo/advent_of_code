package main

import (
	"os"
	"runtime/pprof"

	"github.com/idokendo/aoc/cmd"
)

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	cmd.Execute()
}
