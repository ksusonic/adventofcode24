package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rules, updates := readInput(input)

	partOne(rules, updates)
	// partTwo()
}

func partOne(rules map[int]map[int]struct{}, updates [][]int) {
	sum := 0

	for _, update := range updates {
		if correctUpdate(update, rules) {
			sum += update[len(update)/2]
		}
	}

	fmt.Printf("Part one: %d", sum)
}

func correctUpdate(update []int, rules map[int]map[int]struct{}) bool {
	tracked := make(map[int]struct{})

	for _, p := range update {
		rule := rules[p]

		for p := range rule {
			if _, ok := tracked[p]; ok {
				return false
			}
		}

		tracked[p] = struct{}{}
	}

	return true
}

func readInput(input string) (map[int]map[int]struct{}, [][]int) {
	rulesStr, updatesStr, _ := strings.Cut(input, "\n\n")
	rules := make(map[int]map[int]struct{}, len(rulesStr))
	updates := make([][]int, 0, len(input))

	for _, ruleInput := range strings.Split(rulesStr, "\n") {
		aStr, bStr, _ := strings.Cut(ruleInput, "|")
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)

		rule, ok := rules[a]
		if !ok {
			rule = make(map[int]struct{})
		}
		rule[b] = struct{}{}
		rules[a] = rule
	}

	for _, updateInput := range strings.Split(updatesStr, "\n") {
		if len(updateInput) == 0 {
			break
		}

		updateStr := strings.Split(updateInput, ",")
		update := make([]int, 0, len(updateStr))

		for _, pageNumStr := range updateStr {
			pageNum, _ := strconv.Atoi(pageNumStr)
			update = append(update, pageNum)
		}

		updates = append(updates, update)
	}

	return rules, updates
}