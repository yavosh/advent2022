package advent2022

import (
	"testing"
)

func Tester(t *testing.T, prefix, file string, s Solver, expected string) {
	in, err := ReadFile(file)
	if err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	}

	if res, err := s(in); err != nil {
		t.Errorf("error %v", err)
		t.FailNow()
	} else {
		if res != expected {
			t.Errorf("want %q got %q", expected, res)
		}
	}

}
