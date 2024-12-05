package utils

import (
	"bufio"
	"os"
)

func Load(filename string, lineProces func(string)) {
	input, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		lineProces(scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}
}
