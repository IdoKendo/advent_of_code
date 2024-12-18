package day14_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day14"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string, int, int) int
	}{
		{
			want:  12,
			input: "test1.txt",
			fn:    day14.Part1,
		},
		{
			want:  0,
			input: "test2.txt",
			fn:    day14.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content), 7, 11)
		assert.Equal(t, tt.want, got)
	}
}
