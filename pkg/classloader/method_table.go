package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
)

func parseMethodTable(r io.Reader, h *classHeader) (err error) {
	binary.Read(r, binary.BigEndian, &h.MethodCount)
	// fmt.Printf("Loading %d Method table entries\n", h.MethodCount)
	h.MethodTable = make([]methodEntry, h.MethodCount)
	for i := 0; i < int(h.MethodCount); i++ {
		h.MethodTable[i], err = parseMethodTableEntry(r, h)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseMethodTableEntry(r io.Reader, h *classHeader) (entry methodEntry, err error) {
	binary.Read(r, binary.BigEndian, &entry.AccessFlags)
	binary.Read(r, binary.BigEndian, &entry.NameIndex)
	binary.Read(r, binary.BigEndian, &entry.DescriptorIndex)
	binary.Read(r, binary.BigEndian, &entry.AttributesCount)

	if int(entry.NameIndex) > len(h.ConstantPoolTable) {
		return entry, fmt.Errorf("invalid name index %d > %d", int(entry.NameIndex), len(h.ConstantPoolTable))
	}

	//fmt.Printf("Method: AccessFlags = %d, NameIndex = %d, Name = %s, DescriptorIndex = %d, DescriptorName = %s, AttributesCount = %d\n", entry.AccessFlags, entry.NameIndex, getFieldName(entry.NameIndex, h), entry.DescriptorIndex, getFieldName(entry.DescriptorIndex, h), entry.AttributesCount)
	entry.AttributeInfo = make([]interface{}, entry.AttributesCount)
	for i := 0; i < int(entry.AttributesCount); i++ {
		entry.AttributeInfo[i], err = parseAttributeTableEntry(r, h)
		if err != nil {
			return methodEntry{}, err
		}
	}
	//fmt.Printf("----\n\n")
	return entry, nil
}
