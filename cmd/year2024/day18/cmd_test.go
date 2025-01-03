package day18_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2024/day18"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  string
		input string
		fn    func(string, int, int) string
	}{
		{
			want:  "22",
			input: "test1.txt",
			fn:    day18.Part1,
		},
		{
			want:  "6,1",
			input: "test2.txt",
			fn:    day18.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content), 7, 12)
		assert.Equal(t, tt.want, got)
	}
}
