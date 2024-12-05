package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ksusonic/adventofcode24/utils"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	result := 0

	utils.Load("input.txt", func(s string) {
		split := strings.Split(s, " ")

		numericLine := make([]int, 0, len(split))
		for _, ss := range split {
			number, err := strconv.Atoi(ss)
			if err != nil {
				panic(fmt.Errorf("number in line parsing error: %s", s))
			}

			numericLine = append(numericLine, number)
		}

		if safeLine(numericLine) {
			result++
		}
	})

	fmt.Printf("Part one: %d", result)
}

func partTwo() {
	result := 0

	utils.Load("input.txt", func(s string) {
		split := strings.Split(s, " ")

		numericLine := make([]int, 0, len(split))
		for _, ss := range split {
			number, err := strconv.Atoi(ss)
			if err != nil {
				panic(fmt.Errorf("number in line parsing error: %s", s))
			}

			numericLine = append(numericLine, number)
		}

		if safeWithTolerateLine(numericLine) {
			result++
		}
	})

	fmt.Printf("Part two: %d", result)
}

func safeWithTolerateLine(numbers []int) bool {
	if safeLine(numbers) {
		return true
	}

	for i := 0; i < len(numbers); i++ {
		temp := append([]int{}, numbers[:i]...)
		temp = append(temp, numbers[i+1:]...)
		if safeLine(temp) {
			return true
		}
	}

	return false
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
