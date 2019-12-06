package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
)

func parseFieldTable(r io.Reader, h *classHeader) (err error) {
	binary.Read(r, binary.BigEndian, &h.FieldCount)
	// fmt.Printf("Loading %d field table entries\n", h.FieldCount)
	h.FieldTable = make([]fieldEntry, h.FieldCount)
	for i := 0; i < int(h.FieldCount); i++ {
		h.FieldTable[i], err = parseFieldTableEntry(r, h)
		if err != nil {
			return err
		}
	}
	return nil
}

func getFieldName(index uint16, h *classHeader) string {
	return string(h.ConstantPoolTable[index-1].(constantUTF8InfoEntry).Bytes)
}

func parseFieldTableEntry(r io.Reader, h *classHeader) (entry fieldEntry, err error) {
	binary.Read(r, binary.BigEndian, &entry.AccessFlags)
	binary.Read(r, binary.BigEndian, &entry.NameIndex)
	binary.Read(r, binary.BigEndian, &entry.DescriptorIndex)
	binary.Read(r, binary.BigEndian, &entry.AttributesCount)

	if int(entry.NameIndex) > len(h.ConstantPoolTable) {
		return entry, fmt.Errorf("invalid name index %d > %d", int(entry.NameIndex), len(h.ConstantPoolTable))
	}

	entry.AttributeInfo = make([]interface{}, entry.AttributesCount)
	for i := 0; i < int(entry.AttributesCount); i++ {
		entry.AttributeInfo[i], err = parseAttributeTableEntry(r, h)
		if err != nil {
			return fieldEntry{}, err
		}
	}
	// fmt.Printf("Field: AccessFlags = %d, NameIndex = %d, Name = %s, DescriptorIndex = %d, DescriptorName = %s, AttributesCount = %d\n", entry.AccessFlags, entry.NameIndex, getFieldName(entry.NameIndex, h), entry.DescriptorIndex, getFieldName(entry.DescriptorIndex, h), entry.AttributesCount)
	return entry, nil
}
