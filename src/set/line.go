package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
	"unsafe"
)

type SetLine set

func CreateLine(o *lib.LvObjT) SetLine {
	return SetLine{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetLine) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetLine) SetYInvert(en bool) SetLine {
	C.lv_line_set_y_invert(setter.CStructLvObjT, C.bool(en))

	return setter
}
