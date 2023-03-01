package set

import (
	"lvgl-go/src/lib"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Led set[CObjT]

func CreateLed(o CObjT) Led {
	return Led{
		CObj: o,
	}
}

func (setter Led) Color(color lib.LV_COLOR_T) Led {
	C.lv_led_set_color(setter.CObj, C.lv_color_t(color))

	return setter
}
func (setter Led) Brightness(bright uint8) Led {
	C.lv_led_set_brightness(setter.CObj, C.uint8_t(bright))

	return setter
}
