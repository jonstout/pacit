package pacit

import (
	//	"bytes"
	//	"encoding/binary"
	"errors"
	//"io"
	"math/rand"
	"net"
)

const (
	MSG_BOOT_REQ byte = iota
	MSG_BOOT_RES
)

const (
	_ byte = iota
	MSG_DISCOVER
	MSG_OFFER
	MSG_REQUST
	MSG_DECLINE
	MSG_ACK
	MSG_NAK
	MSG_RELEASE
	MSG_INFORM
)

type DHCPOption interface {
	OptionType() byte
	Bytes() []byte
}

type DHCP struct {
	Operation    byte
	HardwareType byte
	HardwareLen  uint8
	Hops         uint8
	Xid          uint32
	Secs         uint16
	Flags        uint16
	ClientIP     net.IP
	YourIP       net.IP
	ServerIP     net.IP
	GatewayIP    net.IP
	ClientHWAddr net.HardwareAddr
	Options      []DHCPOption
}

const (
	OPT_PAD byte = iota
	OPT_SUBNET_MASK
	OPT_TIME_OFFSET
	OPT_ROUTERS
	OPT_TIME_SERVERS
	OPT_NAME_SERVERS
	OPT_DNS_SERVERS
	OPT_LOG_SERVERS
	OPT_COOKIE_SERVERS
	OPT_LPR_SERVERS
)

const (
	HW_ETHERNET byte = 0x01
)

const (
	FLAG_BROADCAST uint16 = 0x80

//	FLAG_BROADCAST_MASK uint16 = (1 << FLAG_BROADCAST)
)

func NewDHCP(xid uint32, op, hwtype byte, hwaddr net.HardwareAddr, C, Y, S, G net.IP) (*DHCP, error) {
	if xid == 0 {
		xid = rand.Uint32()
	}
	switch hwtype {
	case HW_ETHERNET:
		break
	default:
		return nil, errors.New("Bad HardwareType")
	}
	d := &DHCP{
		Operation:    op,
		HardwareType: hwtype,
		HardwareLen:  byte(len(hwaddr)),
		Xid:          xid,
		Flags:        FLAG_BROADCAST,
		ClientIP:     C,
		YourIP:       Y,
		ServerIP:     S,
		GatewayIP:    G,
	}
	return d, nil
}

//func (d *DHCP) Len() (n uint16) {
//return uint16(4 + len(i.Data))
//}
/*
func (i *ICMP) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, i.Type)
	binary.Write(buf, binary.BigEndian, i.Code)
	binary.Write(buf, binary.BigEndian, i.Checksum)
	binary.Write(buf, binary.BigEndian, i.Data)
	if n, err = buf.Read(b); n == 0 {
		return
	}
	return n, io.EOF
}

/*x
func (i *ICMP) ReadFrom(r io.Reader) (n int64, err error) {
	if err = binary.Read(r, binary.BigEndian, &i.Type); err != nil {
		return
	}
	n += 1
	if err = binary.Read(r, binary.BigEndian, &i.Code); err != nil {
		return
	}
	n += 1
	if err = binary.Read(r, binary.BigEndian, &i.Checksum); err != nil {
		return
	}
	n += 2
	return
}

func (i *ICMP) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
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
	i.Data = make([]byte, len(b)-n)
	if err = binary.Read(buf, binary.BigEndian, &i.Data); err != nil {
		return
	}
	n += len(i.Data)
	return
}
*/
