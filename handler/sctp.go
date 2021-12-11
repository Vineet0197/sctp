package handler

import (
	"encoding/binary"
	"sctp/entities"
	"unsafe"
)

type SCTPHandler struct {
	nativeEndian binary.ByteOrder
	sndRcvInfoSize uintptr
}

func (s *SCTPHandler) init()  {
	init := uint16(1)
	if *(*byte)(unsafe.Pointer(&init)) == 0 {
		s.nativeEndian = binary.BigEndian
	} else {
		s.nativeEndian = binary.LittleEndian
	}
	info := entities.SndRcvInfo{}
	s.sndRcvInfoSize = unsafe.Sizeof(info)

}
