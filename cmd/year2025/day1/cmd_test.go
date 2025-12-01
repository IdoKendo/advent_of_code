package day1_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2025/day1"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  3,
			input: "test1.txt",
			fn:    day1.Part1,
		},
		{
			want:  6,
			input: "test2.txt",
			fn:    day1.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
