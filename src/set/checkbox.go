package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type Checkbox set

func CreateCheckbox(o CObjT) Checkbox {
	return Checkbox{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Checkbox) Text(txt string) Checkbox {
	C.lv_checkbox_set_text(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter Checkbox) TextStatic(txt string) Checkbox {
	C.lv_checkbox_set_text_static(setter.CStructLvObjT, C.CString(txt))

	return setter
}
