package all

import (
	"testing"

	"github.com/yavosh/advent2022"
	"github.com/yavosh/advent2022/day1"
	"github.com/yavosh/advent2022/day2"
	"github.com/yavosh/advent2022/day3"
	"github.com/yavosh/advent2022/day4"
)

func TestAll(t *testing.T) {

	tests := []struct {
		name     string
		file     string
		solution advent2022.Solver
		expected string
	}{
		{"day1-A", "../testdata/day1.txt", day1.Solution, "72718"},
		{"day1-B", "../testdata/day1.txt", day1.SolutionB, "213089"},
		{"day2-A", "../testdata/day2.txt", day2.Solution, "11475"},
		{"day2-B", "../testdata/day2.txt", day2.SolutionB, "16862"},
		{"day3-A", "../testdata/day3.txt", day3.Solution, "7826"},
		{"day3-B", "../testdata/day3.txt", day3.SolutionB, "2577"},
		{"day4-A", "../testdata/day4.txt", day4.Solution, "573"},
		{"day4-B", "../testdata/day4.txt", day4.SolutionB, "867"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			advent2022.Tester(t, tt.name, tt.file, tt.solution, tt.expected)
		})
	}

}
