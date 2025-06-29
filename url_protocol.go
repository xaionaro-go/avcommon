package avcommon

// #include "3rdparty/ffmpeg/url.h"
import "C"
import (
	"unsafe"

	"github.com/xaionaro-go/avcommon/types"
)

type URLProtocol C.URLProtocol

func CWrapURLProtocol(ptr *types.CVoid) *C.URLProtocol {
	return (*C.URLProtocol)(unsafe.Pointer(ptr))
}

func WrapURLProtocol(ptr *types.CVoid) *URLProtocol {
	return (*URLProtocol)(CWrapURLProtocol(ptr))
}

func (urlProt *URLProtocol) Name() string {
	return C.GoString(urlProt.name)
}
