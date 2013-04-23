package pacit

import (
	"io"
	"net"
	"bytes"
	"encoding/binary"
)

const (
	ICMP = 0x01
	IPv4 = 0x04
	TCP = 0x06
	UDP = 0x11
	IPv6 = 0x29
	IPv6ICMP = 0x3a
)

type IP struct {
	Version uint8 //4-bits
	IHL uint8 //4-bits
	DSCP uint8 //6-bits
	ECN uint8 //2-bits
	Length uint16
	ID uint16
	Flags uint16 //3-bits
	FragmentOffset uint16 //13-bits
	TTL uint8
	Protocol uint8
	Checksum uint16
	NWSrc net.IPAddr
	NWDst net.IPAddr
	Options []byte
}

func (i *IP) Len() (n uint16) {
	return uint16(i.IHL*32)
}

func (i *IP) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	var verIhl uint8 = (i.Version << 4) + i.IHL
	binary.Write(buf, binary.BigEndian, verIhl)
	var dscpEcn uint8 = (i.DSCP << 2) + i.ECN
	binary.Write(buf, binary.BigEndian, dscpEcn)
	binary.Write(buf, binary.BigEndian, i.Length)
	binary.Write(buf, binary.BigEndian, i.ID)
	var flagsFrag uint16 = (i.Flags << 13) + i.FragmentOffset
	binary.Write(buf, binary.BigEndian, flagsFrag)
	binary.Write(buf, binary.BigEndian, i.TTL)
	binary.Write(buf, binary.BigEndian, i.Protocol)
	binary.Write(buf, binary.BigEndian, i.Checksum)
	binary.Write(buf, binary.BigEndian, i.NWSrc)
	binary.Write(buf, binary.BigEndian, i.NWDst)
	binary.Write(buf, binary.BigEndian, i.Options)
	if n, err = buf.Read(b); n == 0 {
		return
	}
	return n, io.EOF	
}

func (i *IP) Write(b []byte) (n int, err error) {
	return
}