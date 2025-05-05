package avcommon

// #include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type AVFormatContext C.AVFormatContext

func CWrapAVFormatContext(ptr *types.CVoid) *C.AVFormatContext {
	return (*C.AVFormatContext)(unsafe.Pointer(ptr))
}

func WrapAVFormatContext(ptr *types.CVoid) *AVFormatContext {
	return (*AVFormatContext)(CWrapAVFormatContext(ptr))
}

func (fmtCtx *AVFormatContext) PrivData() *types.CVoid {
	return (*types.CVoid)(fmtCtx.priv_data)
}

func (fmtCtx *AVFormatContext) Pb() *AVIOContext {
	return (*AVIOContext)(fmtCtx.pb)
}
