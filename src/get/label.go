package get

import types "lvgl-go/src/types"

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type GetLabel Get

func CreateLable(o *types.LvObjT) GetLabel {
	return GetLabel{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}
func (setter GetLabel) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (getter GetLabel) GetText() string {
	res := C.lv_label_get_text(getter.CStructLvObjT)

	return C.GoString(res)
}
func (getter GetLabel) GetLongMode() types.LvLabelLongModeT {
	res := C.lv_label_get_long_mode(getter.CStructLvObjT)

	return types.LvLabelLongModeT(res)
}
func (getter GetLabel) GetRecolor() bool {
	res := C.lv_label_get_recolor(getter.CStructLvObjT)

	return bool(res)
}
func (getter GetLabel) GetLetterPos(char_id uint32, pos *types.LvPointT) GetLabel {
	C.lv_label_get_letter_pos(getter.CStructLvObjT, C.uint(char_id), (*C.lv_point_t)(unsafe.Pointer(pos)))

	return getter
}
func (getter GetLabel) GetLetterOn(pos_in *types.LvPointT) uint32 {
	res := C.lv_label_get_letter_on(getter.CStructLvObjT, (*C.lv_point_t)(unsafe.Pointer(pos_in)))

	return uint32(res)
}
func (getter GetLabel) GetTextSelectionStart() uint32 {
	res := C.lv_label_get_text_selection_start(getter.CStructLvObjT)

	return uint32(res)
}
func (getter GetLabel) GetTextSelectionEnd() uint32 {
	res := C.lv_label_get_text_selection_end(getter.CStructLvObjT)

	return uint32(res)
}
