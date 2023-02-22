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

func TextForTextarea(obj *lib.LvObjT, txt string) {
	C.lv_textarea_set_text((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(txt))

}
func PlaceholderTextForTextarea(obj *lib.LvObjT, txt string) {
	C.lv_textarea_set_placeholder_text((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(txt))

}
func CursorPosForTextarea(obj *lib.LvObjT, pos int32) {
	C.lv_textarea_set_cursor_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.int(pos))

}
func CursorClickPosForTextarea(obj *lib.LvObjT, en bool) {
	C.lv_textarea_set_cursor_click_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(en))

}
func PasswordModeForTextarea(obj *lib.LvObjT, en bool) {
	C.lv_textarea_set_password_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(en))

}
func PasswordBulletForTextarea(obj *lib.LvObjT, bullet string) {
	C.lv_textarea_set_password_bullet((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(bullet))

}
func OneLineForTextarea(obj *lib.LvObjT, en bool) {
	C.lv_textarea_set_one_line((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(en))

}
func AcceptedCharsForTextarea(obj *lib.LvObjT, list string) {
	C.lv_textarea_set_accepted_chars((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(list))

}
func MaxLengthForTextarea(obj *lib.LvObjT, num uint32) {
	C.lv_textarea_set_max_length((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(num))

}
func InsertReplaceForTextarea(obj *lib.LvObjT, txt string) {
	C.lv_textarea_set_insert_replace((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(txt))

}
func TextSelectionForTextarea(obj *lib.LvObjT, en bool) {
	C.lv_textarea_set_text_selection((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(en))

}
func PasswordShowTimeForTextarea(obj *lib.LvObjT, time uint16) {
	C.lv_textarea_set_password_show_time((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(time))

}
func AlignForTextarea(obj *lib.LvObjT, align lib.LvTextAlignT) {
	C.lv_textarea_set_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_align_t(align))

}
