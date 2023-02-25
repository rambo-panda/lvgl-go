package set

import (
	types "lvgl-go/src/types"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Arc set

func CreateArc(o CObjT) Arc {
	return Arc{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Arc) StartAngle(start uint16) Arc {
	C.lv_arc_set_start_angle(setter.CStructLvObjT, C.ushort(start))

	return setter
}
func (setter Arc) EndAngle(end uint16) Arc {
	C.lv_arc_set_end_angle(setter.CStructLvObjT, C.ushort(end))

	return setter
}
func (setter Arc) Angles(start uint16, end uint16) Arc {
	C.lv_arc_set_angles(setter.CStructLvObjT, C.ushort(start), C.ushort(end))

	return setter
}
func (setter Arc) BgStartAngle(start uint16) Arc {
	C.lv_arc_set_bg_start_angle(setter.CStructLvObjT, C.ushort(start))

	return setter
}
func (setter Arc) BgEndAngle(end uint16) Arc {
	C.lv_arc_set_bg_end_angle(setter.CStructLvObjT, C.ushort(end))

	return setter
}
func (setter Arc) BgAngles(start uint16, end uint16) Arc {
	C.lv_arc_set_bg_angles(setter.CStructLvObjT, C.ushort(start), C.ushort(end))

	return setter
}
func (setter Arc) Rotation(rotation uint16) Arc {
	C.lv_arc_set_rotation(setter.CStructLvObjT, C.ushort(rotation))

	return setter
}
func (setter Arc) Mode(_type types.LvArcModeT) Arc {
	C.lv_arc_set_mode(setter.CStructLvObjT, C.lv_arc_mode_t(_type))

	return setter
}
func (setter Arc) Value(value int16) Arc {
	C.lv_arc_set_value(setter.CStructLvObjT, C.short(value))

	return setter
}
func (setter Arc) Range(min int16, max int16) Arc {
	C.lv_arc_set_range(setter.CStructLvObjT, C.short(min), C.short(max))

	return setter
}
func (setter Arc) ChangeRate(rate uint16) Arc {
	C.lv_arc_set_change_rate(setter.CStructLvObjT, C.ushort(rate))

	return setter
}
