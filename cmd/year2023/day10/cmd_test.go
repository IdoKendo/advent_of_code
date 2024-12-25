package day10_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2023/day10"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  4,
			input: "test1.txt",
			fn:    day10.Part1,
		},
		{
			want:  8,
			input: "test1b.txt",
			fn:    day10.Part1,
		},
		{
			want:  4,
			input: "test2.txt",
			fn:    day10.Part2,
		},
		{
			want:  4,
			input: "test2b.txt",
			fn:    day10.Part2,
		},
		{
			want:  8,
			input: "test2c.txt",
			fn:    day10.Part2,
		},
		{
			want:  10,
			input: "test2d.txt",
			fn:    day10.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
