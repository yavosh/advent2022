package all

import (
	"testing"

	"github.com/yavosh/advent2022"
	"github.com/yavosh/advent2022/day1"
	"github.com/yavosh/advent2022/day2"
)

func TestAll(t *testing.T) {
	{
		advent2022.Tester(t, "day1-A", "../testdata/day1.txt", day1.Solution)
		advent2022.Tester(t, "day1-B", "../testdata/day1.txt", day1.SolutionB)
	}

	{
		advent2022.Tester(t, "day2-A", "../testdata/day2.txt", day2.Solution)
		advent2022.Tester(t, "day2-B", "../testdata/day2.txt", day2.SolutionB)
	}
}
