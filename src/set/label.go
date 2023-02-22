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

func TextForLabel(obj *lib.LvObjT, text string) {
	C.lv_label_set_text((*C.lv_obj_t)(unsafe.Pointer(obj)), C.CString(text))

}
func TextStaticForLabel(obj *lib.LvObjT, text string) {
	C.lv_label_set_text_static((*C.lv_obj_t)(unsafe.Pointer(obj)), C.CString(text))

}
func LongModeForLabel(obj *lib.LvObjT, long_mode lib.LvLabelLongModeT) {
	C.lv_label_set_long_mode((*C.lv_obj_t)(unsafe.Pointer(obj)), C.lv_label_long_mode_t(long_mode))

}
func RecolorForLabel(obj *lib.LvObjT, en bool) {
	C.lv_label_set_recolor((*C.lv_obj_t)(unsafe.Pointer(obj)), C.bool(en))

}
func TextSelStartForLabel(obj *lib.LvObjT, index uint32) {
	C.lv_label_set_text_sel_start((*C.lv_obj_t)(unsafe.Pointer(obj)), C.uint(index))

}
func TextSelEndForLabel(obj *lib.LvObjT, index uint32) {
	C.lv_label_set_text_sel_end((*C.lv_obj_t)(unsafe.Pointer(obj)), C.uint(index))

}
