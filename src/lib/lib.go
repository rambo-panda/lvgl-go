package lib

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L./ -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
	"lvgl-go/src/types"
	"unsafe"
)

func Ready() {
	C.lv_ready()
}

func TaskHandler(ms uint) {
	C.lv_task_handler2(C.uint(ms))
}

func TaskHandlerAsync(ms uint) {
	go TaskHandler(ms)
}

func C2GoObj(o any) *types.LvObjT {
	return (*types.LvObjT)(o.(unsafe.Pointer))
}

func Go2CObj(o *types.LvObjT, t string) any {
	_o := unsafe.Pointer(o)

	switch t {
	case "get":
		return (*get.LvObjT)(_o)
	case "set":
		return (*set.LvObjT)(_o)
	default:
		return (*C.struct__lv_obj_t)(_o)
	}
}

func GetParent(o any) *types.LvObjT {
	if o == nil {
		o = C.lv_scr_act()
	}

	return C2GoObj(o)
}
