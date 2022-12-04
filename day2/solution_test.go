package day2

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	advent2022.Tester(t, "day2", "../testdata/day2.txt", Solution, "11475")
	advent2022.Tester(t, "day2", "../testdata/day2.txt", SolutionB, "16862")
}
