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

type Obj set

func CreateObj(o CObjT) Obj {
	return Obj{
		CStructLvObjT: o,
	}
}

func (setter Obj) LocalStyleProp(prop types.LvStylePropT, value types.LvStyleValueT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_local_style_prop(setter.CStructLvObjT, C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) LocalStylePropMeta(prop types.LvStylePropT, meta uint16, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_local_style_prop_meta(setter.CStructLvObjT, C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMinWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_min_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMaxWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_max_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleHeight(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_height(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMinHeight(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_min_height(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMaxHeight(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_max_height(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleX(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_x(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleY(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_y(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAlign(value types.LvAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_align(setter.CStructLvObjT, C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformHeight(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_height(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTranslateX(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_translate_x(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTranslateY(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_translate_y(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformZoom(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_zoom(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformAngle(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_angle(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformPivotX(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_pivot_x(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformPivotY(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_pivot_y(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadTop(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_top(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadBottom(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_bottom(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadLeft(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_left(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadRight(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_right(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadRow(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_row(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadColumn(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_column(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGradColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGradDir(value types.LvGradDirT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad_dir(setter.CStructLvObjT, C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgMainStop(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_main_stop(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGradStop(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad_stop(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGrad(value *types.LvGradDscT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad(setter.CStructLvObjT, (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgDitherMode(value types.LvDitherModeT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_dither_mode(setter.CStructLvObjT, C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgSrc(value any, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_src(setter.CStructLvObjT, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgRecolor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_recolor(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgRecolorOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_recolor_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgTiled(value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_tiled(setter.CStructLvObjT, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderSide(value types.LvBorderSideT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_side(setter.CStructLvObjT, C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderPost(value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_post(setter.CStructLvObjT, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlineWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlineColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlineOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlinePad(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_pad(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowOfsX(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_ofs_x(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowOfsY(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_ofs_y(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowSpread(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_spread(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleImgOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_img_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleImgRecolor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_img_recolor(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleImgRecolorOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_img_recolor_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineDashWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_dash_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineDashGap(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_dash_gap(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineRounded(value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_rounded(setter.CStructLvObjT, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcWidth(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_width(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcRounded(value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_rounded(setter.CStructLvObjT, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcImgSrc(value any, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_img_src(setter.CStructLvObjT, unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextColor(value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_color(setter.CStructLvObjT, C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextFont(value *types.LvFontT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_font(setter.CStructLvObjT, (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextLetterSpace(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_letter_space(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextLineSpace(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_line_space(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextDecor(value types.LvTextDecorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_decor(setter.CStructLvObjT, C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextAlign(value types.LvTextAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_align(setter.CStructLvObjT, C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleRadius(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_radius(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleClipCorner(value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_clip_corner(setter.CStructLvObjT, C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleColorFilterDsc(value *types.LvColorFilterDscT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_color_filter_dsc(setter.CStructLvObjT, (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleColorFilterOpa(value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_color_filter_opa(setter.CStructLvObjT, C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAnim(value *types.LvAnimT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_anim(setter.CStructLvObjT, (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAnimTime(value uint32, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_anim_time(setter.CStructLvObjT, C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAnimSpeed(value uint32, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_anim_speed(setter.CStructLvObjT, C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransition(value *types.LvStyleTransitionDscT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transition(setter.CStructLvObjT, (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBlendMode(value types.LvBlendModeT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_blend_mode(setter.CStructLvObjT, C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLayout(value uint16, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_layout(setter.CStructLvObjT, C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBaseDir(value types.LvBaseDirT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_base_dir(setter.CStructLvObjT, C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) Pos(x types.LvCoordT, y types.LvCoordT) Obj {
	C.lv_obj_set_pos(setter.CStructLvObjT, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter Obj) X(x types.LvCoordT) Obj {
	C.lv_obj_set_x(setter.CStructLvObjT, C.lv_coord_t(x))

	return setter
}
func (setter Obj) Y(y types.LvCoordT) Obj {
	C.lv_obj_set_y(setter.CStructLvObjT, C.lv_coord_t(y))

	return setter
}
func (setter Obj) Size(w types.LvCoordT, h types.LvCoordT) Obj {
	C.lv_obj_set_size(setter.CStructLvObjT, C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter Obj) Width(w types.LvCoordT) Obj {
	C.lv_obj_set_width(setter.CStructLvObjT, C.lv_coord_t(w))

	return setter
}
func (setter Obj) Height(h types.LvCoordT) Obj {
	C.lv_obj_set_height(setter.CStructLvObjT, C.lv_coord_t(h))

	return setter
}
func (setter Obj) ContentWidth(w types.LvCoordT) Obj {
	C.lv_obj_set_content_width(setter.CStructLvObjT, C.lv_coord_t(w))

	return setter
}
func (setter Obj) ContentHeight(h types.LvCoordT) Obj {
	C.lv_obj_set_content_height(setter.CStructLvObjT, C.lv_coord_t(h))

	return setter
}
func (setter Obj) Layout(layout uint32) Obj {
	C.lv_obj_set_layout(setter.CStructLvObjT, C.uint(layout))

	return setter
}
func (setter Obj) Align(align types.LvAlignT) Obj {
	C.lv_obj_set_align(setter.CStructLvObjT, C.lv_align_t(align))

	return setter
}
func (setter Obj) ExtClickArea(size types.LvCoordT) Obj {
	C.lv_obj_set_ext_click_area(setter.CStructLvObjT, C.lv_coord_t(size))

	return setter
}
func (setter Obj) ScrollbarMode(mode types.LvScrollbarModeT) Obj {
	C.lv_obj_set_scrollbar_mode(setter.CStructLvObjT, C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter Obj) ScrollDir(dir types.LvDirT) Obj {
	C.lv_obj_set_scroll_dir(setter.CStructLvObjT, C.lv_dir_t(dir))

	return setter
}
func (setter Obj) ScrollSnapX(align types.LvScrollSnapT) Obj {
	C.lv_obj_set_scroll_snap_x(setter.CStructLvObjT, C.lv_scroll_snap_t(align))

	return setter
}
func (setter Obj) ScrollSnapY(align types.LvScrollSnapT) Obj {
	C.lv_obj_set_scroll_snap_y(setter.CStructLvObjT, C.lv_scroll_snap_t(align))

	return setter
}
func (setter Obj) Parent(parent *types.LvObjT) Obj {
	C.lv_obj_set_parent(setter.CStructLvObjT, (*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter Obj) FlexFlow(flow types.LvFlexFlowT) Obj {
	C.lv_obj_set_flex_flow(setter.CStructLvObjT, C.lv_flex_flow_t(flow))

	return setter
}
func (setter Obj) FlexAlign(main_place types.LvFlexAlignT, cross_place types.LvFlexAlignT, track_cross_place types.LvFlexAlignT) Obj {
	C.lv_obj_set_flex_align(setter.CStructLvObjT, C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

	return setter
}
func (setter Obj) FlexGrow(grow uint8) Obj {
	C.lv_obj_set_flex_grow(setter.CStructLvObjT, C.uint8_t(grow))

	return setter
}
func (setter Obj) StyleFlexFlow(value types.LvFlexFlowT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_flex_flow(setter.CStructLvObjT, C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleFlexMainPlace(value types.LvFlexAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_flex_main_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleFlexCrossPlace(value types.LvFlexAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_flex_cross_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleFlexTrackPlace(value types.LvFlexAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_flex_track_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleFlexGrow(value uint8, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_flex_grow(setter.CStructLvObjT, C.uint8_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) GridAlign(column_align types.LvGridAlignT, row_align types.LvGridAlignT) Obj {
	C.lv_obj_set_grid_align(setter.CStructLvObjT, C.lv_grid_align_t(column_align), C.lv_grid_align_t(row_align))

	return setter
}
func (setter Obj) GridCell(column_align types.LvGridAlignT, col_pos uint8, col_span uint8, row_align types.LvGridAlignT, row_pos uint8, row_span uint8) Obj {
	C.lv_obj_set_grid_cell(setter.CStructLvObjT, C.lv_grid_align_t(column_align), C.uint8_t(col_pos), C.uint8_t(col_span), C.lv_grid_align_t(row_align), C.uint8_t(row_pos), C.uint8_t(row_span))

	return setter
}
func (setter Obj) StyleGridRowAlign(value types.LvGridAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_row_align(setter.CStructLvObjT, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridColumnAlign(value types.LvGridAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_column_align(setter.CStructLvObjT, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridCellColumnPos(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_cell_column_pos(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridCellColumnSpan(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_cell_column_span(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridCellRowPos(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_cell_row_pos(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridCellRowSpan(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_cell_row_span(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridCellXAlign(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_cell_x_align(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleGridCellYAlign(value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_grid_cell_y_align(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) Tile(anim_en types.LvAnimEnableT) Obj {
	C.lv_obj_set_tile(setter.CStructLvObjT, setter.CStructLvObjT, C.lv_anim_enable_t(anim_en))

	return setter
}
func (setter Obj) TileId(col_id uint32, row_id uint32, anim_en types.LvAnimEnableT) Obj {
	C.lv_obj_set_tile_id(setter.CStructLvObjT, C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

	return setter
}
