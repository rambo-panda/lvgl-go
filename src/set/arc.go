package set

import (
	lib.LV_FONT_T"lvgl-go/src/lib.LV_FONT_T
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Arc set[CObjT]

func CreateArc(o CObjT) Arc {
	return Arc{
		CObj: o,
	}
}

func (setter Arc) StartAngle(start uint16) Arc {
	C.lv_arc_set_start_angle(setter.CObj, C.ushort(start))

	return setter
}
func (setter Arc) EndAngle(end uint16) Arc {
	C.lv_arc_set_end_angle(setter.CObj, C.ushort(end))

	return setter
}
func (setter Arc) Angles(start uint16, end uint16) Arc {
	C.lv_arc_set_angles(setter.CObj, C.ushort(start), C.ushort(end))

	return setter
}
func (setter Arc) BgStartAngle(start uint16) Arc {
	C.lv_arc_set_bg_start_angle(setter.CObj, C.ushort(start))

	return setter
}
func (setter Arc) BgEndAngle(end uint16) Arc {
	C.lv_arc_set_bg_end_angle(setter.CObj, C.ushort(end))

	return setter
}
func (setter Arc) BgAngles(start uint16, end uint16) Arc {
	C.lv_arc_set_bg_angles(setter.CObj, C.ushort(start), C.ushort(end))

	return setter
}
func (setter Arc) Rotation(rotation uint16) Arc {
	C.lv_arc_set_rotation(setter.CObj, C.ushort(rotation))

	return setter
}
func (setter Arc) Mode(_type lib.LV_FONT_TLvArcModeT) Arc {
	C.lv_arc_set_mode(setter.CObj, C.lv_arc_mode_t(_type))

	return setter
}
func (setter Arc) Value(value int16) Arc {
	C.lv_arc_set_value(setter.CObj, C.short(value))

	return setter
}
func (setter Arc) Range(min int16, max int16) Arc {
	C.lv_arc_set_range(setter.CObj, C.short(min), C.short(max))

	return setter
}
func (setter Arc) ChangeRate(rate uint16) Arc {
	C.lv_arc_set_change_rate(setter.CObj, C.ushort(rate))

	return setter
}
