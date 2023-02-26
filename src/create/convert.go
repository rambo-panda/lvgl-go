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

type tagUint uint8

const (
	screen tagUint = 0
	normal tagUint = 1
)

func _getParent(o CObjT, tag tagUint) CObjT {
	if o == nil {
		switch tag {
		case screen:
			return C.lv_obj_create(nil)
		default:
			return C.lv_scr_act()
		}
	}

	return o
}

func getParent(o CObjT) CObjT {
	return _getParent(o, normal)
}
func getParent2(o CObjT) CObjT {
	return _getParent(o, screen)
}

func toSetObj(o CObjT) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}

func toGetObj(o CObjT) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}
