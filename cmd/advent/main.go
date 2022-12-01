package main

import (
	"fmt"

	"github.com/yavosh/advent2022"
	"github.com/yavosh/advent2022/day1"
)

func main() {

	type problem struct {
		solver advent2022.Solver
		in     string
	}

	solutions := map[string]problem{
		"day1":  {solver: day1.Solution, in: "day1.txt"},
		"day1b": {solver: day1.SolutionB, in: "day1.txt"},
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
