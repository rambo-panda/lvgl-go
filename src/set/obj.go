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

type SetObj set

func CreateObj(o *types.LvObjT) SetObj {
	return SetObj{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetObj) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetObj) SetLocalStyleProp(obj *types.LvObjT, prop types.LvStylePropT, value types.LvStyleValueT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_local_style_prop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetLocalStylePropMeta(obj *types.LvObjT, prop types.LvStylePropT, meta uint16, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_local_style_prop_meta((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMinWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_min_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMaxWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_max_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMinHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_min_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMaxHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_max_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAlign(obj *types.LvObjT, value types.LvAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTranslateX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_translate_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTranslateY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_translate_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformZoom(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_zoom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformAngle(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformPivotX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_pivot_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformPivotY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_pivot_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadTop(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_top((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadBottom(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_bottom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadLeft(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_left((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadRight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_right((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadRow(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_row((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadColumn(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_column((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradDir(obj *types.LvObjT, value types.LvGradDirT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgMainStop(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_main_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradStop(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGrad(obj *types.LvObjT, value *types.LvGradDscT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgDitherMode(obj *types.LvObjT, value types.LvDitherModeT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_dither_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgSrc(obj *types.LvObjT, value any, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgRecolor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgRecolorOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgTiled(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_tiled((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderSide(obj *types.LvObjT, value types.LvBorderSideT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_side((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderPost(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_post((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlinePad(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_pad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOfsX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_ofs_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOfsY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_ofs_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowSpread(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_spread((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgRecolor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgRecolorOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineDashWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_dash_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineDashGap(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_dash_gap((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineRounded(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcRounded(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcImgSrc(obj *types.LvObjT, value any, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextFont(obj *types.LvObjT, value *types.LvFontT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_font((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextLetterSpace(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_letter_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextLineSpace(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_line_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextDecor(obj *types.LvObjT, value types.LvTextDecorT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_decor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextAlign(obj *types.LvObjT, value types.LvTextAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleRadius(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_radius((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleClipCorner(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_clip_corner((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleColorFilterDsc(obj *types.LvObjT, value *types.LvColorFilterDscT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_color_filter_dsc((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleColorFilterOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_color_filter_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnim(obj *types.LvObjT, value *types.LvAnimT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnimTime(obj *types.LvObjT, value uint32, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim_time((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnimSpeed(obj *types.LvObjT, value uint32, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim_speed((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransition(obj *types.LvObjT, value *types.LvStyleTransitionDscT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transition((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBlendMode(obj *types.LvObjT, value types.LvBlendModeT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_blend_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLayout(obj *types.LvObjT, value uint16, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBaseDir(obj *types.LvObjT, value types.LvBaseDirT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_base_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetPos(obj *types.LvObjT, x types.LvCoordT, y types.LvCoordT) SetObj {
	C.lv_obj_set_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter SetObj) SetX(obj *types.LvObjT, x types.LvCoordT) SetObj {
	C.lv_obj_set_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x))

	return setter
}
func (setter SetObj) SetY(obj *types.LvObjT, y types.LvCoordT) SetObj {
	C.lv_obj_set_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(y))

	return setter
}
func (setter SetObj) SetSize(obj *types.LvObjT, w types.LvCoordT, h types.LvCoordT) SetObj {
	C.lv_obj_set_size((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetWidth(obj *types.LvObjT, w types.LvCoordT) SetObj {
	C.lv_obj_set_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	return setter
}
func (setter SetObj) SetHeight(obj *types.LvObjT, h types.LvCoordT) SetObj {
	C.lv_obj_set_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetContentWidth(obj *types.LvObjT, w types.LvCoordT) SetObj {
	C.lv_obj_set_content_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	return setter
}
func (setter SetObj) SetContentHeight(obj *types.LvObjT, h types.LvCoordT) SetObj {
	C.lv_obj_set_content_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetLayout(obj *types.LvObjT, layout uint32) SetObj {
	C.lv_obj_set_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(layout))

	return setter
}
func (setter SetObj) SetAlign(obj *types.LvObjT, align types.LvAlignT) SetObj {
	C.lv_obj_set_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(align))

	return setter
}
func (setter SetObj) SetExtClickArea(obj *types.LvObjT, size types.LvCoordT) SetObj {
	C.lv_obj_set_ext_click_area((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(size))

	return setter
}
func (setter SetObj) SetScrollbarMode(obj *types.LvObjT, mode types.LvScrollbarModeT) SetObj {
	C.lv_obj_set_scrollbar_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter SetObj) SetScrollDir(obj *types.LvObjT, dir types.LvDirT) SetObj {
	C.lv_obj_set_scroll_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dir_t(dir))

	return setter
}
func (setter SetObj) SetScrollSnapX(obj *types.LvObjT, align types.LvScrollSnapT) SetObj {
	C.lv_obj_set_scroll_snap_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

	return setter
}
func (setter SetObj) SetScrollSnapY(obj *types.LvObjT, align types.LvScrollSnapT) SetObj {
	C.lv_obj_set_scroll_snap_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

	return setter
}
func (setter SetObj) SetParent(obj *types.LvObjT, parent *types.LvObjT) SetObj {
	C.lv_obj_set_parent((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter SetObj) SetFlexFlow(flow types.LvFlexFlowT) SetObj {
	C.lv_obj_set_flex_flow(setter.CStructLvObjT, C.lv_flex_flow_t(flow))

	return setter
}
func (setter SetObj) SetFlexAlign(main_place types.LvFlexAlignT, cross_place types.LvFlexAlignT, track_cross_place types.LvFlexAlignT) SetObj {
	C.lv_obj_set_flex_align(setter.CStructLvObjT, C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

	return setter
}
func (setter SetObj) SetFlexGrow(grow uint8) SetObj {
	C.lv_obj_set_flex_grow(setter.CStructLvObjT, C.uint8_t(grow))

	return setter
}
func (setter SetObj) SetStyleFlexFlow(value types.LvFlexFlowT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_flow(setter.CStructLvObjT, C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexMainPlace(value types.LvFlexAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_main_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexCrossPlace(value types.LvFlexAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_cross_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexTrackPlace(value types.LvFlexAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_track_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexGrow(value uint8, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_grow(setter.CStructLvObjT, C.uint8_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetGridAlign(column_align types.LvGridAlignT, row_align types.LvGridAlignT) SetObj {
	C.lv_obj_set_grid_align(setter.CStructLvObjT, C.lv_grid_align_t(column_align), C.lv_grid_align_t(row_align))

	return setter
}
func (setter SetObj) SetGridCell(column_align types.LvGridAlignT, col_pos uint8, col_span uint8, row_align types.LvGridAlignT, row_pos uint8, row_span uint8) SetObj {
	C.lv_obj_set_grid_cell(setter.CStructLvObjT, C.lv_grid_align_t(column_align), C.uint8_t(col_pos), C.uint8_t(col_span), C.lv_grid_align_t(row_align), C.uint8_t(row_pos), C.uint8_t(row_span))

	return setter
}
func (setter SetObj) SetStyleGridRowAlign(value types.LvGridAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_row_align(setter.CStructLvObjT, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridColumnAlign(value types.LvGridAlignT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_column_align(setter.CStructLvObjT, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellColumnPos(value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_column_pos(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellColumnSpan(value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_column_span(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellRowPos(value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_row_pos(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellRowSpan(value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_row_span(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellXAlign(value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_x_align(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellYAlign(value types.LvCoordT, selector types.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_y_align(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetTile(anim_en types.LvAnimEnableT) SetObj {
	C.lv_obj_set_tile(setter.CStructLvObjT, setter.CStructLvObjT, C.lv_anim_enable_t(anim_en))

	return setter
}
func (setter SetObj) SetTileId(col_id uint32, row_id uint32, anim_en types.LvAnimEnableT) SetObj {
	C.lv_obj_set_tile_id(setter.CStructLvObjT, C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

	return setter
}
