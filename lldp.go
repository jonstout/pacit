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

type PortTLV struct {

}

type TTLTLV struct {

}
