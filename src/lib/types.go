package lib

/*
#include "lv_init.h"
*/
import "C"
import "unsafe"

// ================ C Types
type LV_ANIM_EXEC_XCB_T = C.lv_anim_exec_xcb_t

// type LV_STYLE_TRANSITION_DSC_INIT = C.lv_style_transition_dsc_init
type LV_STYLE_TRANSITION_DSC_T = C.lv_style_transition_dsc_t
type LV_COLOR_FILTER_DSC_T = C.lv_color_filter_dsc_t
type LV_LABEL_LONG_MODE_T = C.lv_label_long_mode_t
type C_SHORT = C.uint16_t
type LV_PART_T = C.lv_part_t
type LV_POINT_T = C.lv_point_t
type LV_GRAD_DSC_T = C.lv_grad_dsc_t
type LV_FONT_T = C.lv_font_t
type LV_ANIM_T = C.lv_anim_t
type LV_DISP_T = C.lv_disp_t
type LV_STYLE_SELECTOR_T = C.lv_style_selector_t
type LV_STYLE_VALUE_T = C.lv_style_value_t

type LV_OBJ_T = *C.lv_obj_t // 该变量因经常以指针使用，因此直接定义为指针
type LV_OBJ_CLASS_T = C.lv_obj_class_t

// ================ inferface
type CreateI interface {
	GetObj() unsafe.Pointer
}

type _empty struct{}

func (s _empty) GetObj() unsafe.Pointer {
	return unsafe.Pointer(nil)
}

var CREATE_NIL _empty = _empty{}

// ================ constant
const (
	LV_ANIM_ON              uint32  = C.LV_ANIM_ON
	LV_ANIM_OFF             uint32  = C.LV_ANIM_OFF
	LV_ANIM_REPEAT_INFINITE C_SHORT = C.LV_ANIM_REPEAT_INFINITE

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
	LV_ALIGN_DEFAULT      uint8 = C.LV_ALIGN_DEFAULT
	LV_ALIGN_TOP_LEFT     uint8 = C.LV_ALIGN_TOP_LEFT
	LV_ALIGN_TOP_MID      uint8 = C.LV_ALIGN_TOP_MID
	LV_ALIGN_TOP_RIGHT    uint8 = C.LV_ALIGN_TOP_RIGHT
	LV_ALIGN_BOTTOM_LEFT  uint8 = C.LV_ALIGN_BOTTOM_LEFT
	LV_ALIGN_BOTTOM_MID   uint8 = C.LV_ALIGN_BOTTOM_MID
	LV_ALIGN_BOTTOM_RIGHT uint8 = C.LV_ALIGN_BOTTOM_RIGHT
	LV_ALIGN_LEFT_MID     uint8 = C.LV_ALIGN_LEFT_MID
	LV_ALIGN_RIGHT_MID    uint8 = C.LV_ALIGN_RIGHT_MID
	LV_ALIGN_CENTER       uint8 = C.LV_ALIGN_CENTER

	LV_ALIGN_OUT_TOP_LEFT     uint8 = C.LV_ALIGN_OUT_TOP_LEFT
	LV_ALIGN_OUT_TOP_MID      uint8 = C.LV_ALIGN_OUT_TOP_MID
	LV_ALIGN_OUT_TOP_RIGHT    uint8 = C.LV_ALIGN_OUT_TOP_RIGHT
	LV_ALIGN_OUT_BOTTOM_LEFT  uint8 = C.LV_ALIGN_OUT_BOTTOM_LEFT
	LV_ALIGN_OUT_BOTTOM_MID   uint8 = C.LV_ALIGN_OUT_BOTTOM_MID
	LV_ALIGN_OUT_BOTTOM_RIGHT uint8 = C.LV_ALIGN_OUT_BOTTOM_RIGHT
	LV_ALIGN_OUT_LEFT_TOP     uint8 = C.LV_ALIGN_OUT_LEFT_TOP
	LV_ALIGN_OUT_LEFT_MID     uint8 = C.LV_ALIGN_OUT_LEFT_MID
	LV_ALIGN_OUT_LEFT_BOTTOM  uint8 = C.LV_ALIGN_OUT_LEFT_BOTTOM
	LV_ALIGN_OUT_RIGHT_TOP    uint8 = C.LV_ALIGN_OUT_RIGHT_TOP
	LV_ALIGN_OUT_RIGHT_MID    uint8 = C.LV_ALIGN_OUT_RIGHT_MID
	LV_ALIGN_OUT_RIGHT_BOTTOM uint8 = C.LV_ALIGN_OUT_RIGHT_BOTTOM

	LV_SIZE_CONTENT int16 = C.LV_SIZE_CONTENT
)

const (
	LV_GRAD_DIR_NONE uint8 = C.LV_GRAD_DIR_NONE
	LV_GRAD_DIR_VER  uint8 = C.LV_GRAD_DIR_VER
	LV_GRAD_DIR_HOR  uint8 = C.LV_GRAD_DIR_HOR

	LV_FLEX_ALIGN_START         uint8 = C.LV_FLEX_ALIGN_START
	LV_FLEX_ALIGN_END           uint8 = C.LV_FLEX_ALIGN_END
	LV_FLEX_ALIGN_CENTER        uint8 = C.LV_FLEX_ALIGN_CENTER
	LV_FLEX_ALIGN_SPACE_EVENLY  uint8 = C.LV_FLEX_ALIGN_SPACE_EVENLY
	LV_FLEX_ALIGN_SPACE_AROUND  uint8 = C.LV_FLEX_ALIGN_SPACE_AROUND
	LV_FLEX_ALIGN_SPACE_BETWEEN uint8 = C.LV_FLEX_ALIGN_SPACE_BETWEEN

	LV_FLEX_FLOW_ROW                 uint8 = C.LV_FLEX_FLOW_ROW
	LV_FLEX_FLOW_COLUMN              uint8 = C.LV_FLEX_FLOW_COLUMN
	LV_FLEX_FLOW_ROW_WRAP            uint8 = C.LV_FLEX_FLOW_ROW_WRAP
	LV_FLEX_FLOW_ROW_REVERSE         uint8 = C.LV_FLEX_FLOW_ROW_REVERSE
	LV_FLEX_FLOW_ROW_WRAP_REVERSE    uint8 = C.LV_FLEX_FLOW_ROW_WRAP_REVERSE
	LV_FLEX_FLOW_COLUMN_WRAP         uint8 = C.LV_FLEX_FLOW_COLUMN_WRAP
	LV_FLEX_FLOW_COLUMN_REVERSE      uint8 = C.LV_FLEX_FLOW_COLUMN_REVERSE
	LV_FLEX_FLOW_COLUMN_WRAP_REVERSE uint8 = C.LV_FLEX_FLOW_COLUMN_WRAP_REVERSE
)
