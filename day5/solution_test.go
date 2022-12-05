package day4

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	advent2022.Tester(t, "day5", "../testdata/day5.txt", Solution, "FWSHSPJWM")
	advent2022.Tester(t, "day5", "../testdata/day5.txt", SolutionB, "MCD")
}
