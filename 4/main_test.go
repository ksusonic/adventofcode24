package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMatrix = []string{
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

func Test_FindXMAS(t *testing.T) {
	assert.Equal(t, 18, findXMAS(strings.Join(testMatrix, ""), len(testMatrix[0])))
}

func Test_FindMAS(t *testing.T) {
	assert.Equal(t, 9, findMAS(strings.Join(testMatrix, ""), len(testMatrix[0])))
}
