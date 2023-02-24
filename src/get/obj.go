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

type GetObj Get

func CreateObj(o *types.LvObjT) GetObj {
	return GetObj{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}
func (setter GetObj) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (getter GetObj) GetStyleProp(part types.LvPartT, prop types.LvStylePropT) types.LvStyleValueT {
	res := C.lv_obj_get_style_prop(getter.CStructLvObjT, C.lv_part_t(part), C.lv_style_prop_t(prop))

	return types.LvStyleValueT(res)
}
func (getter GetObj) GetState() types.LvStateT {
	res := C.lv_obj_get_state(getter.CStructLvObjT)

	return types.LvStateT(res)
}
func (getter GetObj) GetGroup() GetObj {
	C.lv_obj_get_group(getter.CStructLvObjT)

	return getter
}
func (getter GetObj) GetClass() *types.LvObjClassT {
	res := C.lv_obj_get_class(getter.CStructLvObjT)

	return (*types.LvObjClassT)(unsafe.Pointer(res))
}
func (getter GetObj) GetCoords(coords *types.LvAreaT) GetObj {
	C.lv_obj_get_coords(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(coords)))

	return getter
}
func (getter GetObj) GetX() int16 {
	res := C.lv_obj_get_x(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetX2() int16 {
	res := C.lv_obj_get_x2(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetY() int16 {
	res := C.lv_obj_get_y(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetY2() int16 {
	res := C.lv_obj_get_y2(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetXAligned() int16 {
	res := C.lv_obj_get_x_aligned(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetYAligned() int16 {
	res := C.lv_obj_get_y_aligned(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetWidth() int16 {
	res := C.lv_obj_get_width(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetHeight() int16 {
	res := C.lv_obj_get_height(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetContentWidth() int16 {
	res := C.lv_obj_get_content_width(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetContentHeight() int16 {
	res := C.lv_obj_get_content_height(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetContentCoords(area *types.LvAreaT) GetObj {
	C.lv_obj_get_content_coords(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(area)))

	return getter
}
func (getter GetObj) GetSelfWidth() int16 {
	res := C.lv_obj_get_self_width(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetSelfHeight() int16 {
	res := C.lv_obj_get_self_height(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetTransformedArea(area *types.LvAreaT, recursive bool, inv bool) GetObj {
	C.lv_obj_get_transformed_area(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(area)), C.bool(recursive), C.bool(inv))

	return getter
}
func (getter GetObj) GetClickArea(area *types.LvAreaT) GetObj {
	C.lv_obj_get_click_area(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(area)))

	return getter
}
func (getter GetObj) GetScrollbarMode() types.LvScrollbarModeT {
	res := C.lv_obj_get_scrollbar_mode(getter.CStructLvObjT)

	return types.LvScrollbarModeT(res)
}
func (getter GetObj) GetScrollDir() types.LvDirT {
	res := C.lv_obj_get_scroll_dir(getter.CStructLvObjT)

	return types.LvDirT(res)
}
func (getter GetObj) GetScrollSnapX() types.LvScrollSnapT {
	res := C.lv_obj_get_scroll_snap_x(getter.CStructLvObjT)

	return types.LvScrollSnapT(res)
}
func (getter GetObj) GetScrollSnapY() types.LvScrollSnapT {
	res := C.lv_obj_get_scroll_snap_y(getter.CStructLvObjT)

	return types.LvScrollSnapT(res)
}
func (getter GetObj) GetScrollX() int16 {
	res := C.lv_obj_get_scroll_x(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetScrollY() int16 {
	res := C.lv_obj_get_scroll_y(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetScrollTop() int16 {
	res := C.lv_obj_get_scroll_top(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetScrollBottom() int16 {
	res := C.lv_obj_get_scroll_bottom(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetScrollLeft() int16 {
	res := C.lv_obj_get_scroll_left(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetScrollRight() int16 {
	res := C.lv_obj_get_scroll_right(getter.CStructLvObjT)

	return int16(res)
}
func (getter GetObj) GetScrollEnd(obj *types.LvObjT, end *types.LvPointT) GetObj {
	C.lv_obj_get_scroll_end((*C.lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_point_t)(unsafe.Pointer(end)))

	return getter
}
func (getter GetObj) GetScrollbarArea(hor *types.LvAreaT, ver *types.LvAreaT) GetObj {
	C.lv_obj_get_scrollbar_area(getter.CStructLvObjT, (*C.lv_area_t)(unsafe.Pointer(hor)), (*C.lv_area_t)(unsafe.Pointer(ver)))

	return getter
}
func (getter GetObj) GetDisp() *types.LvDispT {
	res := C.lv_obj_get_disp(getter.CStructLvObjT)

	return (*types.LvDispT)(unsafe.Pointer(res))
}
func (getter GetObj) GetChildCnt() uint32 {
	res := C.lv_obj_get_child_cnt(getter.CStructLvObjT)

	return uint32(res)
}
func (getter GetObj) GetIndex() uint32 {
	res := C.lv_obj_get_index(getter.CStructLvObjT)

	return uint32(res)
}

// func (getter GetObj) GetEventUserData(event_cb types.LvEventCbT) GetObj {
// 	C.lv_obj_get_event_user_data(getter.CStructLvObjT, C.lv_event_cb_t(event_cb))

// 	return getter
// }
