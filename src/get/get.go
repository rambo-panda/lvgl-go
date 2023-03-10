package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
	"unsafe"
)

type CObjT = *C.struct__lv_obj_t
type CStyleT = *C.lv_style_t
type CAnimT = *C.lv_anim_t

type get[T CAnimT | CObjT | CStyleT] struct {
	CObj T
}

func (g get[T]) Destroy(tag uint8) {
	lib.Destroy(unsafe.Pointer(g.CObj), tag)
	g.CObj = nil
}
