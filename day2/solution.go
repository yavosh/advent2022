package day2

import (
	"fmt"
	"strings"
)

type outcome string

const (
	win  outcome = "W"
	lose outcome = "L"
	draw outcome = "D"
)

func play(left, right string) outcome {
	outcomes := map[string]outcome{
		"AX": draw,
		"AY": win,
		"AZ": lose,

		"BX": lose,
		"BY": draw,
		"BZ": win,

		"CX": win,
		"CY": lose,
		"CZ": draw,
	}

	if val, ok := outcomes[left+right]; ok {
		return val
	}
	panic("situation " + left + right)
}

func playB(left, right string) (outcome, string) {

	switch right {
	case "X":
		// lose
		switch left {
		case "A":
			return lose, "Z"
		case "B":
			return lose, "X"
		case "C":
			return lose, "Y"
		}
	case "Y":
		//draw
		switch left {
		case "A":
			return draw, "X"
		case "B":
			return draw, "Y"
		case "C":
			return draw, "Z"
		}
	case "Z":
		//win
		switch left {
		case "A":
			return win, "Y"
		case "B":
			return win, "Z"
		case "C":
			return win, "X"
		}
	}

	panic("situation " + left + right)
}

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")

	score := 0
	for _, line := range lines {
		sides := strings.Split(strings.TrimSpace(line), " ")
		outcome := play(sides[0], sides[1])

		switch sides[1] {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		switch outcome {
		case draw:
			score += 3
		case win:
			score += 6
		}

	}

	return fmt.Sprintf("%d", score), nil

}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")

	score := 0
	for _, line := range lines {
		sides := strings.Split(strings.TrimSpace(line), " ")
		outcome, res := playB(sides[0], sides[1])

		switch res {
		case "X":
			score += 1
		case "Y":
			score += 2
		case "Z":
			score += 3
		}

		switch outcome {
		case draw:
			score += 3
		case win:
			score += 6
		}

	}

	return fmt.Sprintf("%d", score), nil

}
