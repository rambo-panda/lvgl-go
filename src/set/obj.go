package set

import (
	lib "lvgl-go/src/lib"
	types "lvgl-go/src/types"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Obj set[CObjT]

func CreateObj(o CObjT) *Obj {
	_o := Obj{
		CObj: o,
	}
	return &_o
}

func (setter *Obj) LocalStyleProp(prop types.LvStylePropT, value types.LvStyleValueT, selector uint32) *Obj {
	C.lv_obj_set_local_style_prop(setter.CObj, C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) LocalStylePropMeta(prop types.LvStylePropT, meta uint16, selector uint32) *Obj {
	C.lv_obj_set_local_style_prop_meta(setter.CObj, C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMinWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_min_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMaxWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_max_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleHeight(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMinHeight(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_min_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleMaxHeight(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_max_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleX(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleY(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleAlign(value types.LvAlignT, selector uint32) *Obj {
	C.lv_obj_set_style_align(setter.CObj, C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_transform_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformHeight(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_transform_height(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTranslateX(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_translate_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTranslateY(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_translate_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformZoom(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_transform_zoom(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformAngle(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_transform_angle(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformPivotX(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_transform_pivot_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTransformPivotY(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_transform_pivot_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadTop(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_pad_top(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadBottom(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_pad_bottom(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadLeft(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_pad_left(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadRight(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_pad_right(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadRow(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_pad_row(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StylePadColumn(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_pad_column(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGradColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGradDir(value types.LvGradDirT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad_dir(setter.CObj, C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgMainStop(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_main_stop(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGradStop(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad_stop(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgGrad(value *types.LvGradDscT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_grad(setter.CObj, (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgDitherMode(value types.LvDitherModeT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_dither_mode(setter.CObj, C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgSrc(value any, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_src(setter.CObj, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgRecolor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_recolor(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgRecolorOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_recolor_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBgImgTiled(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_bg_img_tiled(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_border_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_border_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_border_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderSide(value types.LvBorderSideT, selector uint32) *Obj {
	C.lv_obj_set_style_border_side(setter.CObj, C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBorderPost(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_border_post(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlineWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_outline_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlineColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_outline_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlineOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_outline_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOutlinePad(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_outline_pad(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowOfsX(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_ofs_x(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowOfsY(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_ofs_y(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowSpread(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_spread(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleShadowOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_shadow_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleImgOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_img_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleImgRecolor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_img_recolor(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleImgRecolorOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_img_recolor_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_line_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineDashWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_line_dash_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineDashGap(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_line_dash_gap(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineRounded(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_line_rounded(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_line_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLineOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_line_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcWidth(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_arc_width(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcRounded(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_arc_rounded(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_arc_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_arc_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleArcImgSrc(value any, selector uint32) *Obj {
	C.lv_obj_set_style_arc_img_src(setter.CObj, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextColor(value types.LvColorT, selector uint32) *Obj {
	C.lv_obj_set_style_text_color(setter.CObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_text_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextFont(value *types.LvFontT, selector uint32) *Obj {
	C.lv_obj_set_style_text_font(setter.CObj, (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextLetterSpace(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_text_letter_space(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextLineSpace(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_text_line_space(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextDecor(value types.LvTextDecorT, selector uint32) *Obj {
	C.lv_obj_set_style_text_decor(setter.CObj, C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleTextAlign(value types.LvTextAlignT, selector uint32) *Obj {
	C.lv_obj_set_style_text_align(setter.CObj, C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleRadius(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_radius(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleClipCorner(value bool, selector uint32) *Obj {
	C.lv_obj_set_style_clip_corner(setter.CObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleColorFilterDsc(value *types.LvColorFilterDscT, selector uint32) *Obj {
	C.lv_obj_set_style_color_filter_dsc(setter.CObj, (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleColorFilterOpa(value types.LvOpaT, selector uint32) *Obj {
	C.lv_obj_set_style_color_filter_opa(setter.CObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleAnim(value *types.LvAnimT, selector uint32) *Obj {
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
func (setter *Obj) StyleTransition(value *types.LvStyleTransitionDscT, selector uint32) *Obj {
	C.lv_obj_set_style_transition(setter.CObj, (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBlendMode(value types.LvBlendModeT, selector uint32) *Obj {
	C.lv_obj_set_style_blend_mode(setter.CObj, C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleLayout(value uint16, selector uint32) *Obj {
	C.lv_obj_set_style_layout(setter.CObj, C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleBaseDir(value types.LvBaseDirT, selector uint32) *Obj {
	C.lv_obj_set_style_base_dir(setter.CObj, C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) Pos(x types.LvCoordT, y types.LvCoordT) *Obj {
	C.lv_obj_set_pos(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter *Obj) X(x types.LvCoordT) *Obj {
	C.lv_obj_set_x(setter.CObj, C.lv_coord_t(x))

	return setter
}
func (setter *Obj) Y(y types.LvCoordT) *Obj {
	C.lv_obj_set_y(setter.CObj, C.lv_coord_t(y))

	return setter
}
func (setter *Obj) Size(w types.LvCoordT, h types.LvCoordT) *Obj {
	C.lv_obj_set_size(setter.CObj, C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter *Obj) Width(w types.LvCoordT) *Obj {
	C.lv_obj_set_width(setter.CObj, C.lv_coord_t(w))

	return setter
}
func (setter *Obj) Height(h types.LvCoordT) *Obj {
	C.lv_obj_set_height(setter.CObj, C.lv_coord_t(h))

	return setter
}
func (setter *Obj) ContentWidth(w types.LvCoordT) *Obj {
	C.lv_obj_set_content_width(setter.CObj, C.lv_coord_t(w))

	return setter
}
func (setter *Obj) ContentHeight(h types.LvCoordT) *Obj {
	C.lv_obj_set_content_height(setter.CObj, C.lv_coord_t(h))

	return setter
}
func (setter *Obj) Layout(layout uint32) *Obj {
	C.lv_obj_set_layout(setter.CObj, C.uint(layout))

	return setter
}
func (setter *Obj) Align(align types.LvAlignT) *Obj {
	C.lv_obj_set_align(setter.CObj, C.lv_align_t(align))

	return setter
}
func (setter *Obj) Center() *Obj {
	setter.Align(types.LV_ALIGN_CENTER).Pos(0, 0)
	return setter
}
func (setter *Obj) ExtClickArea(size types.LvCoordT) *Obj {
	C.lv_obj_set_ext_click_area(setter.CObj, C.lv_coord_t(size))

	return setter
}
func (setter *Obj) ScrollbarMode(mode types.LvScrollbarModeT) *Obj {
	C.lv_obj_set_scrollbar_mode(setter.CObj, C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter *Obj) ScrollDir(dir types.LvDirT) *Obj {
	C.lv_obj_set_scroll_dir(setter.CObj, C.lv_dir_t(dir))

	return setter
}
func (setter *Obj) ScrollSnapX(align types.LvScrollSnapT) *Obj {
	C.lv_obj_set_scroll_snap_x(setter.CObj, C.lv_scroll_snap_t(align))

	return setter
}
func (setter *Obj) ScrollSnapY(align types.LvScrollSnapT) *Obj {
	C.lv_obj_set_scroll_snap_y(setter.CObj, C.lv_scroll_snap_t(align))

	return setter
}
func (setter *Obj) Parent(parent *types.LvObjT) *Obj {
	C.lv_obj_set_parent(setter.CObj, (*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter *Obj) FlexFlow(flow types.LvFlexFlowT) *Obj {
	C.lv_obj_set_flex_flow(setter.CObj, C.lv_flex_flow_t(flow))

	return setter
}
func (setter *Obj) FlexAlign(main_place types.LvFlexAlignT, cross_place types.LvFlexAlignT, track_cross_place types.LvFlexAlignT) *Obj {
	C.lv_obj_set_flex_align(setter.CObj, C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

	return setter
}
func (setter *Obj) FlexGrow(grow uint8) *Obj {
	C.lv_obj_set_flex_grow(setter.CObj, C.uint8_t(grow))

	return setter
}
func (setter *Obj) StyleFlexFlow(value types.LvFlexFlowT, selector uint32) *Obj {
	C.lv_obj_set_style_flex_flow(setter.CObj, C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexMainPlace(value types.LvFlexAlignT, selector uint32) *Obj {
	C.lv_obj_set_style_flex_main_place(setter.CObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexCrossPlace(value types.LvFlexAlignT, selector uint32) *Obj {
	C.lv_obj_set_style_flex_cross_place(setter.CObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleFlexTrackPlace(value types.LvFlexAlignT, selector uint32) *Obj {
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
func (setter *Obj) StyleGridCellColumnPos(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_column_pos(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellColumnSpan(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_column_span(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellRowPos(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_row_pos(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellRowSpan(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_row_span(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellXAlign(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_x_align(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) StyleGridCellYAlign(value types.LvCoordT, selector uint32) *Obj {
	C.lv_obj_set_style_grid_cell_y_align(setter.CObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter *Obj) Tile(anim_en types.LvAnimEnableT) *Obj {
	C.lv_obj_set_tile(setter.CObj, setter.CObj, C.lv_anim_enable_t(anim_en))

	return setter
}
func (setter *Obj) TileId(col_id uint32, row_id uint32, anim_en types.LvAnimEnableT) *Obj {
	C.lv_obj_set_tile_id(setter.CObj, C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

	return setter
}

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
