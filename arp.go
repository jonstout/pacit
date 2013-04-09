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

func (a *Arp) Read(b []byte) (n int, err error) {
	n, err := Ethernet.Read(b)
	if n == 0 {
		return
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, e.HWType)
	binary.Write(buf, binary.BigEndian, e.ProtoType)
	binary.Write(buf, binary.BigEndian, e.HWLength)
	binary.Write(buf, binary.BigEndian, e.ProtoLength)
	binary.Write(buf, binary.BigEndian, e.Operation)
	binary.Write(buf, binary.BigEndian, e.HWSrc)
	binary.Write(buf, binary.BigEndian, e.IPSrc)
	binary.Write(buf, binary.BigEndian, e.HWDst)
	binary.Write(buf, binary.BigEndian, e.IPDst)
	n, err = buf.Read(b)
	return
}

func (a *Arp) Write(b []byte) (n int, err error) {
	n, err := Ethernet.Write(b)
	if n == 0 {
		return
	}
	buf := bytes.NewBuffer(b[n:])
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
		if err = binary.Read(buf, binary.BigEndian, &a.HWSrc); err != nil {
		return
	}
	n += 6
		if err = binary.Read(buf, binary.BigEndian, &a.IPSrc); err != nil {
		return
	}
	n += 4
		if err = binary.Read(buf, binary.BigEndian, &a.HWDst); err != nil {
		return
	}
	n += 6
		if err = binary.Read(buf, binary.BigEndian, &a.IPDst); err != nil {
		return
	}
	n += 4
	return
}
