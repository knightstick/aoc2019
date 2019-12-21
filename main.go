package main

import (
	"fmt"
	"os"

	"github.com/knightstick/aoc2019/day1"
	"github.com/knightstick/aoc2019/day2"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("usage: aoc2019 1")
		return
	}

	switch args[0] {
	case "1-1":
		fmt.Println(day1.Part1(args[1:]))
	case "1-2":
		fmt.Println(day1.Part2(args[1:]))
	case "2-1":
		fmt.Println(day2.Part1(args[1:]))
	}
}
