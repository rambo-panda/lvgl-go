package create

/*
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
	"reflect"
	"unsafe"
)

type CObjT = *C.struct__lv_obj_t

type tagUint uint8

const (
	screen tagUint = 0
	normal tagUint = 1
	STRUCT tagUint = 2
)

func _getParent(o any, tag tagUint) CObjT {
	if o == nil {
		switch tag {
		case screen:
			return C.lv_obj_create(nil)
		default:
			return C.lv_scr_act()
		}
	}

	defer func() {
		if err := recover(); err != nil {
			panic("getParent3 no return")
		}
	}()

	return (CObjT)(reflect.ValueOf(o).MethodByName("getObj").Call(nil)[0].UnsafePointer())
}

func getParent(o any) CObjT {
	return _getParent(o, normal)
}
func getParent2(o any) CObjT {
	return _getParent(o, screen)
}

func toSetObj(o CObjT) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}

func toGetObj(o CObjT) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}
