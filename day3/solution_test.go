package day3

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	advent2022.Tester(t, "day3", "../testdata/day3.txt", Solution, "7826")
	advent2022.Tester(t, "day3", "../testdata/day3.txt", SolutionB, "2577")
}
