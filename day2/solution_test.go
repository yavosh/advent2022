package day2

import (
	"fmt"
	"testing"

	"github.com/yavosh/advent2022"
)

func TestSolution(t *testing.T) {
	in, err := advent2022.ReadFile("../testdata/day2.txt")
	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}

	if res, err := Solution(in); err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	} else {
		fmt.Printf("res: %s\n", res)
	}
}

func TestSolutionB(t *testing.T) {
	in, err := advent2022.ReadFile("../testdata/day2.txt")
	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}

	if res, err := SolutionB(in); err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	} else {
		fmt.Printf("res: %s\n", res)
	}
}
