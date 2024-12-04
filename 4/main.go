package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/ksusonic/adventofcode24/utils"
)

func main() {
	builder := strings.Builder{}
	builder.Grow(20000)

	rowLen := 0

	err := utils.Load("input.txt", func(s string) error {
		builder.WriteString(s)
		if rowLen == 0 {
			rowLen = len(s)
		}

		if len(s) != rowLen {
			panic("row len differs")
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	partOne(builder.String(), rowLen)
	partTwo(builder.String(), rowLen)
}

func partOne(s string, size int) {
	res := findXMAS(s, size)

	fmt.Printf("Part one: %d", res)
}

func partTwo(s string, size int) {
	res := findMAS(s, size)

	fmt.Printf("Part two: %d", res)
}

func findXMAS(s string, size int) int {
	count := 0

	for i := range s {
		horizontalPos := i % size

		if s[i] == 'X' {
			if i+3 < len(s) && horizontalPos+3 < size && s[i:i+4] == "XMAS" {
				count++
			}

			if 0 <= i-3 && 0 <= horizontalPos-3 && s[i-3:i+1] == "SAMX" {
				count++
			}

			// up vertical
			if i >= size*3 {
				if s[i-size] == 'M' && s[i-size*2] == 'A' && s[i-size*3] == 'S' {
					count++
				}
			}

			// down vertical
			if i+size*3 < len(s) {
				if s[i+size] == 'M' && s[i+size*2] == 'A' && s[i+size*3] == 'S' {
					count++
				}
			}

			// left upper diagonal
			if 0 <= horizontalPos-3 && i >= size*3 {
				if s[i-size-1] == 'M' && s[i-size*2-2] == 'A' && s[i-size*3-3] == 'S' {
					count++
				}
			}

			// right upper diagonal
			if horizontalPos+3 < size && i >= size*3 {
				if s[i-size+1] == 'M' && s[i-size*2+2] == 'A' && s[i-size*3+3] == 'S' {
					count++
				}
			}

			// left lower diagonal
			if 0 <= horizontalPos-3 && i+size*3 < len(s) {
				if s[i+size-1] == 'M' && s[i+size*2-2] == 'A' && s[i+size*3-3] == 'S' {
					count++
				}
			}

			// right lower diagonal
			if horizontalPos+3 < size && i+size*3 < len(s) {
				if s[i+size+1] == 'M' && s[i+size*2+2] == 'A' && s[i+size*3+3] == 'S' {
					count++
				}
			}
		}
	}

	return count
}

func findMAS(s string, size int) int {
	count := 0

	allowed := map[string]struct{}{
		"MMSS": {},
		"MSSM": {},
		"SSMM": {},
		"SMMS": {},
	}

	for i := size + 1; i < len(s)-size-1; i++ {
		horizontalPos := i % size

		if s[i] == 'A' && 0 < horizontalPos && horizontalPos < size-1 {
			buff := bytes.NewBufferString("")
			buff.WriteByte(s[i-size-1])
			buff.WriteByte(s[i-size+1])
			buff.WriteByte(s[i+size+1])
			buff.WriteByte(s[i+size-1])

			if _, ok := allowed[buff.String()]; ok {
				count++
			}
		}
	}

	return count
}
