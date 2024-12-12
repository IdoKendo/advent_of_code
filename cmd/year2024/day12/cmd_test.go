package day12_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day12"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  140,
			input: "test1.txt",
			fn:    day12.Part1,
		},
		{
			want:  772,
			input: "test1b.txt",
			fn:    day12.Part1,
		},
		{
			want:  1930,
			input: "test1c.txt",
			fn:    day12.Part1,
		},
		{
			want:  80,
			input: "test2.txt",
			fn:    day12.Part2,
		},
		{
			want:  436,
			input: "test2b.txt",
			fn:    day12.Part2,
		},
		{
			want:  236,
			input: "test2c.txt",
			fn:    day12.Part2,
		},
		{
			want:  368,
			input: "test2d.txt",
			fn:    day12.Part2,
		},
		{
			want:  1206,
			input: "test2e.txt",
			fn:    day12.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
