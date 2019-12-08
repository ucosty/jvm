package classloader

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"java-hackery/pkg/jvm"
	"log"
	"os"
	"path"
	"path/filepath"
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

//func getClassName(index uint16, f *ClassFile) string {
//	constClass := f.ConstantPoolTable[index-1].(constantClassEntry)
//	name := f.ConstantPoolTable[constClass.NameIndex-1].(constantUTF8InfoEntry)
//	return string(name.Bytes)
//}

func parseHeader(classBytes []byte) (class *jvm.Class, err error) {
	var header ClassFile
	classReader := bytes.NewReader(classBytes)
	class = jvm.NewClass()

	if err := binary.Read(classReader, binary.BigEndian, &header.Magic); err != nil {
		return nil, err
	}
	if header.Magic != classFileMagic {
		return nil, fmt.Errorf("invalid magic header, expected = %x, found = %x", classFileMagic, header.Magic)
	}

	if err := binary.Read(classReader, binary.BigEndian, &header.MinorVersion); err != nil {
		return nil, err
	}
	if err := binary.Read(classReader, binary.BigEndian, &header.MajorVersion); err != nil {
		return nil, err
	}

	// Read the ConstantPoolTable
	if err := parseConstantPool(classReader, &header, class); err != nil {
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
	if err := parseInterfaceTable(classReader, &header, class); err != nil {
		return nil, err
	}

	// Read the FieldTable
	if err := parseFieldTable(classReader, &header, class); err != nil {
		return nil, err
	}

	// Read the MethodTable
	if err := parseMethodTable(classReader, &header, class); err != nil {
		return nil, err
	}

	// Read the AttributeTable
	if err := parseAttributeTable(classReader, &header, class); err != nil {
		return nil, err
	}

	if class.Name, err = class.GetClassNameConstant(header.ClassConstantPoolIndex); err != nil {
		return nil, err
	}
	if class.Superclass, err = class.GetClassNameConstant(header.SuperclassConstantPoolIndex); err != nil {
		return nil, err
	}

	return class, nil
}

func loadFromFile(filename string) (*jvm.Class, error) {
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

func LoadFromClasspath(dir string, metaspace *jvm.Metaspace) error {
	files, err := filepath.Glob(path.Join(dir, "*.class"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		class, err := loadFromFile(file)
		if err != nil {
			return err
		}
		if err := metaspace.LoadFromStruct(class); err != nil {
			return fmt.Errorf("failed to load class %s: %v", class.Name, err)
		}
	}

	return nil
}
