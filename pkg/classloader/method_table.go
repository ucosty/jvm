package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
)

func parseMethodTable(r io.Reader, h *classHeader) (err error) {
	if err := binary.Read(r, binary.BigEndian, &h.MethodCount); err != nil {
		return err
	}

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
			return methodEntry{}, err
		}
	}
	return entry, nil
}
