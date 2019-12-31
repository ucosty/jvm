package jvm

import (
	"fmt"
)

type Metaspace struct {
	Classes map[string]*Class
	Heap    []*ClassInstance
}

type ClassInstance struct {
	Class *Class
}

func NewMetaspace() *Metaspace {
	return &Metaspace{
		Heap:    make([]*ClassInstance, 0),
		Classes: make(map[string]*Class)}
}

func (m *Metaspace) GetClass(name string) (*Class, error) {
	if _, exists := m.Classes[name]; !exists {
		return nil, fmt.Errorf("class %s not found", name)
	}
	return m.Classes[name], nil
}

func (m *Metaspace) LoadFromStruct(class *Class) error {
	if _, exists := m.Classes[class.Name]; exists {
		return fmt.Errorf("class %s already loaded", class.Name)
	}
	fmt.Printf("Loading class %s\n", class.Name)
	m.Classes[class.Name] = class
	return nil
}

func (m *Metaspace) InitClasses() error {
	for k, _ := range m.Classes {
		class := m.Classes[k]
		if err := m.loadConstants(class); err != nil {
			return fmt.Errorf("could not load constants for %s: %v", class.Name, err)
		}
	}
	return nil
}

func (m *Metaspace) InstanceClass(class string, name string, args... interface{}) (*ClassInstance, error) {
	if _, exists := m.Classes[class]; !exists {
		return nil, fmt.Errorf("class not found %s", class)
	}
	fmt.Printf("Instancing class %s\n", class)
	return &ClassInstance{
		Class: m.Classes[class]}, nil
}

func (m *Metaspace) loadConstants(c *Class) error {
	fmt.Printf("Loading constants for class %s\n", c.Name)
	for i, constant := range c.Constants {
		if constant == nil {
			continue
		}

		switch constant.Type {
		case "CONSTANT_String":
			constString := constant.Value.(ConstantString)
			utf8value, err := c.GetUTF8Constant(constString.StringIndex)
			if err != nil {
				return err
			}
			fmt.Printf("Creating string object for constant %d: %s\n", i, utf8value)
			instance, err := m.InstanceClass("java/lang/String", fmt.Sprintf(c.Name, i))
			if err != nil {
				return fmt.Errorf("could not load constant: %v", err)
			}
			m.Heap = append(m.Heap, instance)

			// Create the back reference
			c.Constants[i].Instance = instance
		}
	}
	return nil
}

func (m *Metaspace) DumpHeap() {
	for i, instance := range m.Heap {
		fmt.Printf("%d: class = %s\n", i, instance.Class.Name)
	}
}