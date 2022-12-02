package advent2022

import (
	"fmt"
	"testing"
)

func Tester(t *testing.T, prefix, file string, s Solver) {
	in, err := ReadFile(file)
	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}

	if res, err := s(in); err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	} else {
		fmt.Printf("%s is %q\n", prefix, res)
	}
}
