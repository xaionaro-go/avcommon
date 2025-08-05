package avcommon

// #include "3rdparty/ffmpeg/libavformat/tcp.c"
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type TCPContext C.TCPContext

func CWrapTCPContext(ptr *types.CVoid) *C.TCPContext {
	return (*C.TCPContext)(unsafe.Pointer(ptr))
}

func WrapTCPContext(ptr *types.CVoid) *TCPContext {
	return (*TCPContext)(CWrapTCPContext(ptr))
}

func (tcpCtx *TCPContext) FileDescriptor() int {
	return int(tcpCtx.fd)
}
