package jvm

type Method struct {
	Name            string
	Code            []byte
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	AttributesCount uint16
	Attributes      []*Attribute
}
