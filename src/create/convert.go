package create

/*
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
	"unsafe"
)

type CObjT = *C.struct__lv_obj_t

const SCREEN uint8 = 0

func getParent(o CObjT, args ...uint8) CObjT {
	if o == nil {
		if len(args) == 0 {
			return C.lv_scr_act()
		}
		if args[0] == SCREEN {
			return C.lv_obj_create(nil)
		}
	}

	return o
}

func toSetObj(o CObjT) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}

func toGetObj(o CObjT) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}
