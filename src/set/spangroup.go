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

type SetSpangroup set

func (setter SetSpangroup) SetAlign(align lib.LvTextAlignT) SetSpangroup {
	C.lv_spangroup_set_align(setter.CStructLvObjT, C.lv_text_align_t(align))

	return setter
}
func (setter SetSpangroup) SetOverflow(overflow lib.LvSpanOverflowT) SetSpangroup {
	C.lv_spangroup_set_overflow(setter.CStructLvObjT, C.lv_span_overflow_t(overflow))

	return setter
}
func (setter SetSpangroup) SetIndent(indent lib.LvCoordT) SetSpangroup {
	C.lv_spangroup_set_indent(setter.CStructLvObjT, C.lv_coord_t(indent))

	return setter
}
func (setter SetSpangroup) SetMode(mode lib.LvSpanModeT) SetSpangroup {
	C.lv_spangroup_set_mode(setter.CStructLvObjT, C.lv_span_mode_t(mode))

	return setter
}
func (setter SetSpangroup) SetLines(lines int32) SetSpangroup {
	C.lv_spangroup_set_lines(setter.CStructLvObjT, C.int(lines))

	return setter
}
