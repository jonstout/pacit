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
	Info uint16 //9bits
	TTL uint16
}

func (t *PortTLV) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	var tni uint16 = 0
	typeAndInfo := (tni | t.Type << 9) + (tni | t.Info)
	binary.Write(buf, binary.BigEndian, typeAndInfo)
	binary.Write(buf, binary.BigEndian, t.TTL)
	n, err = buf.Read(b)
	return
}

func (t *PortTLV) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	var typeAndInfo uint16 = 0
	if err = binary.Read(buf, binary.BigEndian, &typeAndInfo); err != nil {
		return
	}
	n += 2
	t.Type = uint8(typeAndInfo >> 9)
	t.Info = uint16( uint16(0) | typeAndInfo)
	if err = binary.Read(buf, binary.BigEndian, &t.TTL); err != nil {
		return
	}
	n += 2
	return
}

type TTLTLV struct {
	Type uint8 //7 bits
	Info uint16 //9 bits
	TTL uint16
}

func (t *TTLTLV) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	var tni uint16 = 0
	typeAndInfo := (tni | t.Type << 9) + (tni | t.Info)
	binary.Write(buf, binary.BigEndian, typeAndInfo)
	binary.Write(buf, binary.BigEndian, t.TTL)
	n, err = buf.Read(b)
	return
}

func (t *TTLTLV) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	var typeAndInfo uint16 = 0
	if err = binary.Read(buf, binary.BigEndian, &typeAndInfo); err != nil {
		return
	}
	n += 2
	t.Type = uint8(typeAndInfo >> 9)
	t.Info = uint16( uint16(0) | typeAndInfo)
	if err = binary.Read(buf, binary.BigEndian, &t.TTL); err != nil {
		return
	}
	n += 2
	return
}
