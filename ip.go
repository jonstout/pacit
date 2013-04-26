package pacit

import (
	"io"
	"net"
	"bytes"
	"encoding/binary"
)

const (
	IP_ICMP = 0x01
	IP_TCP = 0x06
	IP_UDP = 0x11
	IP_IPv6 = 0x29
	IP_IPv6ICMP = 0x3a
)

type IPv4 struct {
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
	NWSrc net.IP
	NWDst net.IP
	Options []byte
}

func (i *IPv4) Len() (n uint16) {
	return uint16(i.IHL*32)
}

func (i *IPv4) Read(b []byte) (n int, err error) {
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

func (i *IPv4) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	var verIhl uint8 = 0
	if err = binary.Read(buf, binary.BigEndian, verIhl); err != nil {
		return
	}
	n += 1
	i.Version = verIhl >> 4
	i.IHL = verIhl & 0x0f
	var dscpEcn uint8 = 0
	if err = binary.Read(buf, binary.BigEndian, dscpEcn); err != nil {
		return
	}
	n += 1
	i.DSCP = dscpEcn >> 2
	i.ECN = dscpEcn & 0x02
	if err = binary.Read(buf, binary.BigEndian, &i.Length); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &i.ID); err != nil {
		return
	}
	n += 2
	var flagsFrag uint16 = 0
	if err = binary.Read(buf, binary.BigEndian, flagsFrag); err != nil {
		return
	}
	n += 2
	i.Flags = flagsFrag >> 13
	i.FragmentOffset = flagsFrag & 0x1fff
	if err = binary.Read(buf, binary.BigEndian, &i.TTL); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &i.Protocol); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &i.Checksum); err != nil {
		return
	}
	n += 2
	i.NWSrc = make([]byte, 4)
	if err = binary.Read(buf, binary.BigEndian, &i.NWSrc); err != nil {
		return
	}
	n += 4
	i.NWSrc = make([]byte, 4)
	if err = binary.Read(buf, binary.BigEndian, &i.NWDst); err != nil {
		return
	}
	n += 4
	i.Options = make([]byte, 4*(int(i.IHL) - 5))
	if err = binary.Read(buf, binary.BigEndian, &i.Options); err != nil {
		return
	}
	n += 4*(int(i.IHL) - 5)
	return
}