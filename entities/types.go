package entities

type NotificationHandler func([]byte) error

type EventSubscribe struct {
	DataIO uint8
	Association uint8
	Address uint8
	SendFailure uint8
	PeerError uint8
	Shutdown uint8
	PartialDelivery uint8
	AdaptationLayer uint8
	Authentication uint8
	SenderDry uint8
}

type InitMsg struct {
	NumOstreams    uint16
	MaxInstreams   uint16
	MaxAttempts    uint16
	MaxInitTimeout uint16
}

type SndRcvInfo struct {
	Stream uint16
	SSN uint16
	Flags uint16
	_ uint16
	PPID uint32
	Context uint32
	TTL uint32
	TSN uint32
	CumTSN uint32
	AssocID int32
}

type SndInfo struct {
	SID uint16
	Flags uint16
	PPID uint32
	Context uint32
	AssocID int32
}

type GetAddrsOld struct {
	AssocID int32
	AddrNum int32
	Addrs uintptr
}

type NotificationHeader struct {
	Type uint16
	Flags uint16
	Length uint32
}
