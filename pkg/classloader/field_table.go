package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
	"github.com/ucosty/jvm/pkg/jvm"
)

func parseFieldTable(r io.Reader, h *ClassFile, c *jvm.Class) (err error) {
	if err := binary.Read(r, binary.BigEndian, &h.FieldCount); err != nil {
		return err
	}

	for i := 0; i < int(h.FieldCount); i++ {
		field, err := parseFieldTableEntry(r, h, c)
		if err != nil {
			return err
		}
		c.AddField(field)
	}
	return nil
}

func parseFieldTableEntry(r io.Reader, h *ClassFile, c *jvm.Class) (field *jvm.Field, err error) {
	field = &jvm.Field{}

	if err := binary.Read(r, binary.BigEndian, &field.AccessFlags); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &field.NameIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &field.DescriptorIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &field.AttributesCount); err != nil {
		return nil, err
	}

	if int(field.NameIndex) > len(c.Constants) {
		return nil, fmt.Errorf("invalid name index %d > %d", int(field.NameIndex), len(c.Constants))
	}

	field.Attributes = make([]*jvm.Attribute, field.AttributesCount)
	for i := 0; i < int(field.AttributesCount); i++ {
		field.Attributes[i], err = parseAttributeTableEntry(r, h, c)
		if err != nil {
			return nil, err
		}
	}

	return field, nil
}
