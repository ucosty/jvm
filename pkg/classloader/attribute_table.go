package classloader

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

func parseAttributeTable(r io.Reader, h *classHeader) (err error) {
	binary.Read(r, binary.BigEndian, &h.AttributeCount)
	// fmt.Printf("Loading %d attribute table entries\n", h.AttributeCount)
	h.AttributeTable = make([]interface{}, h.AttributeCount)
	for i := 0; i < int(h.AttributeCount); i++ {
		h.AttributeTable[i], err = parseAttributeTableEntry(r, h)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseCodeAttribute(info []byte, h *classHeader) (ca CodeAttribute) {

	infoReader := bytes.NewReader(info)

	binary.Read(infoReader, binary.BigEndian, &ca.MaxStack)
	binary.Read(infoReader, binary.BigEndian, &ca.MaxLocals)
	binary.Read(infoReader, binary.BigEndian, &ca.CodeLength)

	ca.Code = make([]byte, ca.CodeLength)
	binary.Read(infoReader, binary.BigEndian, &ca.Code)

	//fmt.Printf("code:\n%s", hex.Dump(ca.Code))

	// binary.Read(methodReader, binary.BigEndian, &ca.ExceptionTableLength)
	// binary.Read(methodReader, binary.BigEndian, &code.ExceptionTable)

	// binary.Read(methodReader, binary.BigEndian, &ca.AttributesCount)
	// binary.Read(methodReader, binary.BigEndian, &code.AttributeInfo)

	return ca
}

func parseAttributeTableEntry(r io.Reader, h *classHeader) (interface{}, error) {
	var entry attributeEntry
	binary.Read(r, binary.BigEndian, &entry.AttributeNameIndex)
	binary.Read(r, binary.BigEndian, &entry.AttributeLength)
	entry.Info = make([]uint8, entry.AttributeLength)
	binary.Read(r, binary.BigEndian, &entry.Info)

	//fmt.Printf("Attribute: AttributeNameIndex = %d, AttributeName = %s, AttributeLength = %d\n", entry.AttributeNameIndex, getFieldName(entry.AttributeNameIndex, h), entry.AttributeLength)

	attributeName := getFieldName(entry.AttributeNameIndex, h)

	switch attributeName {
	case "Code":
		return parseCodeAttribute(entry.Info, h), nil
	case "SourceFile":
		return nil, nil
	case "InnerClasses":
		return nil, nil
	case "BootstrapMethods":
		return nil, nil

	default:
		return nil, fmt.Errorf("Could not parse attribute %s", attributeName)
	}
}
