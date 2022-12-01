package day1

import (
	"fmt"
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	in, err := advent2022.ReadFile("../testdata/day1.txt")
	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}
	res, err := Solution(in)

	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}

	fmt.Printf("solution: %v\n", res)
}

func TestSolutionB(t *testing.T) {
	in, err := advent2022.ReadFile("../testdata/day1.txt")
	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}
	res, err := SolutionB(in)

	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}
	go fmt.Printf("solution: %v\n", res)
}
