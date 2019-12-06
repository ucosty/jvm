package classloader

var (
	javaMagic = []byte{0xca, 0xfe, 0xba, 0xbe}
)

const (
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


type constantPoolEntry interface {
	getTag() uint8
}

type constantFieldMethodInterfaceRefEntry struct {
	Tag              uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c constantFieldMethodInterfaceRefEntry) getTag() uint8 {
	return constantFieldref
}

type constantClassEntry struct {
	Tag       uint8
	NameIndex uint16
}

func (c constantClassEntry) getTag() uint8 {
	return constantClass
}

type constantStringEntry struct {
	Tag         uint8
	StringIndex uint16
}

func (c constantStringEntry) getTag() uint8 {
	return constantString
}

type ConstantIntegerEntry struct {
	Tag   uint8
	Bytes uint32
}
type constantFloatEntry ConstantIntegerEntry

func (c ConstantIntegerEntry) getTag() uint8 {
	return constantInteger
}

func (c constantFloatEntry) getTag() uint8 {
	return constantFloat
}

type constantLongEntry struct {
	Tag       uint8
	HighBytes uint32
	LowBytes  uint32
}
type constantDoubleEntry constantLongEntry

func (c constantLongEntry) getTag() uint8 {
	return constantLong
}

func (c constantDoubleEntry) getTag() uint8 {
	return constantDouble
}

type constantNameAndTypeInfoEntry struct {
	Tag             uint8
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c constantNameAndTypeInfoEntry) getTag() uint8 {
	return constantNameAndType
}

type constantUTF8InfoEntry struct {
	Tag    uint8
	Length uint16
	Bytes  []uint8
}

func (c constantUTF8InfoEntry) getTag() uint8 {
	return constantUtf8
}

type constantMethodHandleInfoEntry struct {
	Tag            uint8
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (c constantMethodHandleInfoEntry) getTag() uint8 {
	return constantMethodHandle
}

type constantMethodTypeInfoEntry struct {
	Tag             uint8
	DescriptorIndex uint16
}

func (c constantMethodTypeInfoEntry) getTag() uint8 {
	return constantMethodType
}

type constantInvokeDynamicInfoEntry struct {
	Tag                      uint8
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (c constantInvokeDynamicInfoEntry) getTag() uint8 {
	return constantInvokeDynamic
}

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

type classHeader struct {
	Magic                       uint32
	MinorVersion                uint16
	MajorVersion                uint16
	ConstantPoolCount           uint16
	ConstantPoolTable           []constantPoolEntry
	AccessFlags                 uint16
	ClassConstantPoolIndex      uint16
	SuperclassConstantPoolIndex uint16
	InterfaceCount              uint16
	InterfaceTable              []interfaceEntry
	FieldCount                  uint16
	FieldTable                  []fieldEntry
	MethodCount                 uint16
	MethodTable                 []methodEntry
	AttributeCount              uint16
	AttributeTable              []interface{}
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
