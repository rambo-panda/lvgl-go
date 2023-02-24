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

type Spangroup set

func CreateSpangroup(o *types.LvObjT) Spangroup {
	return Spangroup{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Spangroup) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter Spangroup) Align(align types.LvTextAlignT) Spangroup {
	C.lv_spangroup_set_align(setter.CStructLvObjT, C.lv_text_align_t(align))

	return setter
}
func (setter Spangroup) Overflow(overflow types.LvSpanOverflowT) Spangroup {
	C.lv_spangroup_set_overflow(setter.CStructLvObjT, C.lv_span_overflow_t(overflow))

	return setter
}
func (setter Spangroup) Indent(indent types.LvCoordT) Spangroup {
	C.lv_spangroup_set_indent(setter.CStructLvObjT, C.lv_coord_t(indent))

	return setter
}
func (setter Spangroup) Mode(mode types.LvSpanModeT) Spangroup {
	C.lv_spangroup_set_mode(setter.CStructLvObjT, C.lv_span_mode_t(mode))

	return setter
}
func (setter Spangroup) Lines(lines int32) Spangroup {
	C.lv_spangroup_set_lines(setter.CStructLvObjT, C.int(lines))

	return setter
}
