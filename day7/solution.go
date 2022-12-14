package day7

import (
	"fmt"
	"strings"

	"github.com/yavosh/advent2022/terminal"
)

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")
	term := terminal.New()
	for _, line := range lines {
		if err := term.Read(line); err != nil {
			panic(err)
		}
	}

	term.Print()
	answer, err := term.SumSizeUpTo(100000)
	return fmt.Sprintf("%d", answer), err
}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")
	for _ = range lines {

	}

	answer := ""
	return fmt.Sprintf("%s", answer), nil
}
