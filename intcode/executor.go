package intcode

// Opcode is a translation from integer to action the program wants to take
type Opcode int

const (
	// ADD takes two input positions and puts the sum in the output
	ADD = 1
	// MULTIPLY takes two input positions and puts the product in the output
	MULTIPLY = 2
	// HALT tells the program to stop
	HALT = 99
)

// An Executor takes a program to put into memory, and then can execute the
// instructionns
type Executor struct {
	memory             *Program
	instructionPointer int
}

// NewExecutor builds a new instance of Executor with the specified program
// as the initial state of the memory
func NewExecutor(program *Program) *Executor {
	executor := new(Executor)
	executor.memory = program
	executor.instructionPointer = 0

	return executor
}

// Execute runs the program in memory until it halts
func (e *Executor) Execute() {
	if e.NextOpcode() == HALT {
		// println("halting")
		return
	}

	if e.NextOpcode() == ADD {
		e.Add()
		e.Execute()
		return
	}

	if e.NextOpcode() == MULTIPLY {
		e.Multiply()
		e.Execute()
		return
	}
}

// Add uses the instruction to take values and add them together, then places
// them in the output
func (e *Executor) Add() {
	input0position := e.memory.ValueAt(e.instructionPointer + 1)
	input1position := e.memory.ValueAt(e.instructionPointer + 2)
	outputPosition := e.memory.ValueAt(e.instructionPointer + 3)

	input0value := e.memory.ValueAt(input0position)
	input1value := e.memory.ValueAt(input1position)

	e.memory.ReplaceAt(outputPosition, input0value+input1value)

	e.Step()
}

// Multiply uses the instruction to take values and multiply them together,
// then places the product in the output
func (e *Executor) Multiply() {
	input0position := e.memory.ValueAt(e.instructionPointer + 1)
	input1position := e.memory.ValueAt(e.instructionPointer + 2)
	outputPosition := e.memory.ValueAt(e.instructionPointer + 3)

	input0value := e.memory.ValueAt(input0position)
	input1value := e.memory.ValueAt(input1position)

	e.memory.ReplaceAt(outputPosition, input0value*input1value)

	e.Step()
}

// NextOpcode finds the Opcode at the start of the next instuction to execute
func (e *Executor) NextOpcode() Opcode {
	return Opcode(e.memory.ValueAt(e.instructionPointer))
}

// Step moves the instruction pointer forward
func (e *Executor) Step() {
	e.instructionPointer += 4
}
