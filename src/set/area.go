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

func WidthForArea(area_p *lib.LvAreaT, w lib.LvCoordT) {
	C.lv_area_set_width((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(w))

}
func HeightForArea(area_p *lib.LvAreaT, h lib.LvCoordT) {
	C.lv_area_set_height((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(h))

}
