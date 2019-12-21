package day2

import (
	"github.com/knightstick/aoc2019/intcode"
)

// Part1 restores the gravity assist program (from the input) to the "1202
// program alarm" state, and then runs the program. The return value is the
// value at position 0 after the program halts.
func Part1(args []string) int {
	if len(args) != 1 {
		panic("expected 1 argument")
	}

	gravityAssistProgramString := args[0]
	gravityAssistProgram := intcode.NewProgram(gravityAssistProgramString)

	restoreTo1202ProgramAlarmState(gravityAssistProgram)
	output := intcode.Execute(gravityAssistProgram)

	return output[0]
}

func restoreTo1202ProgramAlarmState(program *intcode.Program) {
	provideInput(program, 12, 1)
}

func provideInput(program *intcode.Program, noun, verb int) {
	program.ReplaceAt(1, noun)
	program.ReplaceAt(2, verb)
}
