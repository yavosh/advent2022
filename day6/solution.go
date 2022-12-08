package day6

import (
	"fmt"
	"strings"
)

func duplicates(in string) int {
	seen := map[rune]int8{}
	for i, r := range in {
		seen[r]++
		if seen[r] > 1 {
			return i
		}
	}
	return -1
}

func marker(in string, n int) (string, int) {
	buff := make([]rune, 0, n)
	for i, r := range in {
		buff = append(buff, r)

		if len(buff) == n {
			// check
			if duplicates(string(buff)) == -1 {
				return string(buff), i + 1
			}

			// trim first
			buff = buff[1:]
		}
	}

	return string(buff), -1
}

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")
	_, position := marker(lines[0], 4)
	return fmt.Sprintf("%d", position), nil
}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")
	_, position := marker(lines[0], 14)
	return fmt.Sprintf("%d", position), nil
}
