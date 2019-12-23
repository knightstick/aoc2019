package day4

import (
	"strconv"

	"github.com/knightstick/aoc2019/passwords"
)

func Part1(args []string) int {
	if len(args) != 2 {
		panic("expected exactly two arguments; the min and max")
	}

	min, err := strconv.Atoi(args[0])
	if err != nil {
		panic("could not parse min")
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		panic("could not parse max")
	}

	return len(passwords.ValidPasswords(min, max))
}
