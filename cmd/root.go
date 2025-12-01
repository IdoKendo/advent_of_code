package cmd

import (
	"fmt"
	"os"

	"github.com/idokendo/aoc/cmd/year2024"
	"github.com/idokendo/aoc/cmd/year2025"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "aoc",
	Long:  "aoc",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(year2024.Cmd)
	rootCmd.AddCommand(year2025.Cmd)
}
