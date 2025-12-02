package day2_test

import (
	"os"
	"testing"

	"github.com/idokendo/aoc/cmd/year2025/day2"
	"github.com/stretchr/testify/assert"
)

func TestParts(t *testing.T) {
	tests := []struct {
		want  int
		input string
		fn    func(string) int
	}{
		{
			want:  1227775554,
			input: "test1.txt",
			fn:    day2.Part1,
		},
		{
			want:  4174379265,
			input: "test2.txt",
			fn:    day2.Part2,
		},
	}

	for _, tt := range tests {
		content, err := os.ReadFile(tt.input)
		assert.NoError(t, err)
		got := tt.fn(string(content))
		assert.Equal(t, tt.want, got)
	}
}
