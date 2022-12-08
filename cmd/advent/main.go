package main

import (
	"fmt"

	"github.com/yavosh/advent2022"
	"github.com/yavosh/advent2022/day1"
	"github.com/yavosh/advent2022/day2"
	"github.com/yavosh/advent2022/day3"
	"github.com/yavosh/advent2022/day4"
	"github.com/yavosh/advent2022/day5"
	"github.com/yavosh/advent2022/day6"
)

func main() {

	type problem struct {
		solver advent2022.Solver
		in     string
	}

	solutions := map[string]problem{
		"day1a": {solver: day1.Solution, in: "day1.txt"},
		"day1b": {solver: day1.SolutionB, in: "day1.txt"},

		"day2a": {solver: day2.Solution, in: "day2.txt"},
		"day2b": {solver: day2.SolutionB, in: "day2.txt"},

		"day3a": {solver: day3.Solution, in: "day3.txt"},
		"day3b": {solver: day3.SolutionB, in: "day3.txt"},

		"day4a": {solver: day4.Solution, in: "day4.txt"},
		"day4b": {solver: day4.SolutionB, in: "day4.txt"},

		"day5a": {solver: day5.Solution, in: "day5.txt"},
		"day5b": {solver: day5.SolutionB, in: "day5.txt"},

		"day6a": {solver: day6.Solution, in: "day6.txt"},
		"day6b": {solver: day6.SolutionB, in: "day6.txt"},
	}

	for name, s := range solutions {
		f, err := advent2022.ReadFile("./testdata/" + s.in)
		if err != nil {
			fmt.Printf("error solving %s %v\n", name, err)
			continue
		}

		solution, err := s.solver(f)
		if err != nil {
			fmt.Printf("error solving %s %v\n", name, err)
			continue
		}

		fmt.Printf("solution  %s %v\n", name, solution)
	}
}
