package intcode

// An Program represents the computer programs run on the Elf ship
type Program struct{}

// NewProgram constructs a new Program
func NewProgram(string string) *Program {
	ip := new(Program)

	return ip
}

// ValueAt gets the value at the specified position
func (ip *Program) ValueAt(position int) int {
	return 0
}

func (ip *Program) ReplaceAt(position, value int) {

}

func Execute(p *Program) {

}
