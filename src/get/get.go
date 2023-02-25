package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	types "lvgl-go/src/types"
	"unsafe"
)

type get struct {
	CStructLvObjT *C.struct__lv_obj_t
}

func Go2CObj(o *types.LvObjT) *C.struct__lv_obj_t {
	return (*C.struct__lv_obj_t)(unsafe.Pointer(o))
}
