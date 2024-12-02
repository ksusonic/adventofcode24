package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ksusonic/adventofcode24/utils"
)

func main() {
	part_one()
}

func part_one() {
	result := 0

	err := utils.Load("input.txt", func(s string) error {
		split := strings.Split(s, " ")

		numericLine := make([]int, 0, len(split))
		for _, ss := range split {
			number, err := strconv.Atoi(ss)
			if err != nil {
				return fmt.Errorf("number in line parsing error: %s", s)
			}

			numericLine = append(numericLine, int(number))
		}

		if safeLine(numericLine) {
			result++
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part one: %d", result)
}

func safeLine(numbers []int) bool {
	if len(numbers) < 2 {
		panic("unexpected line size")
	}

	asc := numbers[1] > numbers[0]

	for i := 1; i < len(numbers); i++ {
		diff := numbers[i] - numbers[i-1]
		absDiff := int(math.Abs(float64(diff)))

		if absDiff < 1 || absDiff > 3 || asc && diff < 0 || !asc && diff > 0 {
			return false
		}
	}

	return true
}
