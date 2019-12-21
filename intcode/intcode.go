package intcode

func Execute(p *Program) {
	NewExecutor(p).Execute()
}
