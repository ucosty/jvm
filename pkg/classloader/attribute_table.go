package classloader

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"github.com/ucosty/jvm/pkg/jvm"
)

func parseAttributeTable(r io.Reader, h *ClassFile, c *jvm.Class) (err error) {
	if err = binary.Read(r, binary.BigEndian, &h.AttributeCount); err != nil {
		return err
	}
	for i := 0; i < int(h.AttributeCount); i++ {
		attribute, err := parseAttributeTableEntry(r, h, c)
		if err != nil {
			return err
		}
		c.AddAttribute(attribute)
	}
	return nil
}

func parseCodeAttribute(info []byte, h *ClassFile) (attribute *jvm.Attribute, err error) {
	codeAttribute := &jvm.CodeAttribute{}
	infoReader := bytes.NewReader(info)
	if err := binary.Read(infoReader, binary.BigEndian, &codeAttribute.MaxStack); err != nil {
		return nil, err
	}
	if err := binary.Read(infoReader, binary.BigEndian, &codeAttribute.MaxLocals); err != nil {
		return nil, err
	}
	if err := binary.Read(infoReader, binary.BigEndian, &codeAttribute.CodeLength); err != nil {
		return nil, err
	}

	codeAttribute.Code = make([]byte, codeAttribute.CodeLength)
	if err := binary.Read(infoReader, binary.BigEndian, &codeAttribute.Code); err != nil {
		return nil, err
	}

	if err := binary.Read(infoReader, binary.BigEndian, &codeAttribute.ExceptionTableLength); err != nil {
		return nil, err
	}
	// binary.Read(methodReader, binary.BigEndian, &code.ExceptionTable)

	//if err := binary.Read(infoReader, binary.BigEndian, &codeAttribute.AttributesCount); err != nil {
	//	return nil, err
	//}
	// binary.Read(methodReader, binary.BigEndian, &code.Attributes)

	return &jvm.Attribute{
		Type:  codeAttribute.GetType(),
		Value: codeAttribute,
	}, nil
}

func parseAttributeTableEntry(r io.Reader, h *ClassFile, c *jvm.Class) (*jvm.Attribute, error) {
	var entry attributeEntry
	if err := binary.Read(r, binary.BigEndian, &entry.AttributeNameIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &entry.AttributeLength); err != nil {
		return nil, err
	}
	entry.Info = make([]uint8, entry.AttributeLength)
	if err := binary.Read(r, binary.BigEndian, &entry.Info); err != nil {
		return nil, err
	}

	attributeName, err := c.GetUTF8Constant(entry.AttributeNameIndex)
	if err != nil {
		return nil, err
	}

	switch attributeName {
	case "Code":
		return parseCodeAttribute(entry.Info, h)
	case "SourceFile":
		return nil, nil
	case "InnerClasses":
		return nil, nil
	case "BootstrapMethods":
		return nil, nil

	default:
		return nil, fmt.Errorf("could not parse attribute %s", attributeName)
	}
}
