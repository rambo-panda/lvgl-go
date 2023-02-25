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

type _ct *C.struct__lv_obj_t
type set struct {
	CStructLvObjT _ct
}

func Go2CObj(o *types.LvObjT, parent bool) _ct {
	_o := o
	if parent {
		_o = lib.GetParent(o)
	}

	return (_ct)(unsafe.Pointer(_o))
}

func getObj(l any) _ct {
	return (_ct)(l.(unsafe.Pointer))
}
