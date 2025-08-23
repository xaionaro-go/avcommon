package avcommon

// #include <libavformat/avio.h>
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type AVIOContext C.AVIOContext

func CWrapAVIOContext(ptr *types.CVoid) *C.AVIOContext {
	return (*C.AVIOContext)(unsafe.Pointer(ptr))
}

func WrapAVIOContext(ptr *types.CVoid) *AVIOContext {
	return (*AVIOContext)(CWrapAVIOContext(ptr))
}

func (avioCtx *AVIOContext) Opaque() *types.CVoid {
	return (*types.CVoid)(avioCtx.opaque)
}

func (avioCtx *AVIOContext) Buffer() []byte {
	length := uintptr(unsafe.Pointer(avioCtx.buf_ptr)) - uintptr(unsafe.Pointer(avioCtx.buffer))
	s := unsafe.Slice((*byte)(unsafe.Pointer(avioCtx.buffer)), int(avioCtx.buffer_size))
	return s[:length]
}
