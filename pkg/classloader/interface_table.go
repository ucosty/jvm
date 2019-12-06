package classloader

import (
	"encoding/binary"
	"io"
)

func parseInterfaceTable(r io.Reader, h *classHeader) (err error) {
	binary.Read(r, binary.BigEndian, &h.InterfaceCount)
	// fmt.Printf("Loading %d Interface table entries\n", h.InterfaceCount)
	h.InterfaceTable = make([]interfaceEntry, h.InterfaceCount)
	for i := 0; i < int(h.InterfaceCount); i++ {
		h.InterfaceTable[i], err = parseInterfaceTableEntry(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseInterfaceTableEntry(r io.Reader) (interfaceEntry, error) {

	return interfaceEntry{}, nil
}