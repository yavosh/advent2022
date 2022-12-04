package day4

import (
	"fmt"
	"strconv"
	"strings"
)

type section struct {
	from int
	to   int
}

func toSection(in string) section {
	parts := strings.Split(in, "-")
	from, _ := strconv.Atoi(parts[0])
	to, _ := strconv.Atoi(parts[1])
	return section{from: from, to: to}
}

func toSections(in string) []section {
	defs := strings.Split(in, ",")
	result := make([]section, len(defs))
	for i, r := range defs {
		result[i] = toSection(r)
	}

	return result
}

func toSet(in string) []int {
	parts := strings.Split(in, "-")
	from, _ := strconv.Atoi(parts[0])
	to, _ := strconv.Atoi(parts[1])
	result := make([]int, 0, 10)
	for from <= to {
		result = append(result, from)
		from++
	}

	return result
}

func toSets(in string) [][]int {
	defs := strings.Split(in, ",")
	result := make([][]int, len(defs))

	for i, r := range defs {
		result[i] = toSet(r)
	}

	return result
}

func contains(section1, section2 section) bool {
	if section1.from >= section2.from && section1.to <= section2.to {
		return true
	}

	return false
}

func touches(section1, section2 section) bool {
	if section1.from >= section2.from && section1.from <= section2.to {
		return true
	}

	if section1.to <= section2.to && section1.to >= section2.from {
		return true
	}

	return false
}

func Solution(in string) (string, error) {
	lines := strings.Split(in, "\n")
	overlapsCount := 0
	for i, line := range lines {
		sections := toSections(line)
		if len(sections) != 2 {
			return "", fmt.Errorf("error too many sections in line %d %s", i, line)
		}

		if contains(sections[0], sections[1]) || contains(sections[1], sections[0]) {
			overlapsCount++
		}
	}

	return fmt.Sprintf("%d", overlapsCount), nil
}

func SolutionB(in string) (string, error) {
	lines := strings.Split(in, "\n")
	overlapsCount := 0
	for i, line := range lines {
		sections := toSections(line)
		if len(sections) != 2 {
			return "", fmt.Errorf("error too many sections in line %d %s", i, line)
		}

		if touches(sections[0], sections[1]) || touches(sections[1], sections[0]) {
			overlapsCount++
		}
	}

	return fmt.Sprintf("%d", overlapsCount), nil
}
