package set

import (
	lib.LV_FONT_T"lvgl-go/src/lib.LV_FONT_T
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Canvas set[CObjT]

func CreateCanvas(o CObjT) Canvas {
	return Canvas{
		CObj: o,
	}
}

func (setter Canvas) Buffer(buf any, w lib.LV_COLOR_T, h lib.LV_COLOR_T, cf lib.LV_FONT_TLvImgCfT) Canvas {
	C.lv_canvas_set_buffer(setter.CObj, unsafe.Pointer(&buf), C.lv_coord_t(w), C.lv_coord_t(h), C.lv_img_cf_t(cf))

	return setter
}
func (setter Canvas) PxColor(x lib.LV_COLOR_T, y lib.LV_COLOR_T, c lib.LV_COLOR_T) Canvas {
	C.lv_canvas_set_px_color(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter Canvas) PxOpa(x lib.LV_COLOR_T, y lib.LV_COLOR_T, opa lib.LV_OPA_T) Canvas {
	C.lv_canvas_set_px_opa(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter Canvas) Palette(id uint8, c lib.LV_COLOR_T) Canvas {
	C.lv_canvas_set_palette(setter.CObj, C.uint8_t(id), C.lv_color_t(c))

	return setter
}
