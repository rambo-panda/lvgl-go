package lib

/*
#include "lv_init.h"
*/
import "C"

type LV_PALETTE_T = C.lv_palette_t
type LV_COLOR_T = C.lv_color_t
type LV_OPA_T = C.lv_opa_t
type LV_COLOR_HSV_T = C.lv_color_hsv_t
type LV_STYLE_SELECTOR_T = C.lv_style_selector_t
type LV_GRAD_DIR_T = C.lv_grad_dir_t

type LV_ANIM_EXEC_XCB_T = C.lv_anim_exec_xcb_t
type C_SHORT = C.uint16_t

const (
	LV_ANIM_ON uint32 = C.LV_ANIM_ON

	LV_PART_MAIN         int = C.LV_PART_MAIN
	LV_PART_SCROLLBAR    int = C.LV_PART_SCROLLBAR
	LV_PART_INDICATOR    int = C.LV_PART_INDICATOR
	LV_PART_KNOB         int = C.LV_PART_KNOB
	LV_PART_SELECTED     int = C.LV_PART_SELECTED
	LV_PART_ITEMS        int = C.LV_PART_ITEMS
	LV_PART_TICKS        int = C.LV_PART_TICKS
	LV_PART_CURSOR       int = C.LV_PART_CURSOR
	LV_PART_CUSTOM_FIRST int = C.LV_PART_CUSTOM_FIRST
	LV_PART_ANY          int = C.LV_PART_ANY
)

const (
	LV_GRAD_DIR_NONE LV_GRAD_DIR_T = C.LV_GRAD_DIR_NONE
	LV_GRAD_DIR_VER  LV_GRAD_DIR_T = C.LV_GRAD_DIR_VER
	LV_GRAD_DIR_HOR  LV_GRAD_DIR_T = C.LV_GRAD_DIR_HOR
)

const LV_ANIM_REPEAT_INFINITE C_SHORT = C.LV_ANIM_REPEAT_INFINITE
