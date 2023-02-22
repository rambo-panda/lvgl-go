package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
  lib "lvgl-go/src/lib"
)
func WidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func MinWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_min_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func MaxWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_max_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func HeightForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_height((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func MinHeightForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_min_height((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func MaxHeightForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_max_height((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func XForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_x((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func YForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_y((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func AlignForStyle(style *lib.LvStyleT ,value lib.LvAlignT)  {
   C.lv_style_set_align((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_align_t(value))

  
}
func TransformWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_transform_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TransformHeightForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_transform_height((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TranslateXForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_translate_x((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TranslateYForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_translate_y((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TransformZoomForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_transform_zoom((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TransformAngleForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_transform_angle((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TransformPivotXForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_transform_pivot_x((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TransformPivotYForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_transform_pivot_y((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func PadTopForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_pad_top((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func PadBottomForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_pad_bottom((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func PadLeftForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_pad_left((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func PadRightForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_pad_right((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func PadRowForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_pad_row((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func PadColumnForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_pad_column((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func BgColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_bg_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func BgOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_bg_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func BgGradColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_bg_grad_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func BgGradDirForStyle(style *lib.LvStyleT ,value lib.LvGradDirT)  {
   C.lv_style_set_bg_grad_dir((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_grad_dir_t(value))

  
}
func BgMainStopForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_bg_main_stop((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func BgGradStopForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_bg_grad_stop((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func BgGradForStyle(style *lib.LvStyleT ,value *lib.LvGradDscT )  {
   C.lv_style_set_bg_grad((*C.lv_style_t)(unsafe.Pointer(style)),(*C.lv_grad_dsc_t)(unsafe.Pointer(value)))

  
}
func BgDitherModeForStyle(style *lib.LvStyleT ,value lib.LvDitherModeT)  {
   C.lv_style_set_bg_dither_mode((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_dither_mode_t(value))

  
}
func BgImgSrcForStyle(style *lib.LvStyleT ,value any)  {
   C.lv_style_set_bg_img_src((*C.lv_style_t)(unsafe.Pointer(style)),unsafe.Pointer(&value))

  
}
func BgImgOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_bg_img_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func BgImgRecolorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_bg_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func BgImgRecolorOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_bg_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func BgImgTiledForStyle(style *lib.LvStyleT ,value bool)  {
   C.lv_style_set_bg_img_tiled((*C.lv_style_t)(unsafe.Pointer(style)),C.bool(value))

  
}
func BorderColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_border_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func BorderOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_border_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func BorderWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_border_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func BorderSideForStyle(style *lib.LvStyleT ,value lib.LvBorderSideT)  {
   C.lv_style_set_border_side((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_border_side_t(value))

  
}
func BorderPostForStyle(style *lib.LvStyleT ,value bool)  {
   C.lv_style_set_border_post((*C.lv_style_t)(unsafe.Pointer(style)),C.bool(value))

  
}
func OutlineWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_outline_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func OutlineColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_outline_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func OutlineOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_outline_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func OutlinePadForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_outline_pad((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ShadowWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_shadow_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ShadowOfsXForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_shadow_ofs_x((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ShadowOfsYForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_shadow_ofs_y((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ShadowSpreadForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_shadow_spread((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ShadowColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_shadow_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func ShadowOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_shadow_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func ImgOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_img_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func ImgRecolorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_img_recolor((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func ImgRecolorOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_img_recolor_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func LineWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_line_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func LineDashWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_line_dash_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func LineDashGapForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_line_dash_gap((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func LineRoundedForStyle(style *lib.LvStyleT ,value bool)  {
   C.lv_style_set_line_rounded((*C.lv_style_t)(unsafe.Pointer(style)),C.bool(value))

  
}
func LineColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_line_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func LineOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_line_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func ArcWidthForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_arc_width((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ArcRoundedForStyle(style *lib.LvStyleT ,value bool)  {
   C.lv_style_set_arc_rounded((*C.lv_style_t)(unsafe.Pointer(style)),C.bool(value))

  
}
func ArcColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_arc_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func ArcOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_arc_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func ArcImgSrcForStyle(style *lib.LvStyleT ,value any)  {
   C.lv_style_set_arc_img_src((*C.lv_style_t)(unsafe.Pointer(style)),unsafe.Pointer(&value))

  
}
func TextColorForStyle(style *lib.LvStyleT ,value lib.LvColorT)  {
   C.lv_style_set_text_color((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_color_t(value))

  
}
func TextOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_text_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func TextFontForStyle(style *lib.LvStyleT ,value *lib.LvFontT )  {
   C.lv_style_set_text_font((*C.lv_style_t)(unsafe.Pointer(style)),(*C.lv_font_t)(unsafe.Pointer(value)))

  
}
func TextLetterSpaceForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_text_letter_space((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TextLineSpaceForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_text_line_space((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func TextDecorForStyle(style *lib.LvStyleT ,value lib.LvTextDecorT)  {
   C.lv_style_set_text_decor((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_text_decor_t(value))

  
}
func TextAlignForStyle(style *lib.LvStyleT ,value lib.LvTextAlignT)  {
   C.lv_style_set_text_align((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_text_align_t(value))

  
}
func RadiusForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_radius((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func ClipCornerForStyle(style *lib.LvStyleT ,value bool)  {
   C.lv_style_set_clip_corner((*C.lv_style_t)(unsafe.Pointer(style)),C.bool(value))

  
}
func OpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func ColorFilterDscForStyle(style *lib.LvStyleT ,value *lib.LvColorFilterDscT )  {
   C.lv_style_set_color_filter_dsc((*C.lv_style_t)(unsafe.Pointer(style)),(*C.lv_color_filter_dsc_t)(unsafe.Pointer(value)))

  
}
func ColorFilterOpaForStyle(style *lib.LvStyleT ,value lib.LvOpaT)  {
   C.lv_style_set_color_filter_opa((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_opa_t(value))

  
}
func AnimForStyle(style *lib.LvStyleT ,value *lib.LvAnimT )  {
   C.lv_style_set_anim((*C.lv_style_t)(unsafe.Pointer(style)),(*C.lv_anim_t)(unsafe.Pointer(value)))

  
}
func AnimTimeForStyle(style *lib.LvStyleT ,value uint32)  {
   C.lv_style_set_anim_time((*C.lv_style_t)(unsafe.Pointer(style)),C.uint(value))

  
}
func AnimSpeedForStyle(style *lib.LvStyleT ,value uint32)  {
   C.lv_style_set_anim_speed((*C.lv_style_t)(unsafe.Pointer(style)),C.uint(value))

  
}
func TransitionForStyle(style *lib.LvStyleT ,value *lib.LvStyleTransitionDscT )  {
   C.lv_style_set_transition((*C.lv_style_t)(unsafe.Pointer(style)),(*C.lv_style_transition_dsc_t)(unsafe.Pointer(value)))

  
}
func BlendModeForStyle(style *lib.LvStyleT ,value lib.LvBlendModeT)  {
   C.lv_style_set_blend_mode((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_blend_mode_t(value))

  
}
func LayoutForStyle(style *lib.LvStyleT ,value uint16)  {
   C.lv_style_set_layout((*C.lv_style_t)(unsafe.Pointer(style)),C.ushort(value))

  
}
func BaseDirForStyle(style *lib.LvStyleT ,value lib.LvBaseDirT)  {
   C.lv_style_set_base_dir((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_base_dir_t(value))

  
}
func PropForStyle(style *lib.LvStyleT ,prop lib.LvStylePropT,value lib.LvStyleValueT)  {
   C.lv_style_set_prop((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_style_prop_t(prop),C.lv_style_value_t(value))

  
}
func PropMetaForStyle(style *lib.LvStyleT ,prop lib.LvStylePropT,meta uint16)  {
   C.lv_style_set_prop_meta((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_style_prop_t(prop),C.ushort(meta))

  
}
func FlexFlowForStyle(style *lib.LvStyleT ,value lib.LvFlexFlowT)  {
   C.lv_style_set_flex_flow((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_flex_flow_t(value))

  
}
func FlexMainPlaceForStyle(style *lib.LvStyleT ,value lib.LvFlexAlignT)  {
   C.lv_style_set_flex_main_place((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_flex_align_t(value))

  
}
func FlexCrossPlaceForStyle(style *lib.LvStyleT ,value lib.LvFlexAlignT)  {
   C.lv_style_set_flex_cross_place((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_flex_align_t(value))

  
}
func FlexTrackPlaceForStyle(style *lib.LvStyleT ,value lib.LvFlexAlignT)  {
   C.lv_style_set_flex_track_place((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_flex_align_t(value))

  
}
func FlexGrowForStyle(style *lib.LvStyleT ,value uint8)  {
   C.lv_style_set_flex_grow((*C.lv_style_t)(unsafe.Pointer(style)),C.uint8_t(value))

  
}
func GridRowAlignForStyle(style *lib.LvStyleT ,value lib.LvGridAlignT)  {
   C.lv_style_set_grid_row_align((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_grid_align_t(value))

  
}
func GridColumnAlignForStyle(style *lib.LvStyleT ,value lib.LvGridAlignT)  {
   C.lv_style_set_grid_column_align((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_grid_align_t(value))

  
}
func GridCellColumnPosForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_grid_cell_column_pos((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func GridCellColumnSpanForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_grid_cell_column_span((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func GridCellRowPosForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_grid_cell_row_pos((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func GridCellRowSpanForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_grid_cell_row_span((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func GridCellXAlignForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_grid_cell_x_align((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}
func GridCellYAlignForStyle(style *lib.LvStyleT ,value lib.LvCoordT)  {
   C.lv_style_set_grid_cell_y_align((*C.lv_style_t)(unsafe.Pointer(style)),C.lv_coord_t(value))

  
}