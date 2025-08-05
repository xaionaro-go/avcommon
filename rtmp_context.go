package avcommon

// #include "3rdparty/ffmpeg/libavformat/rtmpproto.c"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type RTMPContext C.RTMPContext

func CWrapRTMPContext(ptr *types.CVoid) *C.RTMPContext {
	return (*C.RTMPContext)(unsafe.Pointer(ptr))
}

func WrapRTMPContext(ptr *types.CVoid) *RTMPContext {
	return (*RTMPContext)(CWrapRTMPContext(ptr))
}

func (rtmpCtx *RTMPContext) App() string {
	return (C.GoString)(rtmpCtx.app)
}

func (rtmpCtx *RTMPContext) AppBytes() [C.APP_MAX_LENGTH]byte {
	return *(*[C.APP_MAX_LENGTH]byte)(unsafe.Pointer(rtmpCtx.app))
}

func (rtmpCtx *RTMPContext) SetApp(app string) {
	storage := rtmpCtx.AppBytes()
	copy(storage[:], app)
	if len(app) >= len(storage) {
		storage[len(storage)-1] = 0
	} else {
		storage[len(app)] = 0
	}
}

func (rtmpCtx *RTMPContext) Stream() *URLContext {
	return (*URLContext)(rtmpCtx.stream)
}

func (rtmpCtx *RTMPContext) TCPContext() *TCPContext {
	stream := rtmpCtx.Stream()
	protocolName := stream.Protocol().Name()
	if protocolName != "tcp" {
		panic(fmt.Sprintf("unexpected protocol name: '%s'", protocolName))
	}
	return WrapTCPContext(stream.PrivData())
}

func (rtmpCtx *RTMPContext) GetFileDescriptor() int {
	return rtmpCtx.TCPContext().FileDescriptor()
}
