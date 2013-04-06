package pacit

import (
	"bytes"
	"encoding/binary"
)

type Ethernet struct {
	Preamble [7]uint8
	Delimiter uint8
	HWDst [6]uint8
	HWSrc [6]uint8
	VLANID VLAN
	Ethertype uint16
}

func (e *Ethernet) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, e.Preamble)
	binary.Write(buf, binary.BigEndian, e.Delimiter)
	binary.Write(buf, binary.BigEndian, e.HWDst)
	binary.Write(buf, binary.BigEndian, e.HWSrc)
	if e.VLANID.VID != 0 {
		b := []byte{0, 0}
		e.VLANID.Read(b)
		binary.Write(buf, binary.BigEndian, b)
	}
	binary.Write(buf, binary.BigEndian, e.Ethertype)
	n, err = buf.Read(b)
	return
}

func (e *Ethernet) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &e.Preamble); err != nil {
		return
	}
	n += 7
	if err = binary.Read(buf, binary.BigEndian, &e.Delimiter); err != nil {
		return
	}
	n += 1
	if err = binary.Read(buf, binary.BigEndian, &e.HWDst); err != nil {
		return
	}
	n += 6
	if err = binary.Read(buf, binary.BigEndian, &e.HWSrc); err != nil {
		return
	}
	n += 6
	if err = binary.Read(buf, binary.BigEndian, &e.Ethertype); err != nil {
		return
	}
	n += 2
	// If tagged
	if e.Ethertype == 0x8100 {
		b := make([]byte, 2)
		b[0] = byte(e.Ethertype >> 8); b[1] = byte(e.Ethertype)
		c := buf.Next(2)
		e.VLANID.Write( append(b, c[0], c[1]) )
		n += 2
		if err = binary.Read(buf, binary.BigEndian, &e.Ethertype); err != nil {
			return
		}
		n += 2
		return
	} else {
		e.VLANID.VID = 0
	}
	return
}

const (
	PCP_MASK = 0xe000
	DEI_MASK = 0x1000
	VID_MASK = 0x0fff
)

type VLAN struct {
	TPID uint16
	PCP uint8
	DEI uint8
	VID uint8
}

func (v *VLAN) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, v.TPID)
	var tci uint16 = 0
	tci = (tci | uint16(v.PCP) << 13) + (tci | uint16(v.DEI) << 12) + (tci | uint16(v.VID))
	binary.Write(buf, binary.BigEndian, tci)
	n, err = buf.Read(b)
	return
}

func (v *VLAN) Write(b []byte) (n int, err error) {
	var tci uint16 = 0
	buf := bytes.NewBuffer(b)
	if err = binary.Read(buf, binary.BigEndian, &v.TPID); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &tci); err != nil {
		return
	}
	n += 2
	v.PCP = uint8(PCP_MASK & tci >> 13)
	v.DEI = uint8(DEI_MASK & tci >> 12)
	v.VID = uint8(VID_MASK & tci)
	return
}