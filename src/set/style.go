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

type SetStyle set

func (setter SetStyle) SetWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMinWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_min_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMaxWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_max_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetHeight(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMinHeight(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_min_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetMaxHeight(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_max_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetX(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetY(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetAlign(style *lib.LvStyleT, value lib.LvAlignT) SetStyle {
	C.lv_style_set_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_align_t(value))

	return setter
}
func (setter SetStyle) SetTransformWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_transform_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformHeight(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_transform_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTranslateX(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_translate_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTranslateY(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_translate_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformZoom(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_transform_zoom((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformAngle(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_transform_angle((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformPivotX(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_transform_pivot_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTransformPivotY(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_transform_pivot_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadTop(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_pad_top((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadBottom(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_pad_bottom((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadLeft(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_pad_left((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadRight(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_pad_right((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadRow(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_pad_row((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetPadColumn(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_pad_column((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBgColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_bg_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBgOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_bg_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBgGradColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_bg_grad_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBgGradDir(style *lib.LvStyleT, value lib.LvGradDirT) SetStyle {
	C.lv_style_set_bg_grad_dir((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grad_dir_t(value))

	return setter
}
func (setter SetStyle) SetBgMainStop(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_bg_main_stop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBgGradStop(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_bg_grad_stop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBgGrad(style *lib.LvStyleT, value *lib.LvGradDscT) SetStyle {
	C.lv_style_set_bg_grad((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetBgDitherMode(style *lib.LvStyleT, value lib.LvDitherModeT) SetStyle {
	C.lv_style_set_bg_dither_mode((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_dither_mode_t(value))

	return setter
}
func (setter SetStyle) SetBgImgSrc(style *lib.LvStyleT, value any) SetStyle {
	C.lv_style_set_bg_img_src((*C.lv_style_t)(unsafe.Pointer(style)), unsafe.Pointer(&value))

	return setter
}
func (setter SetStyle) SetBgImgOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_bg_img_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBgImgRecolor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_bg_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBgImgRecolorOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_bg_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBgImgTiled(style *lib.LvStyleT, value bool) SetStyle {
	C.lv_style_set_bg_img_tiled((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetBorderColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_border_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetBorderOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_border_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetBorderWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_border_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetBorderSide(style *lib.LvStyleT, value lib.LvBorderSideT) SetStyle {
	C.lv_style_set_border_side((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_border_side_t(value))

	return setter
}
func (setter SetStyle) SetBorderPost(style *lib.LvStyleT, value bool) SetStyle {
	C.lv_style_set_border_post((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetOutlineWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_outline_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetOutlineColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_outline_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetOutlineOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_outline_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetOutlinePad(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_outline_pad((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_shadow_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowOfsX(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_shadow_ofs_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowOfsY(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_shadow_ofs_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowSpread(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_shadow_spread((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetShadowColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_shadow_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetShadowOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_shadow_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetImgOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_img_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetImgRecolor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetImgRecolorOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetLineWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_line_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetLineDashWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_line_dash_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetLineDashGap(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_line_dash_gap((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetLineRounded(style *lib.LvStyleT, value bool) SetStyle {
	C.lv_style_set_line_rounded((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetLineColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_line_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetLineOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_line_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetArcWidth(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_arc_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetArcRounded(style *lib.LvStyleT, value bool) SetStyle {
	C.lv_style_set_arc_rounded((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetArcColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_arc_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetArcOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_arc_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetArcImgSrc(style *lib.LvStyleT, value any) SetStyle {
	C.lv_style_set_arc_img_src((*C.lv_style_t)(unsafe.Pointer(style)), unsafe.Pointer(&value))

	return setter
}
func (setter SetStyle) SetTextColor(style *lib.LvStyleT, value lib.LvColorT) SetStyle {
	C.lv_style_set_text_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter SetStyle) SetTextOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_text_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetTextFont(style *lib.LvStyleT, value *lib.LvFontT) SetStyle {
	C.lv_style_set_text_font((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_font_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetTextLetterSpace(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_text_letter_space((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTextLineSpace(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_text_line_space((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetTextDecor(style *lib.LvStyleT, value lib.LvTextDecorT) SetStyle {
	C.lv_style_set_text_decor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_text_decor_t(value))

	return setter
}
func (setter SetStyle) SetTextAlign(style *lib.LvStyleT, value lib.LvTextAlignT) SetStyle {
	C.lv_style_set_text_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_text_align_t(value))

	return setter
}
func (setter SetStyle) SetRadius(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_radius((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetClipCorner(style *lib.LvStyleT, value bool) SetStyle {
	C.lv_style_set_clip_corner((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter SetStyle) SetOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetColorFilterDsc(style *lib.LvStyleT, value *lib.LvColorFilterDscT) SetStyle {
	C.lv_style_set_color_filter_dsc((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetColorFilterOpa(style *lib.LvStyleT, value lib.LvOpaT) SetStyle {
	C.lv_style_set_color_filter_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter SetStyle) SetAnim(style *lib.LvStyleT, value *lib.LvAnimT) SetStyle {
	C.lv_style_set_anim((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_anim_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetAnimTime(style *lib.LvStyleT, value uint32) SetStyle {
	C.lv_style_set_anim_time((*C.lv_style_t)(unsafe.Pointer(style)), C.uint(value))

	return setter
}
func (setter SetStyle) SetAnimSpeed(style *lib.LvStyleT, value uint32) SetStyle {
	C.lv_style_set_anim_speed((*C.lv_style_t)(unsafe.Pointer(style)), C.uint(value))

	return setter
}
func (setter SetStyle) SetTransition(style *lib.LvStyleT, value *lib.LvStyleTransitionDscT) SetStyle {
	C.lv_style_set_transition((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter SetStyle) SetBlendMode(style *lib.LvStyleT, value lib.LvBlendModeT) SetStyle {
	C.lv_style_set_blend_mode((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_blend_mode_t(value))

	return setter
}
func (setter SetStyle) SetLayout(style *lib.LvStyleT, value uint16) SetStyle {
	C.lv_style_set_layout((*C.lv_style_t)(unsafe.Pointer(style)), C.ushort(value))

	return setter
}
func (setter SetStyle) SetBaseDir(style *lib.LvStyleT, value lib.LvBaseDirT) SetStyle {
	C.lv_style_set_base_dir((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_base_dir_t(value))

	return setter
}
func (setter SetStyle) SetProp(style *lib.LvStyleT, prop lib.LvStylePropT, value lib.LvStyleValueT) SetStyle {
	C.lv_style_set_prop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_style_prop_t(prop), C.lv_style_value_t(value))

	return setter
}
func (setter SetStyle) SetPropMeta(style *lib.LvStyleT, prop lib.LvStylePropT, meta uint16) SetStyle {
	C.lv_style_set_prop_meta((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_style_prop_t(prop), C.ushort(meta))

	return setter
}
func (setter SetStyle) SetFlexFlow(style *lib.LvStyleT, value lib.LvFlexFlowT) SetStyle {
	C.lv_style_set_flex_flow((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_flow_t(value))

	return setter
}
func (setter SetStyle) SetFlexMainPlace(style *lib.LvStyleT, value lib.LvFlexAlignT) SetStyle {
	C.lv_style_set_flex_main_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter SetStyle) SetFlexCrossPlace(style *lib.LvStyleT, value lib.LvFlexAlignT) SetStyle {
	C.lv_style_set_flex_cross_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter SetStyle) SetFlexTrackPlace(style *lib.LvStyleT, value lib.LvFlexAlignT) SetStyle {
	C.lv_style_set_flex_track_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter SetStyle) SetFlexGrow(style *lib.LvStyleT, value uint8) SetStyle {
	C.lv_style_set_flex_grow((*C.lv_style_t)(unsafe.Pointer(style)), C.uint8_t(value))

	return setter
}
func (setter SetStyle) SetGridRowAlign(style *lib.LvStyleT, value lib.LvGridAlignT) SetStyle {
	C.lv_style_set_grid_row_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

	return setter
}
func (setter SetStyle) SetGridColumnAlign(style *lib.LvStyleT, value lib.LvGridAlignT) SetStyle {
	C.lv_style_set_grid_column_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

	return setter
}
func (setter SetStyle) SetGridCellColumnPos(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_column_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellColumnSpan(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_column_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellRowPos(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_row_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellRowSpan(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_row_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellXAlign(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_x_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter SetStyle) SetGridCellYAlign(style *lib.LvStyleT, value lib.LvCoordT) SetStyle {
	C.lv_style_set_grid_cell_y_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
