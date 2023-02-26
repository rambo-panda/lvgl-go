package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Checkbox set

func CreateCheckbox(o CObjT) Checkbox {
	return Checkbox{
		CObj: o,
	}
}

func (setter Checkbox) Text(txt string) Checkbox {
	C.lv_checkbox_set_text(setter.CObj, C.CString(txt))

	return setter
}
func (setter Checkbox) TextStatic(txt string) Checkbox {
	C.lv_checkbox_set_text_static(setter.CObj, C.CString(txt))

	return setter
}
