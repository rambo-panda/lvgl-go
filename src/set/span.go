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

func CreateSpan(o *lib.LvObjT) SetSpan {
	return SetSpan{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetSpan) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetSpan) SetText(span *lib.LvSpanT, text string) SetSpan {
	C.lv_span_set_text((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
func (setter SetSpan) SetTextStatic(span *lib.LvSpanT, text string) SetSpan {
	C.lv_span_set_text_static((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
