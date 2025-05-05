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
