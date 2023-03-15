package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"
import (
	"unsafe"

	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
)

type CObjT = *C.struct__lv_obj_t
type CStyleT = *C.lv_style_t
type CAnimT = *C.lv_anim_t

type set[T CObjT | CStyleT | CAnimT] struct {
	CObj T
}

func getParentObj(o lib.CreateI) unsafe.Pointer {
	return o.GetObj()
}
