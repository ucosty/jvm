package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
	"github.com/ucosty/jvm/pkg/jvm"
)

func parseConstantFieldref(r io.Reader) (*jvm.Constant, error) {
	fieldref := jvm.ConstantFieldref{}
	if err := binary.Read(r, binary.BigEndian, &fieldref.ClassIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &fieldref.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: fieldref,
		Type:  fieldref.GetType()}, nil
}

func parseConstantMethodref(r io.Reader) (*jvm.Constant, error) {
	methodref := jvm.ConstantMethodref{}
	if err := binary.Read(r, binary.BigEndian, &methodref.ClassIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &methodref.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: methodref,
		Type:  methodref.GetType()}, nil
}

func parseConstantInterfaceMethodref(r io.Reader) (*jvm.Constant, error) {
	interfaceMethodref := jvm.ConstantInterfaceMethodref{}
	if err := binary.Read(r, binary.BigEndian, &interfaceMethodref.ClassIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &interfaceMethodref.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: interfaceMethodref,
		Type:  interfaceMethodref.GetType()}, nil
}

func parseConstantClass(r io.Reader) (*jvm.Constant, error) {
	class := jvm.ConstantClass{}
	if err := binary.Read(r, binary.BigEndian, &class.NameIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: class,
		Type:  class.GetType()}, nil
}

func parseConstantString(r io.Reader) (*jvm.Constant, error) {
	constantString := jvm.ConstantString{}
	if err := binary.Read(r, binary.BigEndian, &constantString.StringIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: constantString,
		Type:  constantString.GetType()}, nil
}

func parseConstantInteger(r io.Reader) (*jvm.Constant, error) {
	constantInteger := jvm.ConstantInteger{}
	if err := binary.Read(r, binary.BigEndian, &constantInteger.Bytes); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: constantInteger,
		Type:  constantInteger.GetType()}, nil
}

func parseConstantFloat(r io.Reader) (*jvm.Constant, error) {
	constantFloat := jvm.ConstantFloat{}
	if err := binary.Read(r, binary.BigEndian, &constantFloat.Bytes); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: constantFloat,
		Type:  constantFloat.GetType()}, nil
}

func parseConstantLong(r io.Reader) (*jvm.Constant, error) {
	constantLong := jvm.ConstantLong{}
	if err := binary.Read(r, binary.BigEndian, &constantLong.HighBytes); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &constantLong.LowBytes); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: constantLong,
		Type:  constantLong.GetType()}, nil
}

func parseConstantDouble(r io.Reader) (*jvm.Constant, error) {
	constantDouble := jvm.ConstantDouble{}
	if err := binary.Read(r, binary.BigEndian, &constantDouble.HighBytes); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &constantDouble.LowBytes); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: constantDouble,
		Type:  constantDouble.GetType()}, nil
}

func parseConstantNameAndTypeInfo(r io.Reader) (*jvm.Constant, error) {
	nameAndType := jvm.ConstantNameAndType{}
	if err := binary.Read(r, binary.BigEndian, &nameAndType.NameIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &nameAndType.DescriptorIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: nameAndType,
		Type:  nameAndType.GetType()}, nil
}

func parseConstantUTF8Info(r io.Reader) (*jvm.Constant, error) {
	utf8string := jvm.ConstantUtf8{}
	var length uint16
	if err := binary.Read(r, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	utf8string.Bytes = make([]byte, length)
	if err := binary.Read(r, binary.BigEndian, &utf8string.Bytes); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: utf8string,
		Type:  utf8string.GetType()}, nil
}

func parseConstantMethodHandle(r io.Reader) (*jvm.Constant, error) {
	methodHandle := jvm.ConstantMethodHandle{}
	if err := binary.Read(r, binary.BigEndian, &methodHandle.ReferenceKind); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &methodHandle.ReferenceIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: methodHandle,
		Type:  methodHandle.GetType()}, nil
}

func parseConstantMethodType(r io.Reader) (*jvm.Constant, error) {
	methodType := jvm.ConstantMethodType{}
	if err := binary.Read(r, binary.BigEndian, &methodType.DescriptorIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: methodType,
		Type:  methodType.GetType()}, nil
}

func parseConstantInvokeDynamic(r io.Reader) (*jvm.Constant, error) {
	invokeDynamic := jvm.ConstantInvokeDynamic{}
	if err := binary.Read(r, binary.BigEndian, &invokeDynamic.BootstrapMethodAttrIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &invokeDynamic.NameAndTypeIndex); err != nil {
		return nil, err
	}
	return &jvm.Constant{
		Value: invokeDynamic,
		Type:  invokeDynamic.GetType()}, nil
}

func parseConstantPool(r io.Reader, f *ClassFile, c *jvm.Class) (err error) {
	if err := binary.Read(r, binary.BigEndian, &f.ConstantPoolCount); err != nil {
		return err
	}
	for i := 0; i < int(f.ConstantPoolCount-1); i++ {
		entry, err := parseConstantPoolEntry(r)
		c.AddConstant(entry)
		if err != nil {
			return fmt.Errorf("loading pool entry %d: %v", i, err)
		}
	}
	return nil
}

func parseConstantPoolEntry(r io.Reader) (*jvm.Constant, error) {
	var tag uint8
	if err := binary.Read(r, binary.BigEndian, &tag); err != nil {
		return nil, err
	}

	switch tag {
	case constantClass:
		return parseConstantClass(r)
	case constantFieldref:
		return parseConstantFieldref(r)
	case constantMethodref:
		return parseConstantMethodref(r)
	case constantInterfaceMethodref:
		return parseConstantInterfaceMethodref(r)
	case constantString:
		return parseConstantString(r)
	case constantInteger:
		return parseConstantInteger(r)
	case constantFloat:
		return parseConstantFloat(r)
	case constantLong:
		return parseConstantLong(r)
	case constantDouble:
		return parseConstantDouble(r)
	case constantNameAndType:
		return parseConstantNameAndTypeInfo(r)
	case constantUtf8:
		return parseConstantUTF8Info(r)
	case constantMethodHandle:
		return parseConstantMethodHandle(r)
	case constantMethodType:
		return parseConstantMethodType(r)
	case constantInvokeDynamic:
		return parseConstantInvokeDynamic(r)
	}

	return nil, fmt.Errorf("could not parse constant pool entry with tag %d", tag)
}
