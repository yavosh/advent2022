package day7

import (
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	t.Run("day7", func(t *testing.T) {
		// 1055975
		// 1055975
		advent2022.Tester(t, "day7", "../testdata/day7.txt", Solution, "x")
	})

	t.Run("day7b", func(t *testing.T) {
		advent2022.Tester(t, "day7", "../testdata/day7t.txt", SolutionB, "x")
	})
}
