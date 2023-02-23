package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
	"unsafe"
)

type SetLed set

func CreateLed(o *lib.LvObjT) SetLed {
	return SetLed{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetLed) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetLed) SetColor(color lib.LvColorT) SetLed {
	C.lv_led_set_color(setter.CStructLvObjT, C.lv_color_t(color))

	return setter
}
func (setter SetLed) SetBrightness(bright uint8) SetLed {
	C.lv_led_set_brightness(setter.CStructLvObjT, C.uint8_t(bright))

	return setter
}
