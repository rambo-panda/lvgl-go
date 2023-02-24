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

type SetSpangroup set

func CreateSpangroup(o *types.LvObjT) SetSpangroup {
	return SetSpangroup{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetSpangroup) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetSpangroup) SetAlign(align types.LvTextAlignT) SetSpangroup {
	C.lv_spangroup_set_align(setter.CStructLvObjT, C.lv_text_align_t(align))

	return setter
}
func (setter SetSpangroup) SetOverflow(overflow types.LvSpanOverflowT) SetSpangroup {
	C.lv_spangroup_set_overflow(setter.CStructLvObjT, C.lv_span_overflow_t(overflow))

	return setter
}
func (setter SetSpangroup) SetIndent(indent types.LvCoordT) SetSpangroup {
	C.lv_spangroup_set_indent(setter.CStructLvObjT, C.lv_coord_t(indent))

	return setter
}
func (setter SetSpangroup) SetMode(mode types.LvSpanModeT) SetSpangroup {
	C.lv_spangroup_set_mode(setter.CStructLvObjT, C.lv_span_mode_t(mode))

	return setter
}
func (setter SetSpangroup) SetLines(lines int32) SetSpangroup {
	C.lv_spangroup_set_lines(setter.CStructLvObjT, C.int(lines))

	return setter
}
