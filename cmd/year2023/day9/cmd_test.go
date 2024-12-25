package day9_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2023/day9"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  114,
			input: "test1.txt",
			fn:    day9.Part1,
		},
		{
			want:  2,
			input: "test2.txt",
			fn:    day9.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
