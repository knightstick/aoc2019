package intcode

func Execute(p *Program) []int {
	e := NewExecutor(p)
	e.Execute()
	return e.memory
}
