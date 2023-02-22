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

func LocalStylePropForObj(obj *lib.LvObjT, prop lib.LvStylePropT, value lib.LvStyleValueT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_local_style_prop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.lv_style_value_t(value), C.lv_style_selector_t(selector))

}
func LocalStylePropMetaForObj(obj *lib.LvObjT, prop lib.LvStylePropT, meta uint16, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_local_style_prop_meta((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_style_prop_t(prop), C.ushort(meta), C.lv_style_selector_t(selector))

}
func StyleWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleMinWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_min_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleMaxWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_max_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleHeightForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleMinHeightForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_min_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleMaxHeightForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_max_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleXForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleYForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleAlignForObj(obj *lib.LvObjT, value lib.LvAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(value), C.lv_style_selector_t(selector))

}
func StyleTransformWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transform_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTransformHeightForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transform_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTranslateXForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_translate_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTranslateYForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_translate_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTransformZoomForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transform_zoom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTransformAngleForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transform_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTransformPivotXForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transform_pivot_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTransformPivotYForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transform_pivot_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StylePadTopForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_pad_top((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StylePadBottomForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_pad_bottom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StylePadLeftForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_pad_left((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StylePadRightForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_pad_right((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StylePadRowForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_pad_row((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StylePadColumnForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_pad_column((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleBgColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleBgOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleBgGradColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_grad_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleBgGradDirForObj(obj *lib.LvObjT, value lib.LvGradDirT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_grad_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grad_dir_t(value), C.lv_style_selector_t(selector))

}
func StyleBgMainStopForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_main_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleBgGradStopForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_grad_stop((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleBgGradForObj(obj *lib.LvObjT, value *lib.LvGradDscT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_grad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_grad_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

}
func StyleBgDitherModeForObj(obj *lib.LvObjT, value lib.LvDitherModeT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_dither_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dither_mode_t(value), C.lv_style_selector_t(selector))

}
func StyleBgImgSrcForObj(obj *lib.LvObjT, value any, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

}
func StyleBgImgOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleBgImgRecolorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleBgImgRecolorOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleBgImgTiledForObj(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_bg_img_tiled((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

}
func StyleBorderColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_border_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleBorderOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_border_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleBorderWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_border_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleBorderSideForObj(obj *lib.LvObjT, value lib.LvBorderSideT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_border_side((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_border_side_t(value), C.lv_style_selector_t(selector))

}
func StyleBorderPostForObj(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_border_post((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

}
func StyleOutlineWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_outline_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleOutlineColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_outline_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleOutlineOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_outline_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleOutlinePadForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_outline_pad((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleShadowWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_shadow_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleShadowOfsXForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_shadow_ofs_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleShadowOfsYForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_shadow_ofs_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleShadowSpreadForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_shadow_spread((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleShadowColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_shadow_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleShadowOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_shadow_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleImgOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_img_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleImgRecolorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_img_recolor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleImgRecolorOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_img_recolor_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleLineWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_line_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleLineDashWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_line_dash_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleLineDashGapForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_line_dash_gap((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleLineRoundedForObj(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_line_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

}
func StyleLineColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_line_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleLineOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_line_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleArcWidthForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_arc_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleArcRoundedForObj(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_arc_rounded((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

}
func StyleArcColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_arc_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleArcOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_arc_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleArcImgSrcForObj(obj *lib.LvObjT, value any, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_arc_img_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&value), C.lv_style_selector_t(selector))

}
func StyleTextColorForObj(obj *lib.LvObjT, value lib.LvColorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_color((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_color_t(value), C.lv_style_selector_t(selector))

}
func StyleTextOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleTextFontForObj(obj *lib.LvObjT, value *lib.LvFontT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_font((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_font_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

}
func StyleTextLetterSpaceForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_letter_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTextLineSpaceForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_line_space((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleTextDecorForObj(obj *lib.LvObjT, value lib.LvTextDecorT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_decor((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_decor_t(value), C.lv_style_selector_t(selector))

}
func StyleTextAlignForObj(obj *lib.LvObjT, value lib.LvTextAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_text_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_text_align_t(value), C.lv_style_selector_t(selector))

}
func StyleRadiusForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_radius((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleClipCornerForObj(obj *lib.LvObjT, value bool, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_clip_corner((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(value), C.lv_style_selector_t(selector))

}
func StyleOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleColorFilterDscForObj(obj *lib.LvObjT, value *lib.LvColorFilterDscT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_color_filter_dsc((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

}
func StyleColorFilterOpaForObj(obj *lib.LvObjT, value lib.LvOpaT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_color_filter_opa((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_opa_t(value), C.lv_style_selector_t(selector))

}
func StyleAnimForObj(obj *lib.LvObjT, value *lib.LvAnimT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_anim((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_anim_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

}
func StyleAnimTimeForObj(obj *lib.LvObjT, value uint32, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_anim_time((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

}
func StyleAnimSpeedForObj(obj *lib.LvObjT, value uint32, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_anim_speed((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(value), C.lv_style_selector_t(selector))

}
func StyleTransitionForObj(obj *lib.LvObjT, value *lib.LvStyleTransitionDscT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_transition((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)), C.lv_style_selector_t(selector))

}
func StyleBlendModeForObj(obj *lib.LvObjT, value lib.LvBlendModeT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_blend_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_blend_mode_t(value), C.lv_style_selector_t(selector))

}
func StyleLayoutForObj(obj *lib.LvObjT, value uint16, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(value), C.lv_style_selector_t(selector))

}
func StyleBaseDirForObj(obj *lib.LvObjT, value lib.LvBaseDirT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_base_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_base_dir_t(value), C.lv_style_selector_t(selector))

}
func PosForObj(obj *lib.LvObjT, x lib.LvCoordT, y lib.LvCoordT) {
	C.lv_obj_set_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x), C.lv_coord_t(y))

}
func XForObj(obj *lib.LvObjT, x lib.LvCoordT) {
	C.lv_obj_set_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x))

}
func YForObj(obj *lib.LvObjT, y lib.LvCoordT) {
	C.lv_obj_set_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(y))

}
func SizeForObj(obj *lib.LvObjT, w lib.LvCoordT, h lib.LvCoordT) {
	C.lv_obj_set_size((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w), C.lv_coord_t(h))

}
func WidthForObj(obj *lib.LvObjT, w lib.LvCoordT) {
	C.lv_obj_set_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

}
func HeightForObj(obj *lib.LvObjT, h lib.LvCoordT) {
	C.lv_obj_set_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

}
func ContentWidthForObj(obj *lib.LvObjT, w lib.LvCoordT) {
	C.lv_obj_set_content_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

}
func ContentHeightForObj(obj *lib.LvObjT, h lib.LvCoordT) {
	C.lv_obj_set_content_height((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(h))

}
func LayoutForObj(obj *lib.LvObjT, layout uint32) {
	C.lv_obj_set_layout((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint(layout))

}
func AlignForObj(obj *lib.LvObjT, align lib.LvAlignT) {
	C.lv_obj_set_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_align_t(align))

}
func ExtClickAreaForObj(obj *lib.LvObjT, size lib.LvCoordT) {
	C.lv_obj_set_ext_click_area((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(size))

}
func ScrollbarModeForObj(obj *lib.LvObjT, mode lib.LvScrollbarModeT) {
	C.lv_obj_set_scrollbar_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scrollbar_mode_t(mode))

}
func ScrollDirForObj(obj *lib.LvObjT, dir lib.LvDirT) {
	C.lv_obj_set_scroll_dir((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_dir_t(dir))

}
func ScrollSnapXForObj(obj *lib.LvObjT, align lib.LvScrollSnapT) {
	C.lv_obj_set_scroll_snap_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

}
func ScrollSnapYForObj(obj *lib.LvObjT, align lib.LvScrollSnapT) {
	C.lv_obj_set_scroll_snap_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_scroll_snap_t(align))

}
func ParentForObj(obj *lib.LvObjT, parent *lib.LvObjT) {
	C.lv_obj_set_parent((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), (*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

}
func FlexFlowForObj(obj *lib.LvObjT, flow lib.LvFlexFlowT) {
	C.lv_obj_set_flex_flow((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_flex_flow_t(flow))

}
func FlexAlignForObj(obj *lib.LvObjT, main_place lib.LvFlexAlignT, cross_place lib.LvFlexAlignT, track_cross_place lib.LvFlexAlignT) {
	C.lv_obj_set_flex_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_flex_align_t(main_place), C.lv_flex_align_t(cross_place), C.lv_flex_align_t(track_cross_place))

}
func FlexGrowForObj(obj *lib.LvObjT, grow uint8) {
	C.lv_obj_set_flex_grow((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint8_t(grow))

}
func StyleFlexFlowForObj(obj *lib.LvObjT, value lib.LvFlexFlowT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_flex_flow((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_flex_flow_t(value), C.lv_style_selector_t(selector))

}
func StyleFlexMainPlaceForObj(obj *lib.LvObjT, value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_flex_main_place((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

}
func StyleFlexCrossPlaceForObj(obj *lib.LvObjT, value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_flex_cross_place((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

}
func StyleFlexTrackPlaceForObj(obj *lib.LvObjT, value lib.LvFlexAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_flex_track_place((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_flex_align_t(value), C.lv_style_selector_t(selector))

}
func StyleFlexGrowForObj(obj *lib.LvObjT, value uint8, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_flex_grow((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.uint8_t(value), C.lv_style_selector_t(selector))

}
func GridAlignForObj(obj *lib.LvObjT, column_align lib.LvGridAlignT, row_align lib.LvGridAlignT) {
	C.lv_obj_set_grid_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grid_align_t(column_align), C.lv_grid_align_t(row_align))

}
func GridCellForObj(obj *lib.LvObjT, column_align lib.LvGridAlignT, col_pos uint8, col_span uint8, row_align lib.LvGridAlignT, row_pos uint8, row_span uint8) {
	C.lv_obj_set_grid_cell((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grid_align_t(column_align), C.uint8_t(col_pos), C.uint8_t(col_span), C.lv_grid_align_t(row_align), C.uint8_t(row_pos), C.uint8_t(row_span))

}
func StyleGridRowAlignForObj(obj *lib.LvObjT, value lib.LvGridAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_row_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

}
func StyleGridColumnAlignForObj(obj *lib.LvObjT, value lib.LvGridAlignT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_column_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_grid_align_t(value), C.lv_style_selector_t(selector))

}
func StyleGridCellColumnPosForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_cell_column_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleGridCellColumnSpanForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_cell_column_span((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleGridCellRowPosForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_cell_row_pos((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleGridCellRowSpanForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_cell_row_span((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleGridCellXAlignForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_cell_x_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func StyleGridCellYAlignForObj(obj *lib.LvObjT, value lib.LvCoordT, selector lib.LvStyleSelectorT) {
	C.lv_obj_set_style_grid_cell_y_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(value), C.lv_style_selector_t(selector))

}
func TileForObj(tv *lib.LvObjT, tile_obj *lib.LvObjT, anim_en lib.LvAnimEnableT) {
	C.lv_obj_set_tile((*C.struct__lv_obj_t)(unsafe.Pointer(tv)), (*C.lv_obj_t)(unsafe.Pointer(tile_obj)), C.lv_anim_enable_t(anim_en))

}
func TileIdForObj(tv *lib.LvObjT, col_id uint32, row_id uint32, anim_en lib.LvAnimEnableT) {
	C.lv_obj_set_tile_id((*C.struct__lv_obj_t)(unsafe.Pointer(tv)), C.uint(col_id), C.uint(row_id), C.lv_anim_enable_t(anim_en))

}
