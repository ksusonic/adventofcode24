package utils

import (
	"bufio"
	"os"
)

func Load(filename string, lineProces func(string) error) error {
	input, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		if err = lineProces(scanner.Text()); err != nil {
			return err
		}
	}

	return scanner.Err()
}
