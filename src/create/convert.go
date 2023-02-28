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

type tagUint uint8

const (
	screen tagUint = 0
	normal tagUint = 1
	STRUCT tagUint = 2
)

func _getParent[T _createI](o T, tag tagUint) CObjT {
	r := o.GetObj()

	if r == nil {
		switch tag {
		case screen:
			return C.lv_obj_create(nil)
		default:
			return C.lv_scr_act()
		}
	}

	return r
}

func getParent[T _createI](o T) CObjT {
	return _getParent(o, normal)
}
func getParent2[T _createI](o T) CObjT {
	return _getParent(o, screen)
}

func toSetObj(o CObjT) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}
func toSetStyleT(o CStyleT) set.CStyleT {
	return (set.CStyleT)(unsafe.Pointer(o))
}
func toSetAnim(o CAnimT) set.CAnimT {
	return (set.CAnimT)(unsafe.Pointer(o))
}

func toGetObj(o CObjT) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}
func toGetStyleT(o CStyleT) get.CStyleT {
	return (get.CStyleT)(unsafe.Pointer(o))
}
func toGetStyleT(o CStyleT) get.CAnimT {
	return (get.CAnimT)(unsafe.Pointer(o))
}
