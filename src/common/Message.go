package common

import (
	"io"
	"encoding/binary"
	"bytes"
	"fmt"
)

func WriteHeader(len int32, seq int32, cmd byte, version byte, buffer io.Writer)  {
	binary.Write(buffer, binary.BigEndian, len)
	binary.Write(buffer, binary.BigEndian, seq)

	t := []byte{cmd, version, 0, 0}
	buffer.Write(t)
}

func ReadHeader(buff []byte) *Header {
	var length int32
	var seq int32
	var cmd int8
	var version int8
	buffer := bytes.NewBuffer(buff)
	binary.Read(buffer, binary.BigEndian, &length)
	binary.Read(buffer, binary.BigEndian, &seq)
	binary.Read(buffer, binary.BigEndian, &cmd)
	binary.Read(buffer, binary.BigEndian, &version)
	return &Header{
		length:length,
		seq:seq,
		cmd:cmd,
		version:version,
	}
}

const HeaderLength = 12

type Header struct {
	length int32
	seq int32
	cmd int8
	version int8
}

func NewHeader(len int32, seq int32, cmd int8, version int8) *Header {
	return &Header{
		length:len,
		seq:seq,
		cmd:cmd,
		version:version,
	}
}

func (header *Header)Description() {
	fmt.Printf("{length = %d, seq = %d, cmd = %d, version = %d}\n", header.length, header.seq, header.cmd, header.version)
}

func (header *Header)ToBytes() []byte {
	buffer := new(bytes.Buffer)
	WriteHeader(header.length, header.seq, byte(header.cmd), byte(header.version), buffer)

	return buffer.Bytes()
}








