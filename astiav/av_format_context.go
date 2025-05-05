package astiav

import (
	"reflect"

	"github.com/asticode/go-astiav"
	"github.com/xaionaro-go/avcommon/types"
	"github.com/xaionaro-go/unsafetools"
)

func CFromAVFormatContext(fmtCtx *astiav.FormatContext) *types.CVoid {
	orig := reflect.ValueOf(fmtCtx)
	unsafePtr := unsafetools.FieldByNameInValue(orig, "c").Elem().UnsafePointer()
	return (*types.CVoid)(unsafePtr)
}
