package lib

/*
#include "lv_init.h"
static const int _LV_SIZE_CONTENT = LV_SIZE_CONTENT;
*/
import "C"

type LvObjT C.lv_obj_t
type LvStylePropT C.lv_style_prop_t
type LvStyleValueT C.lv_style_value_t
type LvStyleSelectorT C.lv_style_selector_t
type LvAlignT C.lv_align_t
type LvColorT C.lv_color_t
type LvOpaT C.lv_opa_t
type LvGradDirT C.lv_grad_dir_t
type LvGradDscT C.lv_grad_dsc_t
type LvDitherModeT C.lv_dither_mode_t
type LvBorderSideT C.lv_border_side_t
type LvFontT C.lv_font_t
type LvTextDecorT C.lv_text_decor_t
type LvTextAlignT C.lv_text_align_t
type LvColorFilterDscT C.lv_color_filter_dsc_t
type LvAnimT C.lv_anim_t
type LvStyleTransitionDscT C.lv_style_transition_dsc_t
type LvBlendModeT C.lv_blend_mode_t
type LvBaseDirT C.lv_base_dir_t
type LvScrollbarModeT C.lv_scrollbar_mode_t
type LvDirT C.lv_dir_t
type LvScrollSnapT C.lv_scroll_snap_t
type LvCoordT C.lv_coord_t
type LvImgSizeModeT C.lv_img_size_mode_t
type LvLabelLongModeT C.lv_label_long_mode_t

const (
	LV_SIZE_CONTENT C.int = C._LV_SIZE_CONTENT

	LV_ALIGN_DEFAULT      C.int = C.LV_ALIGN_DEFAULT
	LV_ALIGN_TOP_LEFT     C.int = C.LV_ALIGN_TOP_LEFT
	LV_ALIGN_TOP_MID      C.int = C.LV_ALIGN_TOP_MID
	LV_ALIGN_TOP_RIGHT    C.int = C.LV_ALIGN_TOP_RIGHT
	LV_ALIGN_BOTTOM_LEFT  C.int = C.LV_ALIGN_BOTTOM_LEFT
	LV_ALIGN_BOTTOM_MID   C.int = C.LV_ALIGN_BOTTOM_MID
	LV_ALIGN_BOTTOM_RIGHT C.int = C.LV_ALIGN_BOTTOM_RIGHT
	LV_ALIGN_LEFT_MID     C.int = C.LV_ALIGN_LEFT_MID
	LV_ALIGN_RIGHT_MID    C.int = C.LV_ALIGN_RIGHT_MID
	LV_ALIGN_CENTER       C.int = C.LV_ALIGN_CENTER

	LV_ALIGN_OUT_TOP_LEFT     C.int = C.LV_ALIGN_OUT_TOP_LEFT
	LV_ALIGN_OUT_TOP_MID      C.int = C.LV_ALIGN_OUT_TOP_MID
	LV_ALIGN_OUT_TOP_RIGHT    C.int = C.LV_ALIGN_OUT_TOP_RIGHT
	LV_ALIGN_OUT_BOTTOM_LEFT  C.int = C.LV_ALIGN_OUT_BOTTOM_LEFT
	LV_ALIGN_OUT_BOTTOM_MID   C.int = C.LV_ALIGN_OUT_BOTTOM_MID
	LV_ALIGN_OUT_BOTTOM_RIGHT C.int = C.LV_ALIGN_OUT_BOTTOM_RIGHT
	LV_ALIGN_OUT_LEFT_TOP     C.int = C.LV_ALIGN_OUT_LEFT_TOP
	LV_ALIGN_OUT_LEFT_MID     C.int = C.LV_ALIGN_OUT_LEFT_MID
	LV_ALIGN_OUT_LEFT_BOTTOM  C.int = C.LV_ALIGN_OUT_LEFT_BOTTOM
	LV_ALIGN_OUT_RIGHT_TOP    C.int = C.LV_ALIGN_OUT_RIGHT_TOP
	LV_ALIGN_OUT_RIGHT_MID    C.int = C.LV_ALIGN_OUT_RIGHT_MID
	LV_ALIGN_OUT_RIGHT_BOTTOM C.int = C.LV_ALIGN_OUT_RIGHT_BOTTOM
)
