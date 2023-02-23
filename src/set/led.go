package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
)

type SetLed set

func (setter SetLed) SetColor(color lib.LvColorT) SetLed {
	C.lv_led_set_color(setter.cObj, C.lv_color_t(color))

	return setter
}
func (setter SetLed) SetBrightness(bright uint8) SetLed {
	C.lv_led_set_brightness(setter.cObj, C.uint8_t(bright))

	return setter
}
