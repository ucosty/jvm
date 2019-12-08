package jvm

type Constant struct {
	Type     string
	Value    ConstantValue
	Instance *ClassInstance
}

type ConstantValue interface {
	GetType() string
}

type StringConstant struct {
	StringRef uint16
}

type UTF8Constant struct {
	Value []byte
}

// Constant Type: Class
type ConstantClass struct {
	NameIndex uint16
}

func (c ConstantClass) GetType() string {
	return "CONSTANT_Class"
}

// Constant Type: Fieldref
type ConstantFieldref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantFieldref) GetType() string {
	return "CONSTANT_Fieldref"
}

// Constant Type: Methodref
type ConstantMethodref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantMethodref) GetType() string {
	return "CONSTANT_Methodref"
}

// Constant Type: InterfaceMethodref
type ConstantInterfaceMethodref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c ConstantInterfaceMethodref) GetType() string {
	return "CONSTANT_InterfaceMethodref"
}

// Constant Type: String
type ConstantString struct {
	StringIndex uint16
}

func (c ConstantString) GetType() string {
	return "CONSTANT_String"
}

// Constant Type: Integer
type ConstantInteger struct {
	Bytes uint32
}

func (c ConstantInteger) GetType() string {
	return "CONSTANT_Integer"
}

// Constant Type: Float
type ConstantFloat struct {
	Bytes uint32
}

func (c ConstantFloat) GetType() string {
	return "CONSTANT_Float"
}

// Constant Type: Long
type ConstantLong struct {
	HighBytes uint32
	LowBytes  uint32
}

func (c ConstantLong) GetType() string {
	return "CONSTANT_Long"
}

// Constant Type: Double
type ConstantDouble struct {
	HighBytes uint32
	LowBytes  uint32
}

func (c ConstantDouble) GetType() string {
	return "CONSTANT_Double"
}

// Constant Type: NameAndType
type ConstantNameAndType struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c ConstantNameAndType) GetType() string {
	return "CONSTANT_NameAndType"
}

// Constant Type: Utf8
type ConstantUtf8 struct {
	Bytes []byte
}

func (c ConstantUtf8) GetType() string {
	return "CONSTANT_Utf8"
}

// Constant Type: MethodHandle
type ConstantMethodHandle struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (c ConstantMethodHandle) GetType() string {
	return "CONSTANT_MethodHandle"
}

// Constant Type: MethodType
type ConstantMethodType struct {
	DescriptorIndex uint16
}

func (c ConstantMethodType) GetType() string {
	return "CONSTANT_MethodType"
}

// Constant Type: InvokeDynamic
type ConstantInvokeDynamic struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (c ConstantInvokeDynamic) GetType() string {
	return "CONSTANT_InvokeDynamic"
}
