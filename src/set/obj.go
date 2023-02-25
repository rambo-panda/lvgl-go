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

func CreateObj(o *types.LvObjT) Obj {
	var _o *C.struct__lv_obj_t

	if nil == o {
		_o = C.lv_obj_create(nil) // create a screen
	} else {
		_o = Go2CObj(o, false)
	}

	return Obj{
		CStructLvObjT: _o,
	}
}

func (setter Obj) LocalStyleProp(obj *types.LvObjT, prop types.LvStylePropT, value types.LvStyleValueT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_local_style_prop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) LocalStylePropMeta(obj *types.LvObjT, prop types.LvStylePropT, meta uint16, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_local_style_prop_meta((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMinWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_min_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMaxWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_max_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMinHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_min_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleMaxHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_max_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAlign(obj *types.LvObjT, value types.LvAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformHeight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTranslateX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_translate_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTranslateY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_translate_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformZoom(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_zoom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformAngle(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformPivotX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_pivot_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransformPivotY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transform_pivot_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadTop(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_top((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadBottom(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_bottom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadLeft(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_left((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadRight(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_right((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadRow(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_row((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StylePadColumn(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_pad_column((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGradColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGradDir(obj *types.LvObjT, value types.LvGradDirT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgMainStop(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_main_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGradStop(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgGrad(obj *types.LvObjT, value *types.LvGradDscT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_grad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgDitherMode(obj *types.LvObjT, value types.LvDitherModeT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_dither_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgSrc(obj *types.LvObjT, value any, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgRecolor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgRecolorOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBgImgTiled(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_bg_img_tiled((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderSide(obj *types.LvObjT, value types.LvBorderSideT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_side((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBorderPost(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_border_post((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlineWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlineColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlineOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOutlinePad(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_outline_pad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowOfsX(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_ofs_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowOfsY(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_ofs_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowSpread(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_spread((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleShadowOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_shadow_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleImgOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleImgRecolor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleImgRecolorOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineDashWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_dash_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineDashGap(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_dash_gap((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineRounded(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLineOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_line_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcWidth(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcRounded(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleArcImgSrc(obj *types.LvObjT, value any, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_arc_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextColor(obj *types.LvObjT, value types.LvColorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextFont(obj *types.LvObjT, value *types.LvFontT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_font((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextLetterSpace(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_letter_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextLineSpace(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_line_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextDecor(obj *types.LvObjT, value types.LvTextDecorT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_decor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTextAlign(obj *types.LvObjT, value types.LvTextAlignT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_text_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleRadius(obj *types.LvObjT, value types.LvCoordT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_radius((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleClipCorner(obj *types.LvObjT, value bool, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_clip_corner((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleColorFilterDsc(obj *types.LvObjT, value *types.LvColorFilterDscT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_color_filter_dsc((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleColorFilterOpa(obj *types.LvObjT, value types.LvOpaT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_color_filter_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAnim(obj *types.LvObjT, value *types.LvAnimT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_anim((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAnimTime(obj *types.LvObjT, value uint32, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_anim_time((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleAnimSpeed(obj *types.LvObjT, value uint32, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_anim_speed((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleTransition(obj *types.LvObjT, value *types.LvStyleTransitionDscT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_transition((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBlendMode(obj *types.LvObjT, value types.LvBlendModeT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_blend_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleLayout(obj *types.LvObjT, value uint16, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) StyleBaseDir(obj *types.LvObjT, value types.LvBaseDirT, selector types.LvStyleSelectorT) Obj {
	C.lv_obj_set_style_base_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter Obj) Pos(obj *types.LvObjT, x types.LvCoordT, y types.LvCoordT) Obj {
	C.lv_obj_set_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter Obj) X(obj *types.LvObjT, x types.LvCoordT) Obj {
	C.lv_obj_set_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x))

	return setter
}
func (setter Obj) Y(obj *types.LvObjT, y types.LvCoordT) Obj {
	C.lv_obj_set_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(y))

	return setter
}
func (setter Obj) Size(obj *types.LvObjT, w types.LvCoordT, h types.LvCoordT) Obj {
	C.lv_obj_set_size((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter Obj) Width(obj *types.LvObjT, w types.LvCoordT) Obj {
	C.lv_obj_set_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	return setter
}
func (setter Obj) Height(obj *types.LvObjT, h types.LvCoordT) Obj {
	C.lv_obj_set_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

	return setter
}
func (setter Obj) ContentWidth(obj *types.LvObjT, w types.LvCoordT) Obj {
	C.lv_obj_set_content_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	return setter
}
func (setter Obj) ContentHeight(obj *types.LvObjT, h types.LvCoordT) Obj {
	C.lv_obj_set_content_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

	return setter
}
func (setter Obj) Layout(obj *types.LvObjT, layout uint32) Obj {
	C.lv_obj_set_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(layout))

	return setter
}
func (setter Obj) Align(obj *types.LvObjT, align types.LvAlignT) Obj {
	C.lv_obj_set_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(align))

	return setter
}
func (setter Obj) ExtClickArea(obj *types.LvObjT, size types.LvCoordT) Obj {
	C.lv_obj_set_ext_click_area((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(size))

	return setter
}
func (setter Obj) ScrollbarMode(obj *types.LvObjT, mode types.LvScrollbarModeT) Obj {
	C.lv_obj_set_scrollbar_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter Obj) ScrollDir(obj *types.LvObjT, dir types.LvDirT) Obj {
	C.lv_obj_set_scroll_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dir_t(dir))

	return setter
}
func (setter Obj) ScrollSnapX(obj *types.LvObjT, align types.LvScrollSnapT) Obj {
	C.lv_obj_set_scroll_snap_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

	return setter
}
func (setter Obj) ScrollSnapY(obj *types.LvObjT, align types.LvScrollSnapT) Obj {
	C.lv_obj_set_scroll_snap_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

	return setter
}
func (setter Obj) Parent(obj *types.LvObjT, parent *types.LvObjT) Obj {
	C.lv_obj_set_parent((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

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
