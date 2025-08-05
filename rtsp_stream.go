package avcommon

// #include "3rdparty/ffmpeg/libavformat/rtsp.h"
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type RTSPStream C.RTSPStream

func CWrapRTSPStream(ptr *types.CVoid) *C.RTSPStream {
	return (*C.RTSPStream)(unsafe.Pointer(ptr))
}

func WrapRTSPStream(ptr *types.CVoid) *RTSPStream {
	return (*RTSPStream)(CWrapRTSPStream(ptr))
}

func (s *RTSPStream) ControlURL() string {
	return C.GoString(&s.control_url[0])
}

func (s *RTSPStream) ControlURLBytes() [C.MAX_URL_SIZE]byte {
	return *(*[C.MAX_URL_SIZE]byte)(unsafe.Pointer(&s.control_url[0]))
}

func (s *RTSPStream) SetControlURL(controlURL string) {
	storage := s.ControlURLBytes()
	copy(storage[:], controlURL)
	if len(controlURL) >= len(storage) {
		storage[len(storage)-1] = 0
	} else {
		storage[len(controlURL)] = 0
	}
}
