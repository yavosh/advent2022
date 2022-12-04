package day1

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	advent2022.Tester(t, "day1", "../testdata/day1.txt", Solution, "72718")
	advent2022.Tester(t, "day1", "../testdata/day1.txt", SolutionB, "213089")
}
