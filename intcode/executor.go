package intcode

type Opcode int

const (
	ADD      = 1
	MULTIPLY = 2
	HALT     = 99
)

type Executor struct {
	program  *Program
	position int
}

func NewExecutor(program *Program) *Executor {
	executor := new(Executor)
	executor.program = program
	executor.position = 0

	return executor
}

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

func (e *Executor) Add() {
	input0position := e.program.ValueAt(e.position + 1)
	input1position := e.program.ValueAt(e.position + 2)
	outputPosition := e.program.ValueAt(e.position + 3)

	input0value := e.program.ValueAt(input0position)
	input1value := e.program.ValueAt(input1position)

	e.program.ReplaceAt(outputPosition, input0value+input1value)

	e.Step()
}

func (e *Executor) Multiply() {
	input0position := e.program.ValueAt(e.position + 1)
	input1position := e.program.ValueAt(e.position + 2)
	outputPosition := e.program.ValueAt(e.position + 3)

	input0value := e.program.ValueAt(input0position)
	input1value := e.program.ValueAt(input1position)

	e.program.ReplaceAt(outputPosition, input0value*input1value)

	e.Step()
}

func (e *Executor) NextOpcode() Opcode {
	return Opcode(e.program.ValueAt(e.position))
}

func (e *Executor) Step() {
	e.position += 4
}
