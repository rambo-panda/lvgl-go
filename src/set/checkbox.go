package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type SetCheckbox set

func (setter SetCheckbox) SetText(txt string) SetCheckbox {
	C.lv_checkbox_set_text(setter.cObj, C.CString(txt))

	return setter
}
func (setter SetCheckbox) SetTextStatic(txt string) SetCheckbox {
	C.lv_checkbox_set_text_static(setter.cObj, C.CString(txt))

	return setter
}
