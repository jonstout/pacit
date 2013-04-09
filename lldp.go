package pacit

import (
	"bytes"
	"encoding/binary"
)

type LLDP struct {
	Ethernet
	Chassis ChassisTLV
	Port PortTLV
	TTL TTLTLV
}

type ChassisTLV struct {

}

// Chassis ID subtypes
const (
	_ = iota
	CHASSIS_COMPONENT
	IFACE_ALIAS
	PORT_COMPONENT
	MAC_ADDR
	NET_ADDR
	IFACE_NAME
	LOCAL_ASSGN
)

// Port ID subtypes
const (
	_ = iota
	IFACE_ALIAS
	PORT_COMPONENT
	MAC_ADDR
	NET_ADDR	
	IFACE_NAME
	CIRCUIT_ID
	LOCAL_ASSGN
)

type PortTLV struct {
	Type uint8 //7bits
	Length uint16 //9bits
	Subtype uint8
	Data []uint8
}

func (t *PortTLV) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	var tni uint16 = 0
	typeAndLen := (tni | t.Type << 9) + (tni | t.Length)
	binary.Write(buf, binary.BigEndian, typeAndLen)
	binary.Write(buf, binary.BigEndian, t.Subtype)
	binary.Write(buf, binary.BigEndian, t.Data)
	n, err = buf.Read(b)
	return
}

func (t *PortTLV) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	var typeAndLen uint16 = 0
	if err = binary.Read(buf, binary.BigEndian, &typeAndLen); err != nil {
		return
	}
	n += 2
	t.Type = uint8(typeAndLen >> 9)
	t.Length = uint16( uint16(0x01ff) & typeAndLen)
	if err = binary.Read(buf, binary.BigEndian, &t.Subtype); err != nil {
		return
	}
	n += 1
	t.Data = make([]uint8, t.Length)
	if err = binary.Read(buf, binary.BigEndian, &t.Data); err != nil {
		return
	}
	n += t.Length
	return
}

type TTLTLV struct {
	Type uint8 //7 bits
	Length uint16 //9 bits
	Seconds uint16
}

func (t *TTLTLV) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	var tni uint16 = 0
	typeAndLen := (tni | t.Type << 9) + (tni | t.Length)
	binary.Write(buf, binary.BigEndian, typeAndLen)
	binary.Write(buf, binary.BigEndian, t.Seconds)
	n, err = buf.Read(b)
	return
}

func (t *TTLTLV) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	var typeAndLen uint16 = 0
	if err = binary.Read(buf, binary.BigEndian, &typeAndLen); err != nil {
		return
	}
	n += 2
	t.Type = uint8(typeAndLen >> 9)
	t.Length = uint16( uint16(0x01ff) & typeAndLen)
	if err = binary.Read(buf, binary.BigEndian, &t.Seconds); err != nil {
		return
	}
	n += 2
	return
}
