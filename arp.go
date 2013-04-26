package pacit

import (
	"io"
	"net"
	"bytes"
	"encoding/binary"
)

type ARP struct {
	HWType uint16
	ProtoType uint16
	HWLength uint8
	ProtoLength uint8
	Operation uint16
	HWSrc net.HardwareAddr
	IPSrc net.IP
	HWDst net.HardwareAddr
	IPDst net.IP
}

func (a *ARP) Len() (n uint16) {
	n += 8
	n += uint16(a.HWLength*2 + a.ProtoLength*2)
	return
}

func (a *ARP) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, a.HWType)
	binary.Write(buf, binary.BigEndian, a.ProtoType)
	binary.Write(buf, binary.BigEndian, a.HWLength)
	binary.Write(buf, binary.BigEndian, a.ProtoLength)
	binary.Write(buf, binary.BigEndian, a.Operation)
	binary.Write(buf, binary.BigEndian, a.HWSrc)
	binary.Write(buf, binary.BigEndian, a.IPSrc)
	binary.Write(buf, binary.BigEndian, a.HWDst)
	binary.Write(buf, binary.BigEndian, a.IPDst)
	n, err = buf.Read(b)
	return n, io.EOF
}

func (a *ARP) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &a.HWType); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.ProtoType); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &a.HWLength); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &a.ProtoLength); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &a.Operation); err != nil {
		return
	}
	n += 2
	a.HWSrc = make([]byte, 6)
	if err = binary.Read(buf, binary.BigEndian, &a.HWSrc); err != nil {
		return
	}
	n += 6
	a.IPSrc = make([]byte, 4)
	if err = binary.Read(buf, binary.BigEndian, &a.IPSrc); err != nil {
		return
	}
	n += 4
	a.HWDst = make([]byte, 6)
	if err = binary.Read(buf, binary.BigEndian, &a.HWDst); err != nil {
		return
	}
	n += 6
	a.IPDst = make([]byte, 4)
	if err = binary.Read(buf, binary.BigEndian, &a.IPDst); err != nil {
		return
	}
	n += 4
	return
}
