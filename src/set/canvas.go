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

type SetCanvas set

func (setter SetCanvas) SetBuffer(buf any, w lib.LvCoordT, h lib.LvCoordT, cf lib.LvImgCfT) SetCanvas {
	C.lv_canvas_set_buffer(setter.cObj, unsafe.Pointer(&buf), C.lv_coord_t(w), C.lv_coord_t(h), C.lv_img_cf_t(cf))

	return setter
}
func (setter SetCanvas) SetPxColor(x lib.LvCoordT, y lib.LvCoordT, c lib.LvColorT) SetCanvas {
	C.lv_canvas_set_px_color(setter.cObj, C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter SetCanvas) SetPxOpa(x lib.LvCoordT, y lib.LvCoordT, opa lib.LvOpaT) SetCanvas {
	C.lv_canvas_set_px_opa(setter.cObj, C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter SetCanvas) SetPalette(id uint8, c lib.LvColorT) SetCanvas {
	C.lv_canvas_set_palette(setter.cObj, C.uint8_t(id), C.lv_color_t(c))

	return setter
}
