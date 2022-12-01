package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/yavosh/advent2022"
)

func sum(lines []string) ([]int, error) {
	totals := make([]int, 0)
	acc := 0
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			if acc > 0 {
				totals = append(totals, acc)
				acc = 0
			}

			continue
		}

		i, err := strconv.Atoi(l)
		if err != nil {
			return totals, fmt.Errorf("error parsing line %s %w", l, err)
		}

		acc += i
	}

	return totals, nil
}

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")

	totals, err := sum(lines)
	if err != nil {
		return "", fmt.Errorf("error summing %w", err)
	}
	if len(totals) == 0 {
		return "0", nil
	}

	sort.Ints(totals)
	return fmt.Sprintf("%d", totals[len(totals)-1]), nil
}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")

	totals, err := sum(lines)
	if err != nil {
		return "", fmt.Errorf("error summing %w", err)
	}
	if len(totals) == 0 {
		return "0", nil
	}

	sort.Ints(totals)

	top := totals[len(totals)-3:]
	tot := advent2022.SumInts(top)
	return fmt.Sprintf("%d", tot), nil
}
