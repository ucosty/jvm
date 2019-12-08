package jvm

type Attribute struct {
	Type string
	Value AttributeValue
}

type AttributeValue interface {
	GetType() string
}

type CodeAttribute struct {
	AttributeNameIndex   uint16
	AttributeLength      uint32
	MaxStack             uint16
	MaxLocals            uint16
	CodeLength           uint32
	Code                 []byte
	ExceptionTableLength uint16
	ExceptionTable       []byte
	AttributesCount      uint16
	AttributeInfo        []byte
}

func (c CodeAttribute) GetType() string {
	return "Code"
}
