package set

import (
	"unsafe"

	lib "gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"

type Obj set[CObjT]

func CreateObj(o CObjT) *Obj {
	_o := Obj{
		CObj: o,
	}
	return &_o
}

func (setter *Obj) LocalStyleProp(prop uint32, value lib.LV_STYLE_VALUE_T, selector uint32) *Obj {
	C.lv_obj_set_local_style_prop(setter.CObj, C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) LocalStylePropMeta(prop uint32, meta uint16, selector uint32) *Obj {
	C.lv_obj_set_local_style_prop_meta(setter.CObj, C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMinWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_min_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMaxWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_max_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleHeight(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMinHeight(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_min_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMaxHeight(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_max_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleX(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleY(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleAlign(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_align(setter.CObj, C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_transform_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformHeight(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_transform_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTranslateX(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_translate_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTranslateY(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_translate_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformZoom(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_transform_zoom(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformAngle(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_transform_angle(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformPivotX(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_transform_pivot_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformPivotY(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_transform_pivot_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadTop(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_pad_top(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadBottom(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_pad_bottom(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadLeft(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_pad_left(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadRight(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_pad_right(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadRow(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_pad_row(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadColumn(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_pad_column(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGradColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGradDir(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad_dir(setter.CObj, C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgMainStop(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_bg_main_stop(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGradStop(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad_stop(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGrad(value *lib.LV_GRAD_DSC_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad(setter.CObj, (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgDitherMode(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_bg_dither_mode(setter.CObj, C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgSrc(value any, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_src(setter.CObj, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgRecolor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_recolor(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgRecolorOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_recolor_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgTiled(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_tiled(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_border_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_border_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_border_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderSide(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_border_side(setter.CObj, C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderPost(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_border_post(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlineWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_outline_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlineColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_outline_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlineOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_outline_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlinePad(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_outline_pad(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowOfsX(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_ofs_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowOfsY(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_ofs_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowSpread(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_spread(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowColor(value lib.LV_COLOR16_T, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleImgOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_img_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleImgRecolor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_img_recolor(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleImgRecolorOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_img_recolor_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_line_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineDashWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_line_dash_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineDashGap(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_line_dash_gap(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineRounded(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_line_rounded(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_line_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_line_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcWidth(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_arc_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcRounded(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_arc_rounded(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_arc_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_arc_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcImgSrc(value any, selector uint32) *Obj {
	C.lv_obj_set_style_arc_img_src(setter.CObj, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextColor(value lib.LV_COLOR_T, selector uint32) *Obj {
	C.lv_obj_set_style_text_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_text_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextFont(value *lib.LV_FONT_T, selector uint32) *Obj {
	C.lv_obj_set_style_text_font(setter.CObj, (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextLetterSpace(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_text_letter_space(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextLineSpace(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_text_line_space(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextDecor(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_text_decor(setter.CObj, C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextAlign(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_text_align(setter.CObj, C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleRadius(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_radius(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleClipCorner(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_clip_corner(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleColorFilterDsc(value *lib.LV_COLOR_FILTER_DSC_T, selector uint32) *Obj {
	C.lv_obj_set_style_color_filter_dsc(setter.CObj, (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleColorFilterOpa(value lib.LV_OPA_T, selector uint32) *Obj {
	C.lv_obj_set_style_color_filter_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleAnim(value *lib.LV_ANIM_T, selector uint32) *Obj {
	C.lv_obj_set_style_anim(setter.CObj, (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleAnimTime(value uint32, selector uint32) *Obj {
	C.lv_obj_set_style_anim_time(setter.CObj, C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleAnimSpeed(value uint32, selector uint32) *Obj {
	C.lv_obj_set_style_anim_speed(setter.CObj, C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransition(value *lib.LV_STYLE_TRANSITION_DSC_T, selector uint32) *Obj {
	C.lv_obj_set_style_transition(setter.CObj, (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBlendMode(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_blend_mode(setter.CObj, C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLayout(value uint16, selector uint32) *Obj {
	C.lv_obj_set_style_layout(setter.CObj, C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBaseDir(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_base_dir(setter.CObj, C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) Pos(x int16, y int16) *Obj {
	C.lv_obj_set_pos(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter *Obj) X(x int16) *Obj {
	C.lv_obj_set_x(setter.CObj, C.lv_coord_t(x))

	return setter
}
func (setter *Obj) Y(y int16) *Obj {
	C.lv_obj_set_y(setter.CObj, C.lv_coord_t(y))

	return setter
}
func (setter *Obj) Size(w int16, h int16) *Obj {
	C.lv_obj_set_size(setter.CObj, C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter *Obj) Width(w int16) *Obj {
	C.lv_obj_set_width(setter.CObj, C.lv_coord_t(w))

	return setter
}
func (setter *Obj) Height(h int16) *Obj {
	C.lv_obj_set_height(setter.CObj, C.lv_coord_t(h))

	return setter
}
func (setter *Obj) ContentWidth(w int16) *Obj {
	C.lv_obj_set_content_width(setter.CObj, C.lv_coord_t(w))

	return setter
}
func (setter *Obj) ContentHeight(h int16) *Obj {
	C.lv_obj_set_content_height(setter.CObj, C.lv_coord_t(h))

	return setter
}
func (setter *Obj) Layout(layout uint32) *Obj {
	C.lv_obj_set_layout(setter.CObj, C.uint(layout))

	return setter
}
func (setter *Obj) Align(align uint8) *Obj {
	C.lv_obj_set_align(setter.CObj, C.lv_align_t(align))

	return setter
}
func (setter *Obj) ToAlign(align uint8, x int16, y int16) *Obj {
	setter.Align(align)
	setter.Pos(x, y)

	return setter
}
func (setter *Obj) AlignTo(o lib.CreateI, align uint8, x int16, y int16) *Obj {
	C.lv_obj_align_to(setter.CObj, (CObjT)(o.GetObj()), C.lv_align_t(align), C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter *Obj) Center() *Obj {
	setter.Align(lib.LV_ALIGN_CENTER).Pos(0, 0)
	return setter
}
func (setter *Obj) ExtClickArea(size int16) *Obj {
	C.lv_obj_set_ext_click_area(setter.CObj, C.lv_coord_t(size))

	return setter
}
func (setter *Obj) ScrollbarMode(mode uint8) *Obj {
	C.lv_obj_set_scrollbar_mode(setter.CObj, C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter *Obj) ScrollDir(dir uint8) *Obj {
	C.lv_obj_set_scroll_dir(setter.CObj, C.lv_dir_t(dir))

	return setter
}
func (setter *Obj) ScrollSnapX(align uint8) *Obj {
	C.lv_obj_set_scroll_snap_x(setter.CObj, C.lv_scroll_snap_t(align))

	return setter
}
func (setter *Obj) ScrollSnapY(align uint8) *Obj {
	C.lv_obj_set_scroll_snap_y(setter.CObj, C.lv_scroll_snap_t(align))

	return setter
}
func (setter *Obj) Parent(parent lib.CreateI) *Obj {
	C.lv_obj_set_parent(setter.CObj, (CObjT)(getParentObj(parent)))

	return setter
}
func (setter *Obj) FlexFlow(flow uint8) *Obj {
	C.lv_obj_set_flex_flow(setter.CObj, C.lv_flex_flow_t(flow))

	return setter
}
func (setter *Obj) FlexAlign(main_place uint8, cross_place uint8, track_cross_place uint8) *Obj {
	C.lv_obj_set_flex_align(setter.CObj, C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

	return setter
}
func (setter *Obj) FlexGrow(grow uint8) *Obj {
	C.lv_obj_set_flex_grow(setter.CObj, C.uint8_t(grow))

	return setter
}
func (setter *Obj) StyleFlexFlow(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_flex_flow(setter.CObj, C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexMainPlace(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_flex_main_place(setter.CObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexCrossPlace(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_flex_cross_place(setter.CObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexTrackPlace(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_flex_track_place(setter.CObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexGrow(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_flex_grow(setter.CObj, C.uint8_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) GridAlign(column_align uint8, row_align uint8) *Obj {
	C.lv_obj_set_grid_align(setter.CObj, C.lv_grid_align_t(column_align), C.lv_grid_align_t(row_align))

	return setter
}
func (setter *Obj) GridCell(column_align uint8, col_pos uint8, col_span uint8, row_align uint8, row_pos uint8, row_span uint8) *Obj {
	C.lv_obj_set_grid_cell(setter.CObj, C.lv_grid_align_t(column_align), C.uint8_t(col_pos), C.uint8_t(col_span), C.lv_grid_align_t(row_align), C.uint8_t(row_pos), C.uint8_t(row_span))

	return setter
}
func (setter *Obj) StyleGridRowAlign(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_grid_row_align(setter.CObj, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridColumnAlign(value uint8, selector uint32) *Obj {
	C.lv_obj_set_style_grid_column_align(setter.CObj, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellColumnPos(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_column_pos(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellColumnSpan(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_column_span(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellRowPos(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_row_pos(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellRowSpan(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_row_span(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellXAlign(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_x_align(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellYAlign(value int16, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_y_align(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) Tile(anim_en uint8) *Obj {
	C.lv_obj_set_tile(setter.CObj, setter.CObj, C.lv_anim_enable_t(anim_en))

	return setter
}
func (setter *Obj) TileId(col_id uint32, row_id uint32, anim_en uint8) *Obj {
	C.lv_obj_set_tile_id(setter.CObj, C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

	return setter
}

// @todo 这玩意会连你之前设置的size都给清楚掉
func (setter *Obj) RemoveStyleAll() *Obj {
	C.lv_obj_remove_style_all(setter.CObj)

	return setter
}

func (setter *Obj) Style(style lib.CreateI, styleSelect int) *Obj {
	var oStyle *C.lv_style_t

	oStyle = (*C.lv_style_t)(getParentObj(style))

	C.lv_obj_add_style(setter.CObj, oStyle, C.lv_style_selector_t(styleSelect))

	return setter
}
