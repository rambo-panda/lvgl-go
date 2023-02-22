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

func StartAngleForArc(obj *lib.LvObjT, start uint16) {
	C.lv_arc_set_start_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(start))

}
func EndAngleForArc(obj *lib.LvObjT, end uint16) {
	C.lv_arc_set_end_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(end))

}
func AnglesForArc(obj *lib.LvObjT, start uint16, end uint16) {
	C.lv_arc_set_angles((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(start), C.ushort(end))

}
func BgStartAngleForArc(obj *lib.LvObjT, start uint16) {
	C.lv_arc_set_bg_start_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(start))

}
func BgEndAngleForArc(obj *lib.LvObjT, end uint16) {
	C.lv_arc_set_bg_end_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(end))

}
func BgAnglesForArc(obj *lib.LvObjT, start uint16, end uint16) {
	C.lv_arc_set_bg_angles((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(start), C.ushort(end))

}
func RotationForArc(obj *lib.LvObjT, rotation uint16) {
	C.lv_arc_set_rotation((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(rotation))

}
func ModeForArc(obj *lib.LvObjT, t lib.LvArcModeT) {
	C.lv_arc_set_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_arc_mode_t(t))

}
func ValueForArc(obj *lib.LvObjT, value int16) {
	C.lv_arc_set_value((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(value))

}
func RangeForArc(obj *lib.LvObjT, min int16, max int16) {
	C.lv_arc_set_range((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(min), C.short(max))

}
func ChangeRateForArc(obj *lib.LvObjT, rate uint16) {
	C.lv_arc_set_change_rate((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(rate))

}
