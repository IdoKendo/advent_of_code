package day16_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day16"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  7036,
			input: "test1.txt",
			fn:    day16.Part1,
		},
		{
			want:  11048,
			input: "test1b.txt",
			fn:    day16.Part1,
		},
		{
			want:  45,
			input: "test2.txt",
			fn:    day16.Part2,
		},
		{
			want:  64,
			input: "test2b.txt",
			fn:    day16.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
