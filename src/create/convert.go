package create

/*
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"
import (
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/get"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/set"
	"unsafe"
)

type tagUint uint8

const (
	screen tagUint = 0
	normal tagUint = 1
	STRUCT tagUint = 2
)

func _2CObj(o unsafe.Pointer) LV_OBJ_T {
	return (LV_OBJ_T)(o)
}

func _getParent[T _createI](o T, tag tagUint) LV_OBJ_T {
	r := o.GetObj()

	if r == nil {
		switch tag {
		case screen:
			return C.lv_obj_create(nil)
		default:
			return C.lv_scr_act()
		}
	}

	return _2CObj(r)
}

func getParent[T _createI](o T) LV_OBJ_T {
	return _getParent(o, normal)
}
func getParent2[T _createI](o T) LV_OBJ_T {
	return _getParent(o, screen)
}

func toSetObj(o LV_OBJ_T) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}
func toSetStyleT(o CStyleT) set.CStyleT {
	return (set.CStyleT)(unsafe.Pointer(o))
}
func toSetAnimT(o CAnimT) set.CAnimT {
	return (set.CAnimT)(unsafe.Pointer(o))
}

func toGetObj(o LV_OBJ_T) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}
func toGetStyleT(o CStyleT) get.CStyleT {
	return (get.CStyleT)(unsafe.Pointer(o))
}
func toGetAnimT(o CAnimT) get.CAnimT {
	return (get.CAnimT)(unsafe.Pointer(o))
}
