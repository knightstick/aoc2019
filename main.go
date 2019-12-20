package main

import (
	"fmt"
	"os"

	"github.com/knightstick/aoc2019/day1"
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
	}
}
