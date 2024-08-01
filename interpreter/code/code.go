package code

import "fmt"

const (
	OpConstant Opcode = iota
)

type Instructions []byte

type Opcode byte

type Definition struct {
	Name string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition {
	OpConstant: {"OpConstant", []int{2}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}

	return def, nil
}
