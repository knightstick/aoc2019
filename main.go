package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/knightstick/aoc2019/day1"
	"github.com/knightstick/aoc2019/day2"
	"github.com/knightstick/aoc2019/day3"
)

var dispatch = map[string]func([]string) int{
	"1-1": day1.Part1,
	"1-2": day1.Part1,
	"2-1": day2.Part1,
	"2-2": day2.Part2,
	"3-1": day3.Part1,
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		println("usage: aoc2019 1")
		return
	}

	dispatchFunc := dispatch[args[0]]
	if dispatchFunc != nil {
		fmt.Println(dispatchFunc(args[1:]))
	} else {
		fmt.Printf("exercise not found: %s\n", args[0])
		fmt.Printf("available exercises: %v", keysString(dispatch))
	}
}

func keys(aMap map[string]func([]string) int) []string {
	keys := make([]string, 0, len(aMap))
	for k := range aMap {
		keys = append(keys, k)
	}
	return keys
}

func keysString(aMap map[string]func([]string) int) string {
	return strings.Join(keys(aMap), ", ")
}
