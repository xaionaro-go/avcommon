package types

import "C"
import "unsafe"

type CVoid C.void

func (ptr *CVoid) UnsafePointer() unsafe.Pointer {
	return unsafe.Pointer(ptr)
}

func AsCVoid[T any](in *T) *CVoid {
	return (*CVoid)((unsafe.Pointer)(in))
}
