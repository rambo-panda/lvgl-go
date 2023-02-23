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

func CreateObj(o *lib.LvObjT) SetObj {
	return SetObj{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetObj) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetObj) SetLocalStyleProp(obj *lib.LvObjT, prop lib.LvStylePropT, value lib.LvStyleValueT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_local_style_prop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetLocalStylePropMeta(obj *lib.LvObjT, prop lib.LvStylePropT, meta uint16, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_local_style_prop_meta((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMinWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_min_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMaxWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_max_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleHeight(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMinHeight(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_min_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleMaxHeight(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_max_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleX(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleY(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAlign(obj *lib.LvObjT, value lib.LvAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformHeight(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTranslateX(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_translate_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTranslateY(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_translate_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformZoom(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_zoom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformAngle(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformPivotX(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_pivot_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransformPivotY(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transform_pivot_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadTop(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_top((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadBottom(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_bottom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadLeft(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_left((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadRight(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_right((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadRow(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_row((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStylePadColumn(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_pad_column((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradDir(obj *lib.LvObjT, value lib.LvGradDirT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgMainStop(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_main_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGradStop(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgGrad(obj *lib.LvObjT, value *lib.LvGradDscT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_grad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgDitherMode(obj *lib.LvObjT, value lib.LvDitherModeT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_dither_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgSrc(obj *lib.LvObjT, value any, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgRecolor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgRecolorOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBgImgTiled(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_bg_img_tiled((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderSide(obj *lib.LvObjT, value lib.LvBorderSideT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_side((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_border_side_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBorderPost(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_border_post((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlineOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOutlinePad(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_outline_pad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOfsX(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_ofs_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOfsY(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_ofs_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowSpread(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_spread((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleShadowOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_shadow_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgRecolor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleImgRecolorOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineDashWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_dash_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineDashGap(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_dash_gap((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineRounded(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLineOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_line_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcWidth(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcRounded(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleArcImgSrc(obj *lib.LvObjT, value any, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_arc_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextColor(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextFont(obj *lib.LvObjT, value *lib.LvFontT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_font((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextLetterSpace(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_letter_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextLineSpace(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_line_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextDecor(obj *lib.LvObjT, value lib.LvTextDecorT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_decor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTextAlign(obj *lib.LvObjT, value lib.LvTextAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_text_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleRadius(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_radius((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleClipCorner(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_clip_corner((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleColorFilterDsc(obj *lib.LvObjT, value *lib.LvColorFilterDscT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_color_filter_dsc((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleColorFilterOpa(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_color_filter_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnim(obj *lib.LvObjT, value *lib.LvAnimT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnimTime(obj *lib.LvObjT, value uint32, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim_time((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleAnimSpeed(obj *lib.LvObjT, value uint32, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_anim_speed((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleTransition(obj *lib.LvObjT, value *lib.LvStyleTransitionDscT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_transition((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBlendMode(obj *lib.LvObjT, value lib.LvBlendModeT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_blend_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleLayout(obj *lib.LvObjT, value uint16, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleBaseDir(obj *lib.LvObjT, value lib.LvBaseDirT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_base_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetPos(obj *lib.LvObjT, x lib.LvCoordT, y lib.LvCoordT) SetObj {
	C.lv_obj_set_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter SetObj) SetX(obj *lib.LvObjT, x lib.LvCoordT) SetObj {
	C.lv_obj_set_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x))

	return setter
}
func (setter SetObj) SetY(obj *lib.LvObjT, y lib.LvCoordT) SetObj {
	C.lv_obj_set_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(y))

	return setter
}
func (setter SetObj) SetSize(obj *lib.LvObjT, w lib.LvCoordT, h lib.LvCoordT) SetObj {
	C.lv_obj_set_size((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetWidth(obj *lib.LvObjT, w lib.LvCoordT) SetObj {
	C.lv_obj_set_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	return setter
}
func (setter SetObj) SetHeight(obj *lib.LvObjT, h lib.LvCoordT) SetObj {
	C.lv_obj_set_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetContentWidth(obj *lib.LvObjT, w lib.LvCoordT) SetObj {
	C.lv_obj_set_content_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	return setter
}
func (setter SetObj) SetContentHeight(obj *lib.LvObjT, h lib.LvCoordT) SetObj {
	C.lv_obj_set_content_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

	return setter
}
func (setter SetObj) SetLayout(obj *lib.LvObjT, layout uint32) SetObj {
	C.lv_obj_set_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(layout))

	return setter
}
func (setter SetObj) SetAlign(obj *lib.LvObjT, align lib.LvAlignT) SetObj {
	C.lv_obj_set_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(align))

	return setter
}
func (setter SetObj) SetExtClickArea(obj *lib.LvObjT, size lib.LvCoordT) SetObj {
	C.lv_obj_set_ext_click_area((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(size))

	return setter
}
func (setter SetObj) SetScrollbarMode(obj *lib.LvObjT, mode lib.LvScrollbarModeT) SetObj {
	C.lv_obj_set_scrollbar_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scrollbar_mode_t(mode))

	return setter
}
func (setter SetObj) SetScrollDir(obj *lib.LvObjT, dir lib.LvDirT) SetObj {
	C.lv_obj_set_scroll_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dir_t(dir))

	return setter
}
func (setter SetObj) SetScrollSnapX(obj *lib.LvObjT, align lib.LvScrollSnapT) SetObj {
	C.lv_obj_set_scroll_snap_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

	return setter
}
func (setter SetObj) SetScrollSnapY(obj *lib.LvObjT, align lib.LvScrollSnapT) SetObj {
	C.lv_obj_set_scroll_snap_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

	return setter
}
func (setter SetObj) SetParent(obj *lib.LvObjT, parent *lib.LvObjT) SetObj {
	C.lv_obj_set_parent((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter SetObj) SetFlexFlow(flow lib.LvFlexFlowT) SetObj {
	C.lv_obj_set_flex_flow(setter.CStructLvObjT, C.lv_flex_flow_t(flow))

	return setter
}
func (setter SetObj) SetFlexAlign(main_place lib.LvFlexAlignT, cross_place lib.LvFlexAlignT, track_cross_place lib.LvFlexAlignT) SetObj {
	C.lv_obj_set_flex_align(setter.CStructLvObjT, C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

	return setter
}
func (setter SetObj) SetFlexGrow(grow uint8) SetObj {
	C.lv_obj_set_flex_grow(setter.CStructLvObjT, C.uint8_t(grow))

	return setter
}
func (setter SetObj) SetStyleFlexFlow(value lib.LvFlexFlowT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_flow(setter.CStructLvObjT, C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexMainPlace(value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_main_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexCrossPlace(value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_cross_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexTrackPlace(value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_track_place(setter.CStructLvObjT, C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleFlexGrow(value uint8, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_flex_grow(setter.CStructLvObjT, C.uint8_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetGridAlign(column_align lib.LvGridAlignT, row_align lib.LvGridAlignT) SetObj {
	C.lv_obj_set_grid_align(setter.CStructLvObjT, C.lv_grid_align_t(column_align), C.lv_grid_align_t(row_align))

	return setter
}
func (setter SetObj) SetGridCell(column_align lib.LvGridAlignT, col_pos uint8, col_span uint8, row_align lib.LvGridAlignT, row_pos uint8, row_span uint8) SetObj {
	C.lv_obj_set_grid_cell(setter.CStructLvObjT, C.lv_grid_align_t(column_align), C.uint8_t(col_pos), C.uint8_t(col_span), C.lv_grid_align_t(row_align), C.uint8_t(row_pos), C.uint8_t(row_span))

	return setter
}
func (setter SetObj) SetStyleGridRowAlign(value lib.LvGridAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_row_align(setter.CStructLvObjT, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridColumnAlign(value lib.LvGridAlignT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_column_align(setter.CStructLvObjT, C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellColumnPos(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_column_pos(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellColumnSpan(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_column_span(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellRowPos(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_row_pos(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellRowSpan(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_row_span(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellXAlign(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_x_align(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetStyleGridCellYAlign(value lib.LvCoordT, selector lib.LvStyleSelectorT) SetObj {
	C.lv_obj_set_style_grid_cell_y_align(setter.CStructLvObjT, C.lv_coord_t(value), C.lv_style_selector_t(selector))

	return setter
}
func (setter SetObj) SetTile(anim_en lib.LvAnimEnableT) SetObj {
	C.lv_obj_set_tile(setter.CStructLvObjT, setter.CStructLvObjT, C.lv_anim_enable_t(anim_en))

	return setter
}
func (setter SetObj) SetTileId(col_id uint32, row_id uint32, anim_en lib.LvAnimEnableT) SetObj {
	C.lv_obj_set_tile_id(setter.CStructLvObjT, C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

	return setter
}
