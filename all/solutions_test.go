package all

import (
	"testing"

	"github.com/yavosh/advent2022"
	"github.com/yavosh/advent2022/day1"
	"github.com/yavosh/advent2022/day2"
	"github.com/yavosh/advent2022/day3"
)

func TestAll(t *testing.T) {
	{
		advent2022.Tester(t, "day1-A", "../testdata/day1.txt", day1.Solution, "72718")
		advent2022.Tester(t, "day1-B", "../testdata/day1.txt", day1.SolutionB, "213089")
	}

	{
		advent2022.Tester(t, "day2-A", "../testdata/day2.txt", day2.Solution, "11475")
		advent2022.Tester(t, "day2-B", "../testdata/day2.txt", day2.SolutionB, "16862")
	}

	{
		advent2022.Tester(t, "day3-A", "../testdata/day3.txt", day3.Solution, "7826")
		advent2022.Tester(t, "day3-B", "../testdata/day3.txt", day3.SolutionB, "2577")
	}

}
