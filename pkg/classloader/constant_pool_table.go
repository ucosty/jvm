package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
)

func parseConstantFieldMethodInterfaceRef(r io.Reader) constantFieldMethodInterfaceRefEntry {
	var classIndex, nameAndTypeIndex uint16
	binary.Read(r, binary.BigEndian, &classIndex)
	binary.Read(r, binary.BigEndian, &nameAndTypeIndex)
	return constantFieldMethodInterfaceRefEntry{
		Tag:              constantMethodref,
		ClassIndex:       classIndex,
		NameAndTypeIndex: nameAndTypeIndex}
}

func parseConstantClass(r io.Reader) constantClassEntry {
	var nameIndex uint16
	binary.Read(r, binary.BigEndian, &nameIndex)
	return constantClassEntry{
		Tag:       constantClass,
		NameIndex: nameIndex}
}

func parseConstantString(r io.Reader) constantStringEntry {
	var stringIndex uint16
	binary.Read(r, binary.BigEndian, &stringIndex)
	return constantStringEntry{
		Tag:         constantClass,
		StringIndex: stringIndex}
}

func parseConstantInteger(r io.Reader) ConstantIntegerEntry {
	var bytes uint32
	binary.Read(r, binary.BigEndian, &bytes)
	return ConstantIntegerEntry{
		Tag:   constantClass,
		Bytes: bytes}
}

func parseConstantFloat(r io.Reader) constantFloatEntry {
	var bytes uint32
	binary.Read(r, binary.BigEndian, &bytes)
	return constantFloatEntry{
		Tag:   constantClass,
		Bytes: bytes}
}

func parseConstantLong(r io.Reader) constantLongEntry {
	var highBytes, lowBytes uint32
	binary.Read(r, binary.BigEndian, &highBytes)
	binary.Read(r, binary.BigEndian, &lowBytes)
	return constantLongEntry{
		Tag:       constantClass,
		HighBytes: highBytes,
		LowBytes:  lowBytes}
}

func parseConstantDouble(r io.Reader) constantDoubleEntry {
	var highBytes, lowBytes uint32
	binary.Read(r, binary.BigEndian, &highBytes)
	binary.Read(r, binary.BigEndian, &lowBytes)
	return constantDoubleEntry{
		Tag:       constantClass,
		HighBytes: highBytes,
		LowBytes:  lowBytes}
}

func parseConstantNameAndTypeInfo(r io.Reader) constantNameAndTypeInfoEntry {
	var nameIndex, descriptorIndex uint16
	binary.Read(r, binary.BigEndian, &nameIndex)
	binary.Read(r, binary.BigEndian, &descriptorIndex)
	return constantNameAndTypeInfoEntry{
		Tag:             constantNameAndType,
		NameIndex:       nameIndex,
		DescriptorIndex: descriptorIndex}
}

func parseConstantUTF8Info(r io.Reader) constantUTF8InfoEntry {
	var length uint16
	binary.Read(r, binary.BigEndian, &length)
	bytes := make([]byte, length)
	binary.Read(r, binary.BigEndian, &bytes)
	return constantUTF8InfoEntry{
		Tag:    constantUtf8,
		Length: length,
		Bytes:  bytes}
}

func parseConstantMethodHandle(r io.Reader) constantMethodHandleInfoEntry {
	var referenceKind uint8
	var referenceIndex uint16
	binary.Read(r, binary.BigEndian, &referenceKind)
	binary.Read(r, binary.BigEndian, &referenceIndex)
	return constantMethodHandleInfoEntry{
		Tag:            constantMethodHandle,
		ReferenceKind:  referenceKind,
		ReferenceIndex: referenceIndex}
}

func parseConstantMethodType(r io.Reader) constantMethodTypeInfoEntry {
	var descriptorIndex uint16
	binary.Read(r, binary.BigEndian, &descriptorIndex)
	return constantMethodTypeInfoEntry{
		Tag:             constantMethodType,
		DescriptorIndex: descriptorIndex,
	}
}

func parseConstantInvokeDynamic(r io.Reader) constantInvokeDynamicInfoEntry {
	var bootstrapMethodAttrIndex, nameAndTypeIndex uint16
	binary.Read(r, binary.BigEndian, &bootstrapMethodAttrIndex)
	binary.Read(r, binary.BigEndian, &nameAndTypeIndex)
	return constantInvokeDynamicInfoEntry{
		Tag:                      constantInvokeDynamic,
		BootstrapMethodAttrIndex: bootstrapMethodAttrIndex,
		NameAndTypeIndex:         nameAndTypeIndex,
	}
}

func parseConstantPool(r io.Reader, h *classHeader) (err error) {
	binary.Read(r, binary.BigEndian, &h.ConstantPoolCount)
	// fmt.Printf("Loading %d constant pool table entries\n", h.ConstantPoolCount)
	h.ConstantPoolTable = make([]constantPoolEntry, h.ConstantPoolCount)
	for i := 0; i < int(h.ConstantPoolCount-1); i++ {
		h.ConstantPoolTable[i], err = parseConstantPoolEntry(r)
		if err != nil {
			return fmt.Errorf("loading pool entry %d: %v", i, err)
		}
	}
	return nil
}

func parseConstantPoolEntry(r io.Reader) (constantPoolEntry, error) {
	var tag uint8
	binary.Read(r, binary.BigEndian, &tag)

	switch tag {
	case constantClass:
		return parseConstantClass(r), nil
	case constantFieldref:
		return parseConstantFieldMethodInterfaceRef(r), nil
	case constantMethodref:
		return parseConstantFieldMethodInterfaceRef(r), nil
	case constantInterfaceMethodref:
		return parseConstantFieldMethodInterfaceRef(r), nil
	case constantString:
		return parseConstantString(r), nil
	case constantInteger:
		return parseConstantInteger(r), nil
	case constantFloat:
		return parseConstantFloat(r), nil
	case constantLong:
		return parseConstantLong(r), nil
	case constantDouble:
		return parseConstantDouble(r), nil
	case constantNameAndType:
		return parseConstantNameAndTypeInfo(r), nil
	case constantUtf8:
		return parseConstantUTF8Info(r), nil
	case constantMethodHandle:
		return parseConstantMethodHandle(r), nil
	case constantMethodType:
		return parseConstantMethodType(r), nil
	case constantInvokeDynamic:
		foo := parseConstantInvokeDynamic(r)
		return foo, nil
		//return parseConstantInvokeDynamic(r), nil
	}

	return nil, fmt.Errorf("could not parse constant pool entry with tag %d", tag)
}
