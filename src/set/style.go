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

type Style set

func CreateStyle(o CObjT) Style {
	return Style{
		CObj: o,
	}
}

func (setter Style) Width(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) MinWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_min_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) MaxWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_max_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) Height(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) MinHeight(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_min_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) MaxHeight(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_max_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) X(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) Y(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) Align(style *types.LvStyleT, value types.LvAlignT) Style {
	C.lv_style_set_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_align_t(value))

	return setter
}
func (setter Style) TransformWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_transform_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TransformHeight(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_transform_height((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TranslateX(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_translate_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TranslateY(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_translate_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TransformZoom(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_transform_zoom((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TransformAngle(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_transform_angle((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TransformPivotX(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_transform_pivot_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TransformPivotY(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_transform_pivot_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) PadTop(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_pad_top((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) PadBottom(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_pad_bottom((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) PadLeft(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_pad_left((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) PadRight(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_pad_right((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) PadRow(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_pad_row((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) PadColumn(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_pad_column((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) BgColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_bg_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) BgOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_bg_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) BgGradColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_bg_grad_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) BgGradDir(style *types.LvStyleT, value types.LvGradDirT) Style {
	C.lv_style_set_bg_grad_dir((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grad_dir_t(value))

	return setter
}
func (setter Style) BgMainStop(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_bg_main_stop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) BgGradStop(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_bg_grad_stop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) BgGrad(style *types.LvStyleT, value *types.LvGradDscT) Style {
	C.lv_style_set_bg_grad((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter Style) BgDitherMode(style *types.LvStyleT, value types.LvDitherModeT) Style {
	C.lv_style_set_bg_dither_mode((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_dither_mode_t(value))

	return setter
}
func (setter Style) BgImgSrc(style *types.LvStyleT, value any) Style {
	C.lv_style_set_bg_img_src((*C.lv_style_t)(unsafe.Pointer(style)), unsafe.Pointer(&value))

	return setter
}
func (setter Style) BgImgOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_bg_img_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) BgImgRecolor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_bg_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) BgImgRecolorOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_bg_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) BgImgTiled(style *types.LvStyleT, value bool) Style {
	C.lv_style_set_bg_img_tiled((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter Style) BorderColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_border_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) BorderOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_border_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) BorderWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_border_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) BorderSide(style *types.LvStyleT, value types.LvBorderSideT) Style {
	C.lv_style_set_border_side((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_border_side_t(value))

	return setter
}
func (setter Style) BorderPost(style *types.LvStyleT, value bool) Style {
	C.lv_style_set_border_post((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter Style) OutlineWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_outline_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) OutlineColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_outline_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) OutlineOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_outline_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) OutlinePad(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_outline_pad((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ShadowWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_shadow_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ShadowOfsX(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_shadow_ofs_x((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ShadowOfsY(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_shadow_ofs_y((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ShadowSpread(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_shadow_spread((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ShadowColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_shadow_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) ShadowOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_shadow_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) ImgOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_img_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) ImgRecolor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) ImgRecolorOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) LineWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_line_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) LineDashWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_line_dash_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) LineDashGap(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_line_dash_gap((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) LineRounded(style *types.LvStyleT, value bool) Style {
	C.lv_style_set_line_rounded((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter Style) LineColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_line_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) LineOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_line_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) ArcWidth(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_arc_width((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ArcRounded(style *types.LvStyleT, value bool) Style {
	C.lv_style_set_arc_rounded((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter Style) ArcColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_arc_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) ArcOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_arc_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) ArcImgSrc(style *types.LvStyleT, value any) Style {
	C.lv_style_set_arc_img_src((*C.lv_style_t)(unsafe.Pointer(style)), unsafe.Pointer(&value))

	return setter
}
func (setter Style) TextColor(style *types.LvStyleT, value types.LvColorT) Style {
	C.lv_style_set_text_color((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_color_t(value))

	return setter
}
func (setter Style) TextOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_text_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) TextFont(style *types.LvStyleT, value *types.LvFontT) Style {
	C.lv_style_set_text_font((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_font_t)(unsafe.Pointer(value)))

	return setter
}
func (setter Style) TextLetterSpace(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_text_letter_space((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TextLineSpace(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_text_line_space((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) TextDecor(style *types.LvStyleT, value types.LvTextDecorT) Style {
	C.lv_style_set_text_decor((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_text_decor_t(value))

	return setter
}
func (setter Style) TextAlign(style *types.LvStyleT, value types.LvTextAlignT) Style {
	C.lv_style_set_text_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_text_align_t(value))

	return setter
}
func (setter Style) Radius(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_radius((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) ClipCorner(style *types.LvStyleT, value bool) Style {
	C.lv_style_set_clip_corner((*C.lv_style_t)(unsafe.Pointer(style)), C.bool(value))

	return setter
}
func (setter Style) Opa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) ColorFilterDsc(style *types.LvStyleT, value *types.LvColorFilterDscT) Style {
	C.lv_style_set_color_filter_dsc((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter Style) ColorFilterOpa(style *types.LvStyleT, value types.LvOpaT) Style {
	C.lv_style_set_color_filter_opa((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_opa_t(value))

	return setter
}
func (setter Style) Anim(style *types.LvStyleT, value *types.LvAnimT) Style {
	C.lv_style_set_anim((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_anim_t)(unsafe.Pointer(value)))

	return setter
}
func (setter Style) AnimTime(style *types.LvStyleT, value uint32) Style {
	C.lv_style_set_anim_time((*C.lv_style_t)(unsafe.Pointer(style)), C.uint(value))

	return setter
}
func (setter Style) AnimSpeed(style *types.LvStyleT, value uint32) Style {
	C.lv_style_set_anim_speed((*C.lv_style_t)(unsafe.Pointer(style)), C.uint(value))

	return setter
}
func (setter Style) Transition(style *types.LvStyleT, value *types.LvStyleTransitionDscT) Style {
	C.lv_style_set_transition((*C.lv_style_t)(unsafe.Pointer(style)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)))

	return setter
}
func (setter Style) BlendMode(style *types.LvStyleT, value types.LvBlendModeT) Style {
	C.lv_style_set_blend_mode((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_blend_mode_t(value))

	return setter
}
func (setter Style) Layout(style *types.LvStyleT, value uint16) Style {
	C.lv_style_set_layout((*C.lv_style_t)(unsafe.Pointer(style)), C.ushort(value))

	return setter
}
func (setter Style) BaseDir(style *types.LvStyleT, value types.LvBaseDirT) Style {
	C.lv_style_set_base_dir((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_base_dir_t(value))

	return setter
}
func (setter Style) Prop(style *types.LvStyleT, prop types.LvStylePropT, value types.LvStyleValueT) Style {
	C.lv_style_set_prop((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_style_prop_t(prop), C.lv_style_value_t(value))

	return setter
}
func (setter Style) PropMeta(style *types.LvStyleT, prop types.LvStylePropT, meta uint16) Style {
	C.lv_style_set_prop_meta((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_style_prop_t(prop), C.ushort(meta))

	return setter
}
func (setter Style) FlexFlow(style *types.LvStyleT, value types.LvFlexFlowT) Style {
	C.lv_style_set_flex_flow((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_flow_t(value))

	return setter
}
func (setter Style) FlexMainPlace(style *types.LvStyleT, value types.LvFlexAlignT) Style {
	C.lv_style_set_flex_main_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter Style) FlexCrossPlace(style *types.LvStyleT, value types.LvFlexAlignT) Style {
	C.lv_style_set_flex_cross_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter Style) FlexTrackPlace(style *types.LvStyleT, value types.LvFlexAlignT) Style {
	C.lv_style_set_flex_track_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

	return setter
}
func (setter Style) FlexGrow(style *types.LvStyleT, value uint8) Style {
	C.lv_style_set_flex_grow((*C.lv_style_t)(unsafe.Pointer(style)), C.uint8_t(value))

	return setter
}
func (setter Style) GridRowAlign(style *types.LvStyleT, value types.LvGridAlignT) Style {
	C.lv_style_set_grid_row_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

	return setter
}
func (setter Style) GridColumnAlign(style *types.LvStyleT, value types.LvGridAlignT) Style {
	C.lv_style_set_grid_column_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

	return setter
}
func (setter Style) GridCellColumnPos(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_grid_cell_column_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) GridCellColumnSpan(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_grid_cell_column_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) GridCellRowPos(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_grid_cell_row_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) GridCellRowSpan(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_grid_cell_row_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) GridCellXAlign(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_grid_cell_x_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
func (setter Style) GridCellYAlign(style *types.LvStyleT, value types.LvCoordT) Style {
	C.lv_style_set_grid_cell_y_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

	return setter
}
