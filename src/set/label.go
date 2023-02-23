package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
)

type SetLabel set

func (setter SetLabel) SetText(text string) SetLabel {
	C.lv_label_set_text(setter.cObj, C.CString(text))

	return setter
}
func (setter SetLabel) SetTextStatic(text string) SetLabel {
	C.lv_label_set_text_static(setter.cObj, C.CString(text))

	return setter
}
func (setter SetLabel) SetLongMode(long_mode lib.LvLabelLongModeT) SetLabel {
	C.lv_label_set_long_mode(setter.cObj, C.lv_label_long_mode_t(long_mode))

	return setter
}
func (setter SetLabel) SetRecolor(en bool) SetLabel {
	C.lv_label_set_recolor(setter.cObj, C.bool(en))

	return setter
}
func (setter SetLabel) SetTextSelStart(index uint32) SetLabel {
	C.lv_label_set_text_sel_start(setter.cObj, C.uint(index))

	return setter
}
func (setter SetLabel) SetTextSelEnd(index uint32) SetLabel {
	C.lv_label_set_text_sel_end(setter.cObj, C.uint(index))

	return setter
}
