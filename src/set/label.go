package set

import (
	"lvgl-go/src/types"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Label set

func CreateLabel(o *types.LvObjT) Label {
	return Label{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}
func (setter Label) Text(text string) Label {
	j := C.CString(text)
	C.lv_label_set_text(setter.CStructLvObjT, j)

	return setter
}
func (setter Label) TextStatic(text string) Label {
	C.lv_label_set_text_static(setter.CStructLvObjT, C.CString(text))

	return setter
}
func (setter Label) LongMode(long_mode types.LvLabelLongModeT) Label {
	C.lv_label_set_long_mode(setter.CStructLvObjT, C.lv_label_long_mode_t(long_mode))

	return setter
}
func (setter Label) Recolor(en bool) Label {
	C.lv_label_set_recolor(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter Label) TextSelStart(index uint32) Label {
	C.lv_label_set_text_sel_start(setter.CStructLvObjT, C.uint(index))

	return setter
}
func (setter Label) TextSelEnd(index uint32) Label {
	C.lv_label_set_text_sel_end(setter.CStructLvObjT, C.uint(index))

	return setter
}
