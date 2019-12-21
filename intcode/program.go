package intcode

import "strings"

import "strconv"

// An Program represents the computer programs run on the Elf ship
type Program struct {
	Values []int
}

// NewProgram constructs a new Program
func NewProgram(string string) *Program {
	ip := new(Program)
	ip.Values = valuesFromString(string)

	return ip
}

// ValueAt gets the value at the specified position
func (ip *Program) ValueAt(position int) int {
	return 0
}

// ReplaceAt overwrites the value at the specified position with the new value
// supplied
func (ip *Program) ReplaceAt(position, value int) {

}

func valuesFromString(string string) (values []int) {
	split := strings.Split(string, ",")

	for _, stringValue := range split {
		value, _ := strconv.Atoi(stringValue)
		values = append(values, value)
	}

	return
}
