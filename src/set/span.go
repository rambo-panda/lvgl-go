package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
	"unsafe"
)

type SetSpan set

func (setter SetSpan) SetText(span *lib.LvSpanT, text string) SetSpan {
	C.lv_span_set_text((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
func (setter SetSpan) SetTextStatic(span *lib.LvSpanT, text string) SetSpan {
	C.lv_span_set_text_static((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
