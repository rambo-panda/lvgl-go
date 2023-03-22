package create

/*
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"
import (
	"unsafe"

	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/get"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/set"
)

type tagUint uint8

const (
	screen    tagUint = 0
	normal    tagUint = 1
	container tagUint = 2
)

func _2CObj(o unsafe.Pointer) _lvObjT {
	return (_lvObjT)(o)
}

func _getParent(o _createI, tag tagUint) _lvObjT {
	r := o.GetObj()

	if r == nil {
		switch tag {
		case screen:
			return C.lv_obj_create(nil)
		case container:
			return C.lv_obj_create(C.lv_scr_act())
		default:
			return C.lv_scr_act()
		}
	}

	return _2CObj(r)
}

func getParent(o _createI, tag tagUint) _lvObjT {
	return _getParent(o, tag)
}

func toSetObj(o _lvObjT) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}
func toSetStyleT(o _PcStyleT) set.CStyleT {
	return (set.CStyleT)(unsafe.Pointer(o))
}
func toSetAnimT(o _pAnimT) set.CAnimT {
	return (set.CAnimT)(unsafe.Pointer(o))
}

func toGetObj(o _lvObjT) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}
func toGetStyleT(o _PcStyleT) get.CStyleT {
	return (get.CStyleT)(unsafe.Pointer(o))
}
func toGetAnimT(o _pAnimT) get.CAnimT {
	return (get.CAnimT)(unsafe.Pointer(o))
}
