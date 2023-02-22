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

func ValueForBar(obj *lib.LvObjT, value int32, anim lib.LvAnimEnableT) {
	C.lv_bar_set_value((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.int(value), C.lv_anim_enable_t(anim))

}
func StartValueForBar(obj *lib.LvObjT, start_value int32, anim lib.LvAnimEnableT) {
	C.lv_bar_set_start_value((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.int(start_value), C.lv_anim_enable_t(anim))

}
func RangeForBar(obj *lib.LvObjT, min int32, max int32) {
	C.lv_bar_set_range((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.int(min), C.int(max))

}
func ModeForBar(obj *lib.LvObjT, mode lib.LvBarModeT) {
	C.lv_bar_set_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_bar_mode_t(mode))

}
