package common

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	MagicNum = 0xEF3C4D78
)

type Header struct {
	Magic   uint32
	Cmd     uint32
	Version uint16
	BodyLen uint16
}

func EncodeHead(head Header) []byte {
	data := make([]byte, binary.Size(head))

	binary.BigEndian.PutUint32(data, head.Magic)
	index := binary.Size(head.Magic)
	binary.BigEndian.PutUint32(data[index:], head.Cmd)
	index = index + binary.Size(head.Cmd)
	binary.BigEndian.PutUint16(data[index:], head.Version)
	index = index + binary.Size(head.Version)
	binary.BigEndian.PutUint16(data[index:], head.BodyLen)

	return data
}

func DecodeHead(data []byte, head *Header) error {
	if len(data) < binary.Size(*head) {
		msg := fmt.Sprintf("decode head the data must %d, but actual is %d",
			binary.Size(*head), len(data))
		return errors.New(msg)
	}

	head.Magic = binary.BigEndian.Uint32(data)
	index := binary.Size(head.Magic)

	head.Cmd = binary.BigEndian.Uint32(data[index:])
	index = index + binary.Size(head.Cmd)

	head.Version = binary.BigEndian.Uint16(data[index:])
	index = index + binary.Size(head.Version)

	head.BodyLen = binary.BigEndian.Uint16(data[index:])

	return nil
}
