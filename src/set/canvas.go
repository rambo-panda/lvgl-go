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

func BufferForCanvas(canvas *lib.LvObjT, buf any, w lib.LvCoordT, h lib.LvCoordT, cf lib.LvImgCfT) {
	C.lv_canvas_set_buffer((*C.struct__lv_obj_t)(unsafe.Pointer(canvas)), unsafe.Pointer(&buf), C.lv_coord_t(w), C.lv_coord_t(h), C.lv_img_cf_t(cf))

}
func PxColorForCanvas(canvas *lib.LvObjT, x lib.LvCoordT, y lib.LvCoordT, c lib.LvColorT) {
	C.lv_canvas_set_px_color((*C.struct__lv_obj_t)(unsafe.Pointer(canvas)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

}
func PxOpaForCanvas(canvas *lib.LvObjT, x lib.LvCoordT, y lib.LvCoordT, opa lib.LvOpaT) {
	C.lv_canvas_set_px_opa((*C.struct__lv_obj_t)(unsafe.Pointer(canvas)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

}
func PaletteForCanvas(canvas *lib.LvObjT, id uint8, c lib.LvColorT) {
	C.lv_canvas_set_palette((*C.struct__lv_obj_t)(unsafe.Pointer(canvas)), C.uint8_t(id), C.lv_color_t(c))

}
