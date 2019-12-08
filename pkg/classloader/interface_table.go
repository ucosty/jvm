package classloader

import (
	"encoding/binary"
	"fmt"
	"io"
	"java-hackery/pkg/jvm"
)

func parseInterfaceTable(r io.Reader, h *ClassFile, c *jvm.Class) (err error) {
	if err := binary.Read(r, binary.BigEndian, &h.InterfaceCount); err != nil {
		return err
	}

	for i := 0; i < int(h.InterfaceCount); i++ {
		classInterface, err := parseInterfaceTableEntry(r)
		if err != nil {
			return err
		}

		fmt.Printf("Got Class Interface: %v\n", classInterface)
	}
	return nil
}

func parseInterfaceTableEntry(r io.Reader) (interfaceEntry, error) {
	return interfaceEntry{}, nil
}
