package day5

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	t.Run("day5", func(t *testing.T) {
		advent2022.Tester(t, "day5", "../testdata/day5.txt", Solution, "FWSHSPJWM")

	})

	t.Run("day5b", func(t *testing.T) {
		advent2022.Tester(t, "day5", "../testdata/day5.txt", SolutionB, "PWPWHGFZS")
	})
}
