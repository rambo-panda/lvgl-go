package set

import (
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
)

/*
#cgo CFLAGS: -I../include/
#include "lv_init.h"
*/
import "C"

type Label set[CObjT]

func CreateLabel(o CObjT) *Label {
	_o := Label{
		CObj: o,
	}
	return &_o
}
func (setter *Label) Text(text string) *Label {
	j := C.CString(text)
	C.lv_label_set_text(setter.CObj, j)

	return setter
}
func (setter *Label) TextStatic(text string) *Label {
	C.lv_label_set_text_static(setter.CObj, C.CString(text))

	return setter
}
func (setter *Label) LongMode(long_mode lib.LV_LABEL_LONG_MODE_T) *Label {
	C.lv_label_set_long_mode(setter.CObj, C.lv_label_long_mode_t(long_mode))

	return setter
}
func (setter *Label) Recolor(en bool) *Label {
	C.lv_label_set_recolor(setter.CObj, C.bool(en))

	return setter
}
func (setter *Label) TextSelStart(index uint32) *Label {
	C.lv_label_set_text_sel_start(setter.CObj, C.uint(index))

	return setter
}
func (setter *Label) TextSelEnd(index uint32) *Label {
	C.lv_label_set_text_sel_end(setter.CObj, C.uint(index))

	return setter
}
