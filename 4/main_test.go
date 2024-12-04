package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindXMAS_Vertical(t *testing.T) {
	matrix := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}

	assert.Equal(t, 18, findXMAS(strings.Join(matrix, ""), len(matrix[0])))
}
