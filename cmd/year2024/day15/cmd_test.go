package day15_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day15"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  10092,
			input: "test1.txt",
			fn:    day15.Part1,
		},
		{
			want:  2028,
			input: "test1b.txt",
			fn:    day15.Part1,
		},
		{
			want:  9021,
			input: "test2.txt",
			fn:    day15.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
