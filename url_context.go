package avcommon

// #include "3rdparty/ffmpeg/url.h"
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type URLContext C.URLContext

func CWrapURLContext(ptr *types.CVoid) *C.URLContext {
	return (*C.URLContext)(unsafe.Pointer(ptr))
}

func WrapURLContext(ptr *types.CVoid) *URLContext {
	return (*URLContext)(CWrapURLContext(ptr))
}

func (urlCtx *URLContext) PrivData() *types.CVoid {
	return (*types.CVoid)(urlCtx.priv_data)
}
