package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ksusonic/adventofcode24/utils"
)

var mulRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func main() {
	partOne()
}

func partOne() {
	line := strings.Builder{}
	line.Grow(21000) // about number of symbols in input

	err := utils.Load("input.txt", func(s string) error {
		line.WriteString(s)

		return nil
	})
	if err != nil {
		panic(err)
	}

	var result int

	allSubmatch := mulRegexp.FindAllStringSubmatch(line.String(), -1)
	for i := range allSubmatch {
		if len(allSubmatch) < 3 {
			panic(fmt.Errorf("wrong line: %s", allSubmatch[i][0]))
		}

		first, err := strconv.Atoi(allSubmatch[i][1])
		if err != nil {
			panic(fmt.Errorf("wrong first int: %s", allSubmatch[i][0]))
		}

		second, err := strconv.Atoi(allSubmatch[i][2])
		if err != nil {
			panic(fmt.Errorf("wrong second int: %s", allSubmatch[i][0]))
		}

		result += first * second
	}

	fmt.Printf("Part one: %d", result)
}
