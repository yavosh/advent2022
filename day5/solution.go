package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type command struct {
	count, from, to int
}

type stack struct {
	data []string
}

type state struct {
	stacks []stack
	ids    []string
}

func (s *stack) size() int {
	return len(s.data)
}

func (s *stack) push(val string) {
	n := make([]string, 0, len(s.data)+1)
	n = append(n, val)
	n = append(n, s.data...)

	s.data = n
}

func (s *stack) pop() string {
	head := s.data[0]
	s.data = s.data[1:]
	return head
}

func (s *state) move9000(cmd command) {
	for i := 0; i < cmd.count; i++ {
		val := s.stacks[cmd.from-1].pop()
		s.stacks[cmd.to-1].push(val)
	}
}

func (s *state) move9001(cmd command) {
	buff := make([]string, 0)
	for i := 0; i < cmd.count; i++ {
		buff = append(buff, s.stacks[cmd.from-1].pop())
	}

	for len(buff) > 0 {
		tail := buff[len(buff)-1]
		buff = buff[0 : len(buff)-1]
		s.stacks[cmd.to-1].push(tail)
	}
}

func (s *state) height() int {
	res := 0

	for i := range s.stacks {
		if s.stacks[i].size() > res {
			res = s.stacks[i].size()
		}
	}
	return res
}

func (s *state) print() {
	for i, s := range s.stacks {
		fmt.Printf("%d | ", i+1)
		for _, val := range s.data {
			fmt.Printf("%s ", val)
		}

		fmt.Println()
	}

	fmt.Println()
}

func load(lines []string) *state {
	tail := lines[len(lines)-1]
	data := lines[:len(lines)-1]
	ids := tokens(tail)

	stacks := make([]stack, len(ids))

	for i := len(data) - 1; i >= 0; i-- {
		line := data[i]
		pos := 0
		for i := range ids {
			val := line[pos : pos+3]
			if strings.TrimSpace(val) != "" {
				stacks[i].push(val)
			}
			pos += 3
			pos += 1
		}
	}

	return &state{ids: ids, stacks: stacks}
}

func tokens(line string) []string {
	res := make([]string, 0, len(line)/3)
	for _, val := range strings.Split(line, " ") {
		if strings.TrimSpace(val) == "" {
			continue
		}

		res = append(res, strings.TrimSpace(val))
	}
	return res
}

func parseCommand(line string) (command, error) {
	cmd := tokens(line)
	res := command{}

	if val, err := strconv.Atoi(cmd[1]); err != nil {
		return command{}, fmt.Errorf("error parsing commands %s", line)
	} else {
		res.count = val
	}

	if val, err := strconv.Atoi(cmd[3]); err != nil {
		return command{}, fmt.Errorf("error parsing commands %s", line)
	} else {
		res.from = val
	}

	if val, err := strconv.Atoi(cmd[5]); err != nil {
		return command{}, fmt.Errorf("error parsing commands %s", line)
	} else {
		res.to = val
	}

	return res, nil
}

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")

	curr := make([]string, 0)
	start := 0
	for i, line := range lines {
		if line == "" {
			start = i
			break
		}
		curr = append(curr, line)
	}

	crates := load(curr)

	for i, line := range lines {
		if i <= start {
			// skip state block
			continue
		}

		if strings.TrimSpace(line) == "" {
			// ignore empty lines
			continue
		}

		cmd, err := parseCommand(line)
		if err != nil {
			return "", fmt.Errorf("error parsing command %d %s", i, line)
		}
		crates.move9000(cmd)
	}

	result := ""
	for i := range crates.ids {
		result += crates.stacks[i].pop()
	}

	result = strings.ReplaceAll(result, "[", "")
	result = strings.ReplaceAll(result, "]", "")

	return result, nil
}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")

	curr := make([]string, 0)
	start := 0
	for i, line := range lines {
		if line == "" {
			start = i
			break
		}
		curr = append(curr, line)
	}

	crates := load(curr)

	for i, line := range lines {
		if i <= start {
			// skip state block
			continue
		}

		if strings.TrimSpace(line) == "" {
			// ignore empty lines
			continue
		}

		cmd, err := parseCommand(line)
		if err != nil {
			return "", fmt.Errorf("error parsing command %d %s", i, line)
		}

		crates.move9001(cmd)
	}

	result := ""
	for i := range crates.ids {
		result += crates.stacks[i].pop()
	}

	result = strings.ReplaceAll(result, "[", "")
	result = strings.ReplaceAll(result, "]", "")

	return result, nil
}
