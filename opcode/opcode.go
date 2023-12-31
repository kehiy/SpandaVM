package opcode

var (
	// EXIT is exit OP code.
	EXIT = 0x00

	// ADD adds two int.
	ADD = 0x01

	// MUL multiplies two int.
	MUL = 0x02

	// SUB subtracts two int.
	SUB = 0x03

	// DIV divides two int.
	DIV = 0x04

	// PUSH pushes a new key/value to storage.
	PUSH = 0x05

	// RM removes a key/value in storage.
	RM = 0x06
)

// Opcode is a holder for a single instruction.
// Note that this doesn't take any account of the arguments which might
// be necessary.
type Opcode struct {
	instruction byte
}

// NewOpcode creates a new Opcode.
func NewOpcode(instruction byte) *Opcode {
	o := &Opcode{}
	o.instruction = instruction
	return o
}

// String converts the given Opcode to a string, but again note that it
// doesn't take into account the value.
func (o *Opcode) String() string {
	switch int(o.instruction) {
	case EXIT:
		return "EXIT"
	case ADD:
		return "ADD"
	case MUL:
		return "MUL"
	case SUB:
		return "SUB"
	case DIV:
		return "DIV"
	case PUSH:
		return "PUSH"
	case RM:
		return "RM"
	default:
		return "UNKNOWN_OPCODE"
	}
}

// Value returns the byte-value of the opcode.
func (o *Opcode) Value() byte {
	return (o.instruction)
}
