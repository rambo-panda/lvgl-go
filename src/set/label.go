package set

import "lvgl-go/src/types"

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type SetLabel set

// func CreateLabel(o CI) SetLabel {
// 	return SetLabel{
// 		CStructLvObjT:
// 	}
// }

func CreateLabel(o *types.LvObjT) SetLabel {
	return SetLabel{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetLabel) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetLabel) SetText(text string) SetLabel {
	j := C.CString(text)
	C.lv_label_set_text(setter.CStructLvObjT, j)

	return setter
}
func (setter SetLabel) SetTextStatic(text string) SetLabel {
	C.lv_label_set_text_static(setter.CStructLvObjT, C.CString(text))

	return setter
}
func (setter SetLabel) SetLongMode(long_mode types.LvLabelLongModeT) SetLabel {
	C.lv_label_set_long_mode(setter.CStructLvObjT, C.lv_label_long_mode_t(long_mode))

	return setter
}
func (setter SetLabel) SetRecolor(en bool) SetLabel {
	C.lv_label_set_recolor(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter SetLabel) SetTextSelStart(index uint32) SetLabel {
	C.lv_label_set_text_sel_start(setter.CStructLvObjT, C.uint(index))

	return setter
}
func (setter SetLabel) SetTextSelEnd(index uint32) SetLabel {
	C.lv_label_set_text_sel_end(setter.CStructLvObjT, C.uint(index))

	return setter
}
