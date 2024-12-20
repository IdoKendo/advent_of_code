package day20_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day20"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		save  int
		input string
		fn    func(string, int) int
	}{
		{
			want:  8,
			save:  12,
			input: "test1.txt",
			fn:    day20.Part1,
		},
		{
			want:  41,
			save:  70,
			input: "test2.txt",
			fn:    day20.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content), tt.save)
		assert.Equal(t, tt.want, got)
	}
}
