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

type SetLed set

func CreateLed(o *types.LvObjT) SetLed {
	return SetLed{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetLed) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetLed) SetColor(color types.LvColorT) SetLed {
	C.lv_led_set_color(setter.CStructLvObjT, C.lv_color_t(color))

	return setter
}
func (setter SetLed) SetBrightness(bright uint8) SetLed {
	C.lv_led_set_brightness(setter.CStructLvObjT, C.uint8_t(bright))

	return setter
}
