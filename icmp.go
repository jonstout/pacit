package pacit

import (
	"io"
	"bytes"
	"encoding/binary"
)

type ICMP struct {
	IPv4
	Type uint8
	Code uint8
	Checksum uint16
	Data []byte
}

func (i *ICMP) Len() (n uint16) {
	return i.IPv4.Len() + uint16(4 + len(i.Data))
}

func (i *ICMP) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	var m int64 = 0
	if m, err = buf.ReadFrom(&i.IPv4); int(m) == 0 {
		return
	}
	n += int(m)
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.Code)
	binary.Write(buf, binary.BigEndian, i.Checksum)
	binary.Write(buf, binary.BigEndian, i.Data)
	if n, err = buf.Read(b); n == 0 {
		return
	}
	return n, io.EOF
}

func (i *ICMP) Write(b []byte) (n int, err error) {
	if n, err = i.IPv4.Write(b); n == 0 {
		return
	}
	buf := bytes.NewBuffer(b[n:])
	if err = binary.Read(buf, binary.BigEndian, &i.Type); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &i.Code); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &i.Checksum); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &i.Data); err != nil {
		return
	}
	n += len(i.Data)
	return
}
