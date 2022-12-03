package day3

import (
	"fmt"
	"strings"
)

type dedupT map[rune]int

// val will return the ordinal of the char
func val(c rune) int {
	b := byte(c)

	if b >= 'a' && b <= 'z' {
		return int(b) - ('a') + 1
	}

	if b >= 'A' && b <= 'Z' {
		return int(b) - ('A') + 1 + 26
	}

	panic("bad item " + string(b))
}

func duplicates(in string) dedupT {
	dup := dedupT{}
	left := in[0 : len(in)/2]
	right := in[len(in)/2:]

	for _, l := range left {
		for _, r := range right {
			if l == r {
				dup[l] = dup[l] + 1
			}
		}
	}

	return dup
}

func duplicatesB(in []string) dedupT {
	dup := dedupT{}

	for _, d1 := range in[0] {
		for _, d2 := range in[1] {
			for _, d3 := range in[2] {
				if d1 == d2 && d2 == d3 {
					dup[d1]++
				}
			}
		}
	}

	return dup
}

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")

	sum := 0

	for _, line := range lines {
		if len(line)%2 != 0 {
			return "", fmt.Errorf("error line checksum (len not even)")
		}

		dup := duplicates(strings.TrimSpace(line))

		acc := 0
		for key := range dup {
			acc += val(key)
		}

		sum += acc
	}

	return fmt.Sprintf("%d", sum), nil
}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")
	sum := 0
	group := make([]string, 3)
	results := make([]dedupT, 0)

	for i, line := range lines {
		if i%3 == 0 && i != 0 {
			// group
			dup := duplicatesB(group)
			results = append(results, dup)
		}

		group[i%3] = line
	}

	// last group
	dup := duplicatesB(group)
	results = append(results, dup)

	fmt.Printf("results: %d %v\n", len(results), results)
	for i, r := range results {
		if len(r) != 1 {
			panic("multiple results in group " + fmt.Sprintf("%d", i))
		}

		for k := range r {
			// should be only one
			sum += val(k)
		}
	}

	return fmt.Sprintf("%d", sum), nil
}
