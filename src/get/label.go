package get

import (
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Label get[CObjT]

func CreateLable(o CObjT) *Label {
	_o := Label{
		CObj: o,
	}
	return &_o
}

func (getter *Label) Text() string {
	res := C.lv_label_get_text(getter.CObj)

	return C.GoString(res)
}
func (getter *Label) LongMode() uint8 {
	res := C.lv_label_get_long_mode(getter.CObj)

	return uint8(res)
}
func (getter *Label) Recolor() bool {
	res := C.lv_label_get_recolor(getter.CObj)

	return bool(res)
}

// TODO: 这里填充的暂不支持
// func (getter *Label) LetterPos(char_id uint32, pos_in *lib.LV_POINT_T) Label {
// 	C.lv_label_get_letter_pos(getter.CObj, C.uint(char_id), (*C.lv_point_t)(unsafe.Pointer(pos)))

//		return getter
//	}
func (getter *Label) LetterOn(pos_in *lib.LV_POINT_T) uint32 {
	res := C.lv_label_get_letter_on(getter.CObj, (*C.lv_point_t)(unsafe.Pointer(pos_in)))

	return uint32(res)
}
func (getter *Label) TextSelectionStart() uint32 {
	res := C.lv_label_get_text_selection_start(getter.CObj)

	return uint32(res)
}
func (getter *Label) TextSelectionEnd() uint32 {
	res := C.lv_label_get_text_selection_end(getter.CObj)

	return uint32(res)
}
