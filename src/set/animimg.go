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

func DurationForAnimimg(img *lib.LvObjT, duration uint32) {
	C.lv_animimg_set_duration((*C.struct__lv_obj_t)(unsafe.Pointer(img)), C.uint(duration))

}
func RepeatCountForAnimimg(img *lib.LvObjT, count uint16) {
	C.lv_animimg_set_repeat_count((*C.struct__lv_obj_t)(unsafe.Pointer(img)), C.ushort(count))

}
