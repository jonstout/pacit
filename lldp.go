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

}

type TTLTLV struct {

}
