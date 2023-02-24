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

type Span set

func CreateSpan(o *types.LvObjT) Span {
	return Span{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Span) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter Span) Text(span *types.LvSpanT, text string) Span {
	C.lv_span_set_text((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
func (setter Span) TextStatic(span *types.LvSpanT, text string) Span {
	C.lv_span_set_text_static((*C.lv_span_t)(unsafe.Pointer(span)), C.CString(text))

	return setter
}
