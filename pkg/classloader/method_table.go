package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
	"java-hackery/pkg/jvm"
)

func parseMethodTable(r io.Reader, h *ClassFile, c *jvm.Class) (err error) {
	if err := binary.Read(r, binary.BigEndian, &h.MethodCount); err != nil {
		return err
	}
	for i := 0; i < int(h.MethodCount); i++ {
		method, err := parseMethodTableEntry(r, h, c)
		if err != nil {
			return err
		}
		c.AddMethod(method)
	}
	return nil
}

func parseMethodTableEntry(r io.Reader, h *ClassFile, c *jvm.Class) (method *jvm.Method, err error) {
	method = &jvm.Method{}
	if err := binary.Read(r, binary.BigEndian, &method.AccessFlags); err != nil {
		return method, err
	}
	if err := binary.Read(r, binary.BigEndian, &method.NameIndex); err != nil {
		return method, err
	}
	if err := binary.Read(r, binary.BigEndian, &method.DescriptorIndex); err != nil {
		return method, err
	}
	if err := binary.Read(r, binary.BigEndian, &method.AttributesCount); err != nil {
		return method, err
	}

	if int(method.NameIndex) > len(c.Constants) {
		return method, fmt.Errorf("invalid name index %d > %d", int(method.NameIndex), len(c.Constants))
	}
	if method.Name, err = c.GetUTF8Constant(method.NameIndex); err != nil {
		return nil, err
	}

	method.Attributes = make([]*jvm.Attribute, method.AttributesCount)
	for i := 0; i < int(method.AttributesCount); i++ {
		 method.Attributes[i], err = parseAttributeTableEntry(r, h, c)
		if err != nil {
			return nil, err
		}
	}

	return method, nil
}
