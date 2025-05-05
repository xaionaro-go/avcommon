package avcommon

// #include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type AVCodecContext C.AVCodecContext

func CWrapAVCodecContext(ptr *types.CVoid) *C.AVCodecContext {
	return (*C.AVCodecContext)(unsafe.Pointer(ptr))
}

func WrapAVCodecContext(ptr *types.CVoid) *AVCodecContext {
	return (*AVCodecContext)(CWrapAVCodecContext(ptr))
}

func (codecCtx *AVCodecContext) PrivData() *types.CVoid {
	return (*types.CVoid)(codecCtx.priv_data)
}
