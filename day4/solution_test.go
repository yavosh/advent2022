package day4

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	advent2022.Tester(t, "day4", "../testdata/day4.txt", Solution, "573")
	advent2022.Tester(t, "day4", "../testdata/day4.txt", SolutionB, "867")
}
