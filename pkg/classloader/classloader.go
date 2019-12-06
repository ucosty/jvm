package classloader

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"log"
	"os"
)

func readClassFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size = stats.Size()
	fileBytes := make([]byte, size)

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(fileBytes)

	return fileBytes, err
}

func getClassName(index uint16, h *classHeader) string {
	constClass := h.ConstantPoolTable[index-1].(constantClassEntry)
	name := h.ConstantPoolTable[constClass.NameIndex-1].(constantUTF8InfoEntry)
	return string(name.Bytes)
}

type JavaClass struct {
	Name         string
	Superclass   string
	Header       *classHeader
	Methods      map[string]methodEntry
	ConstantPool []constantPoolEntry
}

func createMethodMap(methods []methodEntry, h *classHeader) map[string]methodEntry {
	methodMap := make(map[string]methodEntry)
	for _, method := range methods {
		methodMap[getFieldName(method.NameIndex, h)] = method
	}
	return methodMap
}

func parseHeader(classBytes []byte) (class *JavaClass, err error) {
	var header classHeader
	classReader := bytes.NewReader(classBytes)

	if err := binary.Read(classReader, binary.BigEndian, &header.Magic); err != nil {
		return nil, err
	}
	if err := binary.Read(classReader, binary.BigEndian, &header.MinorVersion); err != nil {
		return nil, err
	}
	if err := binary.Read(classReader, binary.BigEndian, &header.MajorVersion); err != nil {
		return nil, err
	}

	// Read the ConstantPoolTable
	if err := parseConstantPool(classReader, &header); err != nil {
		return nil, err
	}

	if err := binary.Read(classReader, binary.BigEndian, &header.AccessFlags); err != nil {
		return nil, err
	}
	if err := binary.Read(classReader, binary.BigEndian, &header.ClassConstantPoolIndex); err != nil {
		return nil, err
	}
	if err := binary.Read(classReader, binary.BigEndian, &header.SuperclassConstantPoolIndex); err != nil {
		return nil, err
	}

	// Read the InterfaceTable
	if err := parseInterfaceTable(classReader, &header); err != nil {
		return nil, err
	}

	// Read the FieldTable
	if err := parseFieldTable(classReader, &header); err != nil {
		return nil, err
	}

	// Read the MethodTable
	if err := parseMethodTable(classReader, &header); err != nil {
		return nil, err
	}

	// Read the AttributeTable
	if err := parseAttributeTable(classReader, &header); err != nil {
		return nil, err
	}

	class = &JavaClass{
		Name:         getClassName(header.ClassConstantPoolIndex, &header),
		Superclass:   getClassName(header.SuperclassConstantPoolIndex, &header),
		Header:       &header,
		ConstantPool: header.ConstantPoolTable,
		Methods:      createMethodMap(header.MethodTable, &header)}
	return class, nil
}


func LoadFromFile(filename string) (*JavaClass, error) {
	classBytes, err := readClassFile(filename)
	if err != nil {
		log.Fatal("Could not read class file: ", err)
	}

	class, err := parseHeader(classBytes)
	if err != nil {
		log.Fatal("Could not parse class file: ", err)
	}

	return class, nil
}