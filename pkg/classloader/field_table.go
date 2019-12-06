package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
)

func parseFieldTable(r io.Reader, h *classHeader) (err error) {
	if err := binary.Read(r, binary.BigEndian, &h.FieldCount); err != nil {
		return err
	}

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
	if err := binary.Read(r, binary.BigEndian, &entry.AccessFlags); err != nil {
		return entry, err
	}
	if err := binary.Read(r, binary.BigEndian, &entry.NameIndex); err != nil {
		return entry, err
	}
	if err := binary.Read(r, binary.BigEndian, &entry.DescriptorIndex); err != nil {
		return entry, err
	}
	if err := binary.Read(r, binary.BigEndian, &entry.AttributesCount); err != nil {
		return entry, err
	}

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

	return entry, nil
}
