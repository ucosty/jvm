package jvm

type Field struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []*Attribute

	// Yeah...
	ValueInt       int
	ValueLong      int64
	ValueFloat     float32
	ValueDouble    float64
	ValueReference *ClassInstance
}

func (f *Field) GetValue() {
}
