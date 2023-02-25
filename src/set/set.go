package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/lib"
	types "lvgl-go/src/types"
	"unsafe"
)

type _cs *C.struct__lv_obj_t
type set struct {
	CStructLvObjT _cs
}

func Go2CObj(o *types.LvObjT, parent bool) *C.struct__lv_obj_t {
	_o := o
	if parent {
		_o = lib.GetParent(o)
	}

	return (*C.struct__lv_obj_t)(unsafe.Pointer(_o))
}

func getObj(l any) _cs {
	return (_cs)(l.(unsafe.Pointer))
}
