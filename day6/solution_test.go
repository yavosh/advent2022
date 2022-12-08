package day6

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	t.Run("day6", func(t *testing.T) {
		advent2022.Tester(t, "day6", "../testdata/day6.txt", Solution, "1198")

	})

	t.Run("day6b", func(t *testing.T) {
		advent2022.Tester(t, "day6", "../testdata/day6.txt", SolutionB, "3120")
	})
}
