package set

import (
	"lvgl-go/src/lib"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Style set[CStyleT]
type lvCoordT = C.lv_coord_t

func CreateStyle(o CStyleT) *Style {
	_o := Style{
		CObj: o,
	}
	return &_o
}

func (setter *Style) Width(value lvCoordT) *Style {
	C.lv_style_set_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) MinWidth(value lvCoordT) *Style {
	C.lv_style_set_min_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) MaxWidth(value lvCoordT) *Style {
	C.lv_style_set_max_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) Height(value lvCoordT) *Style {
	C.lv_style_set_height(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) MinHeight(value lvCoordT) *Style {
	C.lv_style_set_min_height(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) MaxHeight(value lvCoordT) *Style {
	C.lv_style_set_max_height(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) X(value lvCoordT) *Style {
	C.lv_style_set_x(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) Y(value lvCoordT) *Style {
	C.lv_style_set_y(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) Align(value uint8) *Style {
	C.lv_style_set_align(setter.CObj, C.lv_align_t(value))

	return setter
}
func (setter *Style) TransformWidth(value lvCoordT) *Style {
	C.lv_style_set_transform_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TransformHeight(value lvCoordT) *Style {
	C.lv_style_set_transform_height(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TranslateX(value lvCoordT) *Style {
	C.lv_style_set_translate_x(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TranslateY(value lvCoordT) *Style {
	C.lv_style_set_translate_y(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TransformZoom(value lvCoordT) *Style {
	C.lv_style_set_transform_zoom(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TransformAngle(value lvCoordT) *Style {
	C.lv_style_set_transform_angle(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TransformPivotX(value lvCoordT) *Style {
	C.lv_style_set_transform_pivot_x(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TransformPivotY(value lvCoordT) *Style {
	C.lv_style_set_transform_pivot_y(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) Pad(value lvCoordT) *Style {
	C.lv_style_set_pad_all(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) PadTop(value lvCoordT) *Style {
	C.lv_style_set_pad_top(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) PadBottom(value lvCoordT) *Style {
	C.lv_style_set_pad_bottom(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) PadLeft(value lvCoordT) *Style {
	C.lv_style_set_pad_left(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) PadRight(value lvCoordT) *Style {
	C.lv_style_set_pad_right(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) PadRow(value lvCoordT) *Style {
	C.lv_style_set_pad_row(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) PadColumn(value lvCoordT) *Style {
	C.lv_style_set_pad_column(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) BgColor(color uint32) *Style {
	C.lv_style_set_bg_color(setter.CObj, C.lv_color_t(lib.ToColor(lib.HEX, color)))

	return setter
}
func (setter *Style) BgOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_bg_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) BgGradColor(color uint32) *Style {
	C.lv_style_set_bg_grad_color(setter.CObj, C.lv_color_t(lib.ToColor(lib.HEX, color)))

	return setter
}
func (setter *Style) BgGradDir(value int) *Style {
	C.lv_style_set_bg_grad_dir(setter.CObj, C.lv_grad_dir_t(value))

	return setter
}
func (setter *Style) BgMainStop(value lvCoordT) *Style {
	C.lv_style_set_bg_main_stop(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) BgGradStop(value lvCoordT) *Style {
	C.lv_style_set_bg_grad_stop(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) BgGrad(value *lib.LV_GRAD_DSC_T) *Style {
	C.lv_style_set_bg_grad(setter.CObj, (*C.lv_grad_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter *Style) BgDitherMode(value uint8) *Style {
	C.lv_style_set_bg_dither_mode(setter.CObj, C.lv_dither_mode_t(value))

	return setter
}
func (setter *Style) BgImgSrc(value any) *Style {
	C.lv_style_set_bg_img_src(setter.CObj, unsafe.Pointer(&value))

	return setter
}
func (setter *Style) BgImgOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_bg_img_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) BgImgRecolor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_bg_img_recolor(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) BgImgRecolorOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_bg_img_recolor_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) BgImgTiled(value bool) *Style {
	C.lv_style_set_bg_img_tiled(setter.CObj, C.bool(value))

	return setter
}
func (setter *Style) BorderColor(color uint32) *Style {
	C.lv_style_set_border_color(setter.CObj, C.lv_color_t(lib.ToColor(lib.HEX, color)))

	return setter
}
func (setter *Style) BorderOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_border_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) BorderWidth(value lvCoordT) *Style {
	C.lv_style_set_border_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) BorderSide(value uint8) *Style {
	C.lv_style_set_border_side(setter.CObj, C.lv_border_side_t(value))

	return setter
}
func (setter *Style) BorderPost(value bool) *Style {
	C.lv_style_set_border_post(setter.CObj, C.bool(value))

	return setter
}
func (setter *Style) OutlineWidth(value lvCoordT) *Style {
	C.lv_style_set_outline_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) OutlineColor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_outline_color(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) OutlineOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_outline_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) OutlinePad(value lvCoordT) *Style {
	C.lv_style_set_outline_pad(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ShadowWidth(value lvCoordT) *Style {
	C.lv_style_set_shadow_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ShadowOfsX(value lvCoordT) *Style {
	C.lv_style_set_shadow_ofs_x(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ShadowOfsY(value lvCoordT) *Style {
	C.lv_style_set_shadow_ofs_y(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ShadowSpread(value lvCoordT) *Style {
	C.lv_style_set_shadow_spread(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ShadowColor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_shadow_color(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) ShadowOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_shadow_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) ImgOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_img_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) ImgRecolor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_img_recolor(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) ImgRecolorOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_img_recolor_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) LineWidth(value lvCoordT) *Style {
	C.lv_style_set_line_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) LineDashWidth(value lvCoordT) *Style {
	C.lv_style_set_line_dash_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) LineDashGap(value lvCoordT) *Style {
	C.lv_style_set_line_dash_gap(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) LineRounded(value bool) *Style {
	C.lv_style_set_line_rounded(setter.CObj, C.bool(value))

	return setter
}
func (setter *Style) LineColor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_line_color(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) LineOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_line_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) ArcWidth(value lvCoordT) *Style {
	C.lv_style_set_arc_width(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ArcRounded(value bool) *Style {
	C.lv_style_set_arc_rounded(setter.CObj, C.bool(value))

	return setter
}
func (setter *Style) ArcColor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_arc_color(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) ArcOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_arc_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) ArcImgSrc(value any) *Style {
	C.lv_style_set_arc_img_src(setter.CObj, unsafe.Pointer(&value))

	return setter
}
func (setter *Style) TextColor(value lib.LV_COLOR_T) *Style {
	C.lv_style_set_text_color(setter.CObj, C.lv_color_t(value))

	return setter
}
func (setter *Style) TextOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_text_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) TextFont(value *lib.LV_FONT_T) *Style {
	C.lv_style_set_text_font(setter.CObj, (*C.lv_font_t)(unsafe.Pointer(value)))

	return setter
}
func (setter *Style) TextLetterSpace(value lvCoordT) *Style {
	C.lv_style_set_text_letter_space(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TextLineSpace(value lvCoordT) *Style {
	C.lv_style_set_text_line_space(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) TextDecor(value uint8) *Style {
	C.lv_style_set_text_decor(setter.CObj, C.lv_text_decor_t(value))

	return setter
}
func (setter *Style) TextAlign(value uint8) *Style {
	C.lv_style_set_text_align(setter.CObj, C.lv_text_align_t(value))

	return setter
}
func (setter *Style) Radius(value lvCoordT) *Style {
	C.lv_style_set_radius(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) ClipCorner(value bool) *Style {
	C.lv_style_set_clip_corner(setter.CObj, C.bool(value))

	return setter
}
func (setter *Style) Opa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) ColorFilterDsc(value *lib.LV_COLOR_FILTER_DSC_T) *Style {
	C.lv_style_set_color_filter_dsc(setter.CObj, (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter *Style) ColorFilterOpa(value lib.LV_OPA_T) *Style {
	C.lv_style_set_color_filter_opa(setter.CObj, C.lv_opa_t(value))

	return setter
}
func (setter *Style) Anim(value *lib.LV_ANIM_T) *Style {
	C.lv_style_set_anim(setter.CObj, (*C.lv_anim_t)(unsafe.Pointer(value)))

	return setter
}
func (setter *Style) AnimTime(value uint32) *Style {
	C.lv_style_set_anim_time(setter.CObj, C.uint(value))

	return setter
}
func (setter *Style) AnimSpeed(value uint32) *Style {
	C.lv_style_set_anim_speed(setter.CObj, C.uint(value))

	return setter
}
func (setter *Style) Transition(value *lib.LV_STYLE_TRANSITION_DSC_T) *Style {
	C.lv_style_set_transition(setter.CObj, (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter *Style) BlendMode(value uint8) *Style {
	C.lv_style_set_blend_mode(setter.CObj, C.lv_blend_mode_t(value))

	return setter
}
func (setter *Style) Layout(value uint16) *Style {
	C.lv_style_set_layout(setter.CObj, C.ushort(value))

	return setter
}
func (setter *Style) BaseDir(value uint8) *Style {
	C.lv_style_set_base_dir(setter.CObj, C.lv_base_dir_t(value))

	return setter
}
func (setter *Style) Prop(prop uint32, value lib.LV_STYLE_VALUE_T) *Style {
	C.lv_style_set_prop(setter.CObj, C.lv_style_prop_t(prop), C.lv_style_value_t(value))

	return setter
}
func (setter *Style) PropMeta(prop uint32, meta uint16) *Style {
	C.lv_style_set_prop_meta(setter.CObj, C.lv_style_prop_t(prop), C.ushort(meta))

	return setter
}
func (setter *Style) FlexFlow(value uint8) *Style {
	C.lv_style_set_flex_flow(setter.CObj, C.lv_flex_flow_t(value))

	return setter
}
func (setter *Style) FlexMainPlace(value uint8) *Style {
	C.lv_style_set_flex_main_place(setter.CObj, C.lv_flex_align_t(value))

	return setter
}
func (setter *Style) FlexCrossPlace(value uint8) *Style {
	C.lv_style_set_flex_cross_place(setter.CObj, C.lv_flex_align_t(value))

	return setter
}
func (setter *Style) FlexTrackPlace(value uint8) *Style {
	C.lv_style_set_flex_track_place(setter.CObj, C.lv_flex_align_t(value))

	return setter
}
func (setter *Style) FlexGrow(value uint8) *Style {
	C.lv_style_set_flex_grow(setter.CObj, C.uint8_t(value))

	return setter
}
func (setter *Style) GridRowAlign(value uint8) *Style {
	C.lv_style_set_grid_row_align(setter.CObj, C.lv_grid_align_t(value))

	return setter
}
func (setter *Style) GridColumnAlign(value uint8) *Style {
	C.lv_style_set_grid_column_align(setter.CObj, C.lv_grid_align_t(value))

	return setter
}
func (setter *Style) GridCellColumnPos(value lvCoordT) *Style {
	C.lv_style_set_grid_cell_column_pos(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) GridCellColumnSpan(value lvCoordT) *Style {
	C.lv_style_set_grid_cell_column_span(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) GridCellRowPos(value lvCoordT) *Style {
	C.lv_style_set_grid_cell_row_pos(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) GridCellRowSpan(value lvCoordT) *Style {
	C.lv_style_set_grid_cell_row_span(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) GridCellXAlign(value lvCoordT) *Style {
	C.lv_style_set_grid_cell_x_align(setter.CObj, lvCoordT(value))

	return setter
}
func (setter *Style) GridCellYAlign(value lvCoordT) *Style {
	C.lv_style_set_grid_cell_y_align(setter.CObj, lvCoordT(value))

	return setter
}
