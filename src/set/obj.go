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

type SetObj set

func (setter SetObj) SetLocalStyleProp(prop lib.LvStylePropT, value lib.LvStyleValueT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_local_style_prop(setter.cObj, C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetLocalStylePropMeta(prop lib.LvStylePropT, meta uint16, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_local_style_prop_meta(setter.cObj, C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMinWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_min_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMaxWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_max_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleHeight(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_height(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMinHeight(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_min_height(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMaxHeight(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_max_height(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleX(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_x(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleY(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_y(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAlign(value lib.LvAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_align(setter.cObj, C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformHeight(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_height(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTranslateX(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_translate_x(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTranslateY(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_translate_y(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformZoom(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_zoom(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformAngle(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_angle(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformPivotX(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_pivot_x(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformPivotY(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_pivot_y(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadTop(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_top(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadBottom(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_bottom(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadLeft(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_left(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadRight(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_right(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadRow(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_row(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadColumn(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_column(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradDir(value lib.LvGradDirT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_dir(setter.cObj, C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgMainStop(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_main_stop(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradStop(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_stop(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGrad(value *lib.LvGradDscT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad(setter.cObj, (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgDitherMode(value lib.LvDitherModeT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_dither_mode(setter.cObj, C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgSrc(value any, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_src(setter.cObj, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgRecolor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_recolor(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgRecolorOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_recolor_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgTiled(value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_tiled(setter.cObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderSide(value lib.LvBorderSideT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_side(setter.cObj, C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderPost(value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_post(setter.cObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlinePad(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_pad(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOfsX(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_ofs_x(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOfsY(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_ofs_y(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowSpread(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_spread(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgRecolor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_recolor(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgRecolorOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_recolor_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineDashWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_dash_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineDashGap(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_dash_gap(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineRounded(value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_rounded(setter.cObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcWidth(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_width(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcRounded(value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_rounded(setter.cObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcImgSrc(value any, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_img_src(setter.cObj, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextColor(value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_color(setter.cObj, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextFont(value *lib.LvFontT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_font(setter.cObj, (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextLetterSpace(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_letter_space(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextLineSpace(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_line_space(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextDecor(value lib.LvTextDecorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_decor(setter.cObj, C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextAlign(value lib.LvTextAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_align(setter.cObj, C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleRadius(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_radius(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleClipCorner(value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_clip_corner(setter.cObj, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleColorFilterDsc(value *lib.LvColorFilterDscT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_color_filter_dsc(setter.cObj, (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleColorFilterOpa(value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_color_filter_opa(setter.cObj, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnim(value *lib.LvAnimT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim(setter.cObj, (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnimTime(value uint32, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim_time(setter.cObj, C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnimSpeed(value uint32, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim_speed(setter.cObj, C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransition(value *lib.LvStyleTransitionDscT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transition(setter.cObj, (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBlendMode(value lib.LvBlendModeT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_blend_mode(setter.cObj, C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLayout(value uint16, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_layout(setter.cObj, C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBaseDir(value lib.LvBaseDirT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_base_dir(setter.cObj, C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetPos(x lib.LvCoordT, y lib.LvCoordT) SetObj {
	C.lv_obj_set_pos(setter.cObj, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter SetObj) SetX(x lib.LvCoordT) SetObj {
	C.lv_obj_set_x(setter.cObj, C.lv_coord_t(x))

	return setter
}
func (setter SetObj) SetY(y lib.LvCoordT) SetObj {
	C.lv_obj_set_y(setter.cObj, C.lv_coord_t(y))

	return setter
}
func (setter SetObj) SetSize(w lib.LvCoordT, h lib.LvCoordT) SetObj {
	C.lv_obj_set_size(setter.cObj, C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetWidth(w lib.LvCoordT) SetObj {
	C.lv_obj_set_width(setter.cObj, C.lv_coord_t(w))

	return setter
}
func (setter SetObj) SetHeight(h lib.LvCoordT) SetObj {
	C.lv_obj_set_height(setter.cObj, C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetContentWidth(w lib.LvCoordT) SetObj {
	C.lv_obj_set_content_width(setter.cObj, C.lv_coord_t(w))

	return setter
}
func (setter SetObj) SetContentHeight(h lib.LvCoordT) SetObj {
	C.lv_obj_set_content_height(setter.cObj, C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetLayout(layout uint32) SetObj {
	C.lv_obj_set_layout(setter.cObj, C.uint(layout))

	return setter
}
func (setter SetObj) SetAlign(align lib.LvAlignT) SetObj {
	C.lv_obj_set_align(setter.cObj, C.lv_align_t(align))

	return setter
}
func (setter SetObj) SetExtClickArea(size lib.LvCoordT) SetObj {
	C.lv_obj_set_ext_click_area(setter.cObj, C.lv_coord_t(size))

	return setter
}
func (setter SetObj) SetScrollbarMode(mode lib.LvScrollbarModeT) SetObj {
	C.lv_obj_set_scrollbar_mode(setter.cObj, C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter SetObj) SetScrollDir(dir lib.LvDirT) SetObj {
	C.lv_obj_set_scroll_dir(setter.cObj, C.lv_dir_t(dir))

	return setter
}
func (setter SetObj) SetScrollSnapX(align lib.LvScrollSnapT) SetObj {
	C.lv_obj_set_scroll_snap_x(setter.cObj, C.lv_scroll_snap_t(align))

	return setter
}
func (setter SetObj) SetScrollSnapY(align lib.LvScrollSnapT) SetObj {
	C.lv_obj_set_scroll_snap_y(setter.cObj, C.lv_scroll_snap_t(align))

	return setter
}
func (setter SetObj) SetParent() SetObj {
	C.lv_obj_set_parent(setter.cObj, setter.cObj)

	return setter
}
func (setter SetObj) SetFlexFlow(flow lib.LvFlexFlowT) SetObj {
	C.lv_obj_set_flex_flow(setter.cObj, C.lv_flex_flow_t(flow))

	return setter
}
func (setter SetObj) SetFlexAlign(main_place lib.LvFlexAlignT, cross_place lib.LvFlexAlignT, track_cross_place lib.LvFlexAlignT) SetObj {
	C.lv_obj_set_flex_align(setter.cObj, C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

	return setter
}
func (setter SetObj) SetFlexGrow(grow uint8) SetObj {
	C.lv_obj_set_flex_grow(setter.cObj, C.uint8_t(grow))

	return setter
}
func (setter SetObj) SetStyleFlexFlow(value lib.LvFlexFlowT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_flow(setter.cObj, C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexMainPlace(value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_main_place(setter.cObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexCrossPlace(value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_cross_place(setter.cObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexTrackPlace(value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_track_place(setter.cObj, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexGrow(value uint8, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_grow(setter.cObj, C.uint8_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetGridAlign(column_align lib.LvGridAlignT, row_align lib.LvGridAlignT) SetObj {
	C.lv_obj_set_grid_align(setter.cObj, C.lv_grid_align_t(column_align), C.lv_grid_align_t(row_align))

	return setter
}
func (setter SetObj) SetGridCell(column_align lib.LvGridAlignT, col_pos uint8, col_span uint8, row_align lib.LvGridAlignT, row_pos uint8, row_span uint8) SetObj {
	C.lv_obj_set_grid_cell(setter.cObj, C.lv_grid_align_t(column_align), C.uint8_t(col_pos), C.uint8_t(col_span), C.lv_grid_align_t(row_align), C.uint8_t(row_pos), C.uint8_t(row_span))

	return setter
}
func (setter SetObj) SetStyleGridRowAlign(value lib.LvGridAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_row_align(setter.cObj, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridColumnAlign(value lib.LvGridAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_column_align(setter.cObj, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellColumnPos(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_column_pos(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellColumnSpan(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_column_span(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellRowPos(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_row_pos(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellRowSpan(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_row_span(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellXAlign(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_x_align(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellYAlign(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_y_align(setter.cObj, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetTile(anim_en lib.LvAnimEnableT) SetObj {
	C.lv_obj_set_tile(setter.cObj, setter.cObj, C.lv_anim_enable_t(anim_en))

	return setter
}
func (setter SetObj) SetTileId(col_id uint32, row_id uint32, anim_en lib.LvAnimEnableT) SetObj {
	C.lv_obj_set_tile_id(setter.cObj, C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

	return setter
}
