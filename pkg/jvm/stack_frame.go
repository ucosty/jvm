package jvm

type StackFrame struct {
	Parent       *StackFrame
	Class        *Class
	IP           int           // Instruction Pointer
	OperandStack []interface{} // Operand Stack
}

func NewStackFrame() *StackFrame {
	return &StackFrame{
		Parent:       nil,
		Class:        nil,
		IP:           0,
		OperandStack: make([]interface{}, 0),
	}
}
