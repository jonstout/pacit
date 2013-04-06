package pacit

type ARP struct {
	Ethernet
	HWType uint16
	ProtoType uint16
	HWLength uint8
	ProtoLength uint8
	Operation uint16
	HWSrc [6]uint8
	IPSrc uint32
	HWDst [6]uint8
	IPDst uint32
}

