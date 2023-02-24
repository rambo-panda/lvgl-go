package get

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

type Obj get

func CreateObj(o *types.LvObjT) Obj {
	return Obj{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (getter Obj) StyleProp(part types.LvPartT, prop types.LvStylePropT) types.LvStyleValueT {
	res := C.lv_obj_get_style_prop(getter.CStructLvObjT, C.lv_part_t(part), C.lv_style_prop_t(prop))

	return types.LvStyleValueT(res)
}
func (getter Obj) State() types.LvStateT {
	res := C.lv_obj_get_state(getter.CStructLvObjT)

	return types.LvStateT(res)
}
func (getter Obj) Group() Obj {
	C.lv_obj_get_group(getter.CStructLvObjT)

	return getter
}
func (getter Obj) Class() *types.LvObjClassT {
	res := C.lv_obj_get_class(getter.CStructLvObjT)

	return (*types.LvObjClassT)(unsafe.Pointer(res))
}
func (getter Obj) Coords(coords *types.LvAreaT) Obj {
	C.lv_obj_get_coords(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(coords)))

	return getter
}
func (getter Obj) X() int16 {
	res := C.lv_obj_get_x(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) X2() int16 {
	res := C.lv_obj_get_x2(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) Y() int16 {
	res := C.lv_obj_get_y(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) Y2() int16 {
	res := C.lv_obj_get_y2(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) XAligned() int16 {
	res := C.lv_obj_get_x_aligned(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) YAligned() int16 {
	res := C.lv_obj_get_y_aligned(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) Width() int16 {
	res := C.lv_obj_get_width(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) Height() int16 {
	res := C.lv_obj_get_height(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ContentWidth() int16 {
	res := C.lv_obj_get_content_width(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ContentHeight() int16 {
	res := C.lv_obj_get_content_height(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ContentCoords(area *types.LvAreaT) Obj {
	C.lv_obj_get_content_coords(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(area)))

	return getter
}
func (getter Obj) SelfWidth() int16 {
	res := C.lv_obj_get_self_width(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) SelfHeight() int16 {
	res := C.lv_obj_get_self_height(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) TransformedArea(area *types.LvAreaT, recursive bool, inv bool) Obj {
	C.lv_obj_get_transformed_area(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(area)), C.bool(recursive), C.bool(inv))

	return getter
}
func (getter Obj) ClickArea(area *types.LvAreaT) Obj {
	C.lv_obj_get_click_area(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(area)))

	return getter
}
func (getter Obj) ScrollbarMode() types.LvScrollbarModeT {
	res := C.lv_obj_get_scrollbar_mode(getter.CStructLvObjT)

	return types.LvScrollbarModeT(res)
}
func (getter Obj) ScrollDir() types.LvDirT {
	res := C.lv_obj_get_scroll_dir(getter.CStructLvObjT)

	return types.LvDirT(res)
}
func (getter Obj) ScrollSnapX() types.LvScrollSnapT {
	res := C.lv_obj_get_scroll_snap_x(getter.CStructLvObjT)

	return types.LvScrollSnapT(res)
}
func (getter Obj) ScrollSnapY() types.LvScrollSnapT {
	res := C.lv_obj_get_scroll_snap_y(getter.CStructLvObjT)

	return types.LvScrollSnapT(res)
}
func (getter Obj) ScrollX() int16 {
	res := C.lv_obj_get_scroll_x(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ScrollY() int16 {
	res := C.lv_obj_get_scroll_y(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ScrollTop() int16 {
	res := C.lv_obj_get_scroll_top(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ScrollBottom() int16 {
	res := C.lv_obj_get_scroll_bottom(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ScrollLeft() int16 {
	res := C.lv_obj_get_scroll_left(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ScrollRight() int16 {
	res := C.lv_obj_get_scroll_right(getter.CStructLvObjT)

	return int16(res)
}
func (getter Obj) ScrollEnd(obj *types.LvObjT, end *types.LvPointT) Obj {
	C.lv_obj_get_scroll_end((*C.lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_point_t)(unsafe.Pointer(end)))

	return getter
}
func (getter Obj) ScrollbarArea(hor *types.LvAreaT, ver *types.LvAreaT) Obj {
	C.lv_obj_get_scrollbar_area(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(hor)), (*C.lv_area_t)(unsafe.Pointer(ver)))

	return getter
}
func (getter Obj) Disp() *types.LvDispT {
	res := C.lv_obj_get_disp(getter.CStructLvObjT)

	return (*types.LvDispT)(unsafe.Pointer(res))
}
func (getter Obj) ChildCnt() uint32 {
	res := C.lv_obj_get_child_cnt(getter.CStructLvObjT)

	return uint32(res)
}
func (getter Obj) Index() uint32 {
	res := C.lv_obj_get_index(getter.CStructLvObjT)

	return uint32(res)
}

// func (getter Obj) EventUserData(event_cb types.LvEventCbT) Obj {
// 	C.lv_obj_get_event_user_data(getter.CStructLvObjT, C.lv_event_cb_t(event_cb))

// 	return getter
// }
