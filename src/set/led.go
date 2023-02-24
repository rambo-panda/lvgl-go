package set

import types "lvgl-go/src/types"

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type Led set

func CreateLed(o *types.LvObjT) Led {
	return Led{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Led) Color(color types.LvColorT) Led {
	C.lv_led_set_color(setter.CStructLvObjT, C.lv_color_t(color))

	return setter
}
func (setter Led) Brightness(bright uint8) Led {
	C.lv_led_set_brightness(setter.CStructLvObjT, C.uint8_t(bright))

	return setter
}
