package avcommon

// #include "3rdparty/ffmpeg/rtsp.h"
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type RTSPState C.RTSPState

func CWrapRTSPState(ptr *types.CVoid) *C.RTSPState {
	return (*C.RTSPState)(unsafe.Pointer(ptr))
}

func WrapRTSPState(ptr *types.CVoid) *RTSPState {
	return (*RTSPState)(CWrapRTSPState(ptr))
}

func (s *RTSPState) RTSPStreams() []*RTSPStream {
	return unsafe.Slice((**RTSPStream)(unsafe.Pointer(s.rtsp_streams)), int(s.nb_rtsp_streams))
}

func (s *RTSPState) ControlURI() string {
	return C.GoString(&s.control_uri[0])
}

func (s *RTSPState) ControlURIBytes() [C.MAX_URL_SIZE]byte {
	return *(*[C.MAX_URL_SIZE]byte)(unsafe.Pointer(&s.control_uri[0]))
}

func (s *RTSPState) SetControlURI(controlURI string) {
	storage := s.ControlURIBytes()
	copy(storage[:], controlURI)
	if len(controlURI) >= len(storage) {
		storage[len(storage)-1] = 0
	} else {
		storage[len(controlURI)] = 0
	}
}
