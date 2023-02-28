package set

import (
	types "lvgl-go/src/types"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Spangroup set[CObjT]

func CreateSpangroup(o CObjT) Spangroup {
	return Spangroup{
		CObj: o,
	}
}

func (setter Spangroup) Align(align types.LvTextAlignT) Spangroup {
	C.lv_spangroup_set_align(setter.CObj, C.lv_text_align_t(align))

	return setter
}
func (setter Spangroup) Overflow(overflow types.LvSpanOverflowT) Spangroup {
	C.lv_spangroup_set_overflow(setter.CObj, C.lv_span_overflow_t(overflow))

	return setter
}
func (setter Spangroup) Indent(indent types.LvCoordT) Spangroup {
	C.lv_spangroup_set_indent(setter.CObj, C.lv_coord_t(indent))

	return setter
}
func (setter Spangroup) Mode(mode types.LvSpanModeT) Spangroup {
	C.lv_spangroup_set_mode(setter.CObj, C.lv_span_mode_t(mode))

	return setter
}
func (setter Spangroup) Lines(lines int32) Spangroup {
	C.lv_spangroup_set_lines(setter.CObj, C.int(lines))

	return setter
}
