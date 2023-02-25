package get

import (
	"lvgl-go/src/lib"
	types "lvgl-go/src/types"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Label get

func CreateLable(o *types.LvObjT) Label {
	label := C.lv_label_create((*C.struct__lv_obj_t)(unsafe.Pointer(lib.GetParent(o))))

	return Label{
		CStructLvObjT: label,
	}
}

func (getter Label) Text() string {
	res := C.lv_label_get_text(getter.CStructLvObjT)

	return C.GoString(res)
}
func (getter Label) LongMode() types.LvLabelLongModeT {
	res := C.lv_label_get_long_mode(getter.CStructLvObjT)

	return types.LvLabelLongModeT(res)
}
func (getter Label) Recolor() bool {
	res := C.lv_label_get_recolor(getter.CStructLvObjT)

	return bool(res)
}
func (getter Label) LetterPos(char_id uint32, pos *types.LvPointT) Label {
	C.lv_label_get_letter_pos(getter.CStructLvObjT, C.uint(char_id), (*C.lv_point_t)(unsafe.Pointer(pos)))

	return getter
}
func (getter Label) LetterOn(pos_in *types.LvPointT) uint32 {
	res := C.lv_label_get_letter_on(getter.CStructLvObjT, (*C.lv_point_t)(unsafe.Pointer(pos_in)))

	return uint32(res)
}
func (getter Label) TextSelectionStart() uint32 {
	res := C.lv_label_get_text_selection_start(getter.CStructLvObjT)

	return uint32(res)
}
func (getter Label) TextSelectionEnd() uint32 {
	res := C.lv_label_get_text_selection_end(getter.CStructLvObjT)

	return uint32(res)
}
