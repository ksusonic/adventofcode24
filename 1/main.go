package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/ksusonic/adventofcode24/utils"
)

func main() {
	part_one()
	part_two()
}

func part_one() {
	var (
		left  = make([]int, 0, 1000)
		right = make([]int, 0, 1000)
	)

	if err := utils.Load("input.txt", func(s string) error {
		split := strings.Split(s, "   ")
		if len(split) != 2 {
			return fmt.Errorf("Unknown string for split: %s", s)
		}

		leftNum, err := strconv.ParseInt(split[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Error parsing left number: %s", s)
		}

		rightNum, err := strconv.ParseInt(split[1], 10, 32)
		if err != nil {
			return fmt.Errorf("Error parsing right number: %s", s)
		}

		left = append(left, int(leftNum))
		right = append(right, int(rightNum))

		return nil
	}); err != nil {
		panic(err)
	}

	slices.Sort(left)
	slices.Sort(right)

	if len(left) != len(right) {
		panic(fmt.Errorf("different left and right len: %d vs %d", len(left), len(right)))
	}

	diffSum := 0

	for i := range len(left) {
		if left[i] > right[i] {
			diffSum += left[i] - right[i]
		} else {
			diffSum += right[i] - left[i]
		}
	}

	fmt.Printf("Part 2 result: %d", diffSum)
}

func part_two() {
	var (
		leftUnique = make(map[int]struct{}, 1000)
		right      = make(map[int]int, 1000)
	)

	if err := utils.Load("input.txt", func(s string) error {
		split := strings.Split(s, "   ")
		if len(split) != 2 {
			return fmt.Errorf("Unknown string for split: %s", s)
		}

		leftNum, err := strconv.ParseInt(split[0], 10, 32)
		if err != nil {
			return fmt.Errorf("Error parsing left number: %s", s)
		}

		rightNum, err := strconv.ParseInt(split[1], 10, 32)
		if err != nil {
			return fmt.Errorf("Error parsing right number: %s", s)
		}

		leftUnique[int(leftNum)] = struct{}{}
		right[int(rightNum)]++

		return nil
	}); err != nil {
		panic(err)
	}

	simScore := 0

	for left := range leftUnique {
		simScore += left * right[left]
	}

	fmt.Printf("Part 2 result: %d", simScore)
}
