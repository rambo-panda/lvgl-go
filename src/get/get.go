package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type CObjT = *C.struct__lv_obj_t
type CStyleT = *C.lv_style_t
type CAnimT = *C.lv_anim_t

type get[T CAnimT | CObjT | CStyleT] struct {
	CObj T
}

func (g get[T]) Destroy() {
	g.CObj = nil
}
