package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/lib"
	"unsafe"
)

type CObjT = *C.struct__lv_obj_t
type CStyleT = *C.lv_style_t
type CAnimT = *C.lv_anim_t

type set[T CObjT | CStyleT | CAnimT] struct {
	CObj T
}

func (s set[T]) Destroy() {
	s.CObj = nil
}

func getParentObj(o lib.CreateI) unsafe.Pointer {
	return o.GetObj()
}
