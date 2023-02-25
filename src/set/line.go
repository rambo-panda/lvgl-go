package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type Line set

func CreateLine(o CObjT) Line {
	return Line{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Line) YInvert(en bool) Line {
	C.lv_line_set_y_invert(setter.CStructLvObjT, C.bool(en))

	return setter
}
