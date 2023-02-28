package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type CObjT = *C.struct__lv_obj_t
type CStyleT = *C.lv_style_t

type set[T CObjT | CStyleT] struct {
	CObj T
}
