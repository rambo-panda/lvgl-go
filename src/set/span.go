package set

import (
	"lvgl-go/src/lib"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Span set[CObjT]

func CreateSpan(o CObjT) Span {
	return Span{
		CObj: o,
	}
}

func (setter Span) Text(span *lib.LV_SPAN_T, text string) Span {
	C.lv_span_set_text((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
func (setter Span) TextStatic(span *lib.LV_SPAN_T, text string) Span {
	C.lv_span_set_text_static((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
