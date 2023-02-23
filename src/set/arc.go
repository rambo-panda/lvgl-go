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

type SetArc set

func (setter SetArc) SetStartAngle(start uint16) SetArc {
	C.lv_arc_set_start_angle(setter.CStructLvObjT, C.ushort(start))

	return setter
}
func (setter SetArc) SetEndAngle(end uint16) SetArc {
	C.lv_arc_set_end_angle(setter.CStructLvObjT, C.ushort(end))

	return setter
}
func (setter SetArc) SetAngles(start uint16, end uint16) SetArc {
	C.lv_arc_set_angles(setter.CStructLvObjT, C.ushort(start), C.ushort(end))

	return setter
}
func (setter SetArc) SetBgStartAngle(start uint16) SetArc {
	C.lv_arc_set_bg_start_angle(setter.CStructLvObjT, C.ushort(start))

	return setter
}
func (setter SetArc) SetBgEndAngle(end uint16) SetArc {
	C.lv_arc_set_bg_end_angle(setter.CStructLvObjT, C.ushort(end))

	return setter
}
func (setter SetArc) SetBgAngles(start uint16, end uint16) SetArc {
	C.lv_arc_set_bg_angles(setter.CStructLvObjT, C.ushort(start), C.ushort(end))

	return setter
}
func (setter SetArc) SetRotation(rotation uint16) SetArc {
	C.lv_arc_set_rotation(setter.CStructLvObjT, C.ushort(rotation))

	return setter
}
func (setter SetArc) SetMode(_type lib.LvArcModeT) SetArc {
	C.lv_arc_set_mode(setter.CStructLvObjT, C.lv_arc_mode_t(_type))

	return setter
}
func (setter SetArc) SetValue(value int16) SetArc {
	C.lv_arc_set_value(setter.CStructLvObjT, C.short(value))

	return setter
}
func (setter SetArc) SetRange(min int16, max int16) SetArc {
	C.lv_arc_set_range(setter.CStructLvObjT, C.short(min), C.short(max))

	return setter
}
func (setter SetArc) SetChangeRate(rate uint16) SetArc {
	C.lv_arc_set_change_rate(setter.CStructLvObjT, C.ushort(rate))

	return setter
}
