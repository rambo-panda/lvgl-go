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

type SetStyle set

func CreateStyle(o *types.LvObjT) SetStyle {
	return SetStyle{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetStyle) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetStyle) SetWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMinWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_min_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMaxWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_max_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetHeight(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMinHeight(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_min_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMaxHeight(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_max_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetX(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetY(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetAlign(style *types.LvStyleT, value types.LvAlignT) SetStyle {
	C.lv_style_set_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_align_t(value))

	return setter
}
func (setter SetStyle) SetTransformWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_transform_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformHeight(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_transform_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTranslateX(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_translate_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTranslateY(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_translate_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformZoom(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_transform_zoom((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformAngle(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_transform_angle((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformPivotX(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_transform_pivot_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformPivotY(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_transform_pivot_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadTop(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_pad_top((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadBottom(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_pad_bottom((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadLeft(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_pad_left((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadRight(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_pad_right((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadRow(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_pad_row((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadColumn(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_pad_column((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBgColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_bg_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBgOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_bg_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBgGradColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_bg_grad_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBgGradDir(style *types.LvStyleT, value types.LvGradDirT) SetStyle {
	C.lv_style_set_bg_grad_dir((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grad_dir_t(value))

	return setter
}
func (setter SetStyle) SetBgMainStop(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_bg_main_stop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBgGradStop(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_bg_grad_stop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBgGrad(style *types.LvStyleT, value *types.LvGradDscT) SetStyle {
	C.lv_style_set_bg_grad((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetBgDitherMode(style *types.LvStyleT, value types.LvDitherModeT) SetStyle {
	C.lv_style_set_bg_dither_mode((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_dither_mode_t(value))

	return setter
}
func (setter SetStyle) SetBgImgSrc(style *types.LvStyleT, value any) SetStyle {
	C.lv_style_set_bg_img_src((*C.lv_style_t)(unsafe.Pointer(style)), unsafe.Pointer(&value))

	return setter
}
func (setter SetStyle) SetBgImgOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_bg_img_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBgImgRecolor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_bg_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBgImgRecolorOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_bg_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBgImgTiled(style *types.LvStyleT, value bool) SetStyle {
	C.lv_style_set_bg_img_tiled((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetBorderColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_border_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBorderOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_border_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBorderWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_border_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBorderSide(style *types.LvStyleT, value types.LvBorderSideT) SetStyle {
	C.lv_style_set_border_side((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_border_side_t(value))

	return setter
}
func (setter SetStyle) SetBorderPost(style *types.LvStyleT, value bool) SetStyle {
	C.lv_style_set_border_post((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetOutlineWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_outline_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetOutlineColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_outline_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetOutlineOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_outline_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetOutlinePad(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_outline_pad((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_shadow_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowOfsX(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_shadow_ofs_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowOfsY(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_shadow_ofs_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowSpread(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_shadow_spread((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_shadow_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetShadowOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_shadow_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetImgOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_img_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetImgRecolor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetImgRecolorOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetLineWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_line_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetLineDashWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_line_dash_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetLineDashGap(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_line_dash_gap((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetLineRounded(style *types.LvStyleT, value bool) SetStyle {
	C.lv_style_set_line_rounded((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetLineColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_line_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetLineOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_line_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetArcWidth(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_arc_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetArcRounded(style *types.LvStyleT, value bool) SetStyle {
	C.lv_style_set_arc_rounded((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetArcColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_arc_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetArcOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_arc_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetArcImgSrc(style *types.LvStyleT, value any) SetStyle {
	C.lv_style_set_arc_img_src((*C.lv_style_t)(unsafe.Pointer(style)), unsafe.Pointer(&value))

	return setter
}
func (setter SetStyle) SetTextColor(style *types.LvStyleT, value types.LvColorT) SetStyle {
	C.lv_style_set_text_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetTextOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_text_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetTextFont(style *types.LvStyleT, value *types.LvFontT) SetStyle {
	C.lv_style_set_text_font((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_font_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetTextLetterSpace(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_text_letter_space((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTextLineSpace(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_text_line_space((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTextDecor(style *types.LvStyleT, value types.LvTextDecorT) SetStyle {
	C.lv_style_set_text_decor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_text_decor_t(value))

	return setter
}
func (setter SetStyle) SetTextAlign(style *types.LvStyleT, value types.LvTextAlignT) SetStyle {
	C.lv_style_set_text_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_text_align_t(value))

	return setter
}
func (setter SetStyle) SetRadius(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_radius((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetClipCorner(style *types.LvStyleT, value bool) SetStyle {
	C.lv_style_set_clip_corner((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetColorFilterDsc(style *types.LvStyleT, value *types.LvColorFilterDscT) SetStyle {
	C.lv_style_set_color_filter_dsc((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetColorFilterOpa(style *types.LvStyleT, value types.LvOpaT) SetStyle {
	C.lv_style_set_color_filter_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetAnim(style *types.LvStyleT, value *types.LvAnimT) SetStyle {
	C.lv_style_set_anim((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_anim_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetAnimTime(style *types.LvStyleT, value uint32) SetStyle {
	C.lv_style_set_anim_time((*C.lv_style_t)(unsafe.Pointer(style)), C.uint(value))

	return setter
}
func (setter SetStyle) SetAnimSpeed(style *types.LvStyleT, value uint32) SetStyle {
	C.lv_style_set_anim_speed((*C.lv_style_t)(unsafe.Pointer(style)), C.uint(value))

	return setter
}
func (setter SetStyle) SetTransition(style *types.LvStyleT, value *types.LvStyleTransitionDscT) SetStyle {
	C.lv_style_set_transition((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetBlendMode(style *types.LvStyleT, value types.LvBlendModeT) SetStyle {
	C.lv_style_set_blend_mode((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_blend_mode_t(value))

	return setter
}
func (setter SetStyle) SetLayout(style *types.LvStyleT, value uint16) SetStyle {
	C.lv_style_set_layout((*C.lv_style_t)(unsafe.Pointer(style)), C.ushort(value))

	return setter
}
func (setter SetStyle) SetBaseDir(style *types.LvStyleT, value types.LvBaseDirT) SetStyle {
	C.lv_style_set_base_dir((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_base_dir_t(value))

	return setter
}
func (setter SetStyle) SetProp(style *types.LvStyleT, prop types.LvStylePropT, value types.LvStyleValueT) SetStyle {
	C.lv_style_set_prop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_style_prop_t(prop), C.lv_style_value_t(value))

	return setter
}
func (setter SetStyle) SetPropMeta(style *types.LvStyleT, prop types.LvStylePropT, meta uint16) SetStyle {
	C.lv_style_set_prop_meta((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_style_prop_t(prop), C.ushort(meta))

	return setter
}
func (setter SetStyle) SetFlexFlow(style *types.LvStyleT, value types.LvFlexFlowT) SetStyle {
	C.lv_style_set_flex_flow((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_flow_t(value))

	return setter
}
func (setter SetStyle) SetFlexMainPlace(style *types.LvStyleT, value types.LvFlexAlignT) SetStyle {
	C.lv_style_set_flex_main_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter SetStyle) SetFlexCrossPlace(style *types.LvStyleT, value types.LvFlexAlignT) SetStyle {
	C.lv_style_set_flex_cross_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter SetStyle) SetFlexTrackPlace(style *types.LvStyleT, value types.LvFlexAlignT) SetStyle {
	C.lv_style_set_flex_track_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter SetStyle) SetFlexGrow(style *types.LvStyleT, value uint8) SetStyle {
	C.lv_style_set_flex_grow((*C.lv_style_t)(unsafe.Pointer(style)), C.uint8_t(value))

	return setter
}
func (setter SetStyle) SetGridRowAlign(style *types.LvStyleT, value types.LvGridAlignT) SetStyle {
	C.lv_style_set_grid_row_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

	return setter
}
func (setter SetStyle) SetGridColumnAlign(style *types.LvStyleT, value types.LvGridAlignT) SetStyle {
	C.lv_style_set_grid_column_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

	return setter
}
func (setter SetStyle) SetGridCellColumnPos(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_column_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellColumnSpan(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_column_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellRowPos(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_row_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellRowSpan(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_row_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellXAlign(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_x_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellYAlign(style *types.LvStyleT, value types.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_y_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
