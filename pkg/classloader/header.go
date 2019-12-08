package classloader

const (
	classFileMagic             = 0xcafebabe
	constantUtf8               = 1
	constantInteger            = 3
	constantFloat              = 4
	constantLong               = 5
	constantDouble             = 6
	constantClass              = 7
	constantString             = 8
	constantFieldref           = 9
	constantMethodref          = 10
	constantInterfaceMethodref = 11
	constantNameAndType        = 12
	constantMethodHandle       = 15
	constantMethodType         = 16
	constantInvokeDynamic      = 18
)

type interfaceEntry struct{}
type fieldEntry struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	AttributeInfo   []interface{}
}
type methodEntry struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	AttributeInfo   []interface{}
}
type attributeEntry struct {
	AttributeNameIndex uint16
	AttributeLength    uint32
	Info               []byte
}

type ClassFile struct {
	Magic                       uint32
	MinorVersion                uint16
	MajorVersion                uint16
	ConstantPoolCount           uint16
	AccessFlags                 uint16
	ClassConstantPoolIndex      uint16
	SuperclassConstantPoolIndex uint16
	InterfaceCount              uint16
	FieldCount                  uint16
	MethodCount                 uint16
	AttributeCount              uint16
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
