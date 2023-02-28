package set

import (
	types "lvgl-go/src/types"
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

func (setter Canvas) Buffer(buf any, w types.LvCoordT, h types.LvCoordT, cf types.LvImgCfT) Canvas {
	C.lv_canvas_set_buffer(setter.CObj, unsafe.Pointer(&buf), C.lv_coord_t(w), C.lv_coord_t(h), C.lv_img_cf_t(cf))

	return setter
}
func (setter Canvas) PxColor(x types.LvCoordT, y types.LvCoordT, c types.LvColorT) Canvas {
	C.lv_canvas_set_px_color(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter Canvas) PxOpa(x types.LvCoordT, y types.LvCoordT, opa types.LvOpaT) Canvas {
	C.lv_canvas_set_px_opa(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter Canvas) Palette(id uint8, c types.LvColorT) Canvas {
	C.lv_canvas_set_palette(setter.CObj, C.uint8_t(id), C.lv_color_t(c))

	return setter
}
