package jvm

import (
	"fmt"
)

type Class struct {
	Name            string
	Superclass      string
	Methods         map[string]*Method
	Constants       []*Constant
	ClassInterfaces []*ClassInterface
	Attributes      []*Attribute
	Fields          []*Field
}

func NewClass() *Class {
	return &Class{
		Methods:   make(map[string]*Method),
		Constants: make([]*Constant, 1),
	}
}

func (c *Class) Invoke(methodName string, args ...interface{}) error {
	if _, exists := c.Methods[methodName]; !exists {
		return fmt.Errorf("method not found: %s in class %s", methodName, c.Name)
	}
	method := c.Methods[methodName]

	// get the code attribute
	codeAttribute := method.Attributes[0].Value.(*CodeAttribute)
	//asm.Disassemble(codeAttribute.Code)

	baseFrame := NewStackFrame()

	if err := baseFrame.PushOperands(args); err != nil {
		return err
	}
	return Execute(baseFrame, codeAttribute.Code, c)
}

func (c *Class) AddAttribute(attribute *Attribute) {
	c.Attributes = append(c.Attributes, attribute)
}

func (c *Class) GetUTF8Constant(index uint16) (string, error) {
	if int(index) > len(c.Constants) {
		return "", fmt.Errorf("invalid constant pool index %d > %d", index, len(c.Constants))
	}
	return string(c.Constants[index].Value.(ConstantUtf8).Bytes), nil
}

func (c *Class) GetClassNameConstant(index uint16) (string, error) {
	if int(index) > len(c.Constants) {
		return "", fmt.Errorf("invalid constant pool index %d > %d", index, len(c.Constants))
	}

	constantClass := c.Constants[index].Value.(ConstantClass)
	if int(constantClass.NameIndex) > len(c.Constants) {
		return "", fmt.Errorf("invalid constant pool index %d > %d", index, len(c.Constants))
	}
	return c.GetUTF8Constant(constantClass.NameIndex)
}


func (c *Class) GetNameAndType(index uint16) (name string, descriptor string, err error) {
	if c.Constants[index].Type != "CONSTANT_NameAndType" {
		return "", "", fmt.Errorf("expected type CONSTANT_NameAndType, found %s", c.Constants[index].Type)
	}

	nameAndType := c.Constants[index].Value.(ConstantNameAndType)
	if name, err = c.GetUTF8Constant(nameAndType.NameIndex); err != nil {
		return "", "", err
	}
	if descriptor, err = c.GetUTF8Constant(nameAndType.DescriptorIndex); err != nil {
		return "", "", err
	}
	return name, descriptor, nil
}

func (c *Class) AddMethod(method *Method) {
	//fmt.Printf("Adding method %s\n", method.Name)
	c.Methods[method.Name] = method
}

func (c *Class) AddConstant(constant *Constant) {
	c.Constants = append(c.Constants, constant)
}

func (c *Class) AddInterface(classInterface *ClassInterface) {
	c.ClassInterfaces = append(c.ClassInterfaces, classInterface)
}

func (c *Class) AddField(field *Field) {
	c.Fields = append(c.Fields, field)
}
