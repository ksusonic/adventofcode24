package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ksusonic/adventofcode24/utils"
)

func main() {
	builder := strings.Builder{}
	builder.Grow(21000) // about number of symbols in input

	err := utils.Load("input.txt", func(s string) error {
		builder.WriteString(s)

		return nil
	})
	if err != nil {
		panic(err)
	}

	line := builder.String()

	partOne(line)
	partTwo(line)
}

func partOne(line string) {
	var (
		result    int
		mulRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	)

	allSubmatch := mulRegexp.FindAllStringSubmatch(line, -1)
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

func partTwo(line string) {
	var (
		result    int
		mulRegexp = regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`)
	)

	allSubmatch := mulRegexp.FindAllStringSubmatch(line, -1)
	do := true

	for i := range allSubmatch {
		line := allSubmatch[i][0]

		switch {
		case strings.HasPrefix(line, "mul"):
			if !do {
				continue
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
		case strings.HasPrefix(line, "don't"):
			do = false
		case strings.HasPrefix(line, "do"):
			do = true
		default:
			panic(fmt.Errorf("unknown match: %s", line))
		}
	}

	fmt.Printf("Part two: %d", result)
}
