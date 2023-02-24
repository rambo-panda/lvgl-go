package set

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

type SetCheckbox set

func CreateCheckbox(o *types.LvObjT) SetCheckbox {
	return SetCheckbox{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetCheckbox) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetCheckbox) SetText(txt string) SetCheckbox {
	C.lv_checkbox_set_text(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter SetCheckbox) SetTextStatic(txt string) SetCheckbox {
	C.lv_checkbox_set_text_static(setter.CStructLvObjT, C.CString(txt))

	return setter
}
