package lib

/*
#include "lv_init.h"
static const int _LV_SIZE_CONTENT = LV_SIZE_CONTENT;
*/
import "C"

type LvDispT C.lv_disp_t
type LvDispDrvT C.lv_disp_drv_t
type LvCoordT C.lv_coord_t
type LvColorT C.lv_color_t
type LvOpaT C.lv_opa_t
type LvTimerT C.lv_timer_t
type LvTimerCbT C.lv_timer_cb_t
type LvStyleT C.lv_style_t
type LvAlignT C.lv_align_t
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
type LvStylePropT C.lv_style_prop_t
type LvStyleValueT C.lv_style_value_t
type LvAnimTimelineT C.lv_anim_timeline_t
type LvAreaT C.lv_area_t
type LvObjT C.lv_obj_t
type LvFlexAlignT C.lv_flex_align_t
type LvGridAlignT C.lv_grid_align_t
type LvChartTypeT C.lv_chart_type_t
type LvChartAxisT C.lv_chart_axis_t
type LvChartUpdateModeT C.lv_chart_update_mode_t
type LvChartSeriesT C.lv_chart_series_t
type LvChartCursorT C.lv_chart_cursor_t
type LvPointT C.lv_point_t
type LvImgbtnStateT C.lv_imgbtn_state_t
type LvSpanT C.lv_span_t
type LvSpanOverflowT C.lv_span_overflow_t
type LvSpanModeT C.lv_span_mode_t
type LvColorHsvT C.lv_color_hsv_t
type LvColorwheelModeT C.lv_colorwheel_mode_t
type LvMenuModeHeaderT C.lv_menu_mode_header_t
type LvMenuModeRootBackBtnT C.lv_menu_mode_root_back_btn_t
type LvMeterScaleT C.lv_meter_scale_t
type LvMeterIndicatorT C.lv_meter_indicator_t
type LvKeyboardModeT C.lv_keyboard_mode_t
type LvImgDscT C.lv_img_dsc_t
type LvImgDecoderT C.lv_img_decoder_t
type LvImgDecoderInfoFT C.lv_img_decoder_info_f_t
type LvImgDecoderOpenFT C.lv_img_decoder_open_f_t
type LvImgDecoderReadLineFT C.lv_img_decoder_read_line_f_t
type LvImgDecoderCloseFT C.lv_img_decoder_close_f_t
type LvImgSizeModeT C.lv_img_size_mode_t
type LvArcModeT C.lv_arc_mode_t
type LvBarModeT C.lv_bar_mode_t
type LvLabelLongModeT C.lv_label_long_mode_t
type LvRollerModeT C.lv_roller_mode_t
type LvBtnmatrixCtrlT C.lv_btnmatrix_ctrl_t
type LvStyleSelectorT C.lv_style_selector_t
type LvGroupT C.lv_group_t
type LvGroupFocusCbT C.lv_group_focus_cb_t
type LvGroupEdgeCbT C.lv_group_edge_cb_t
type LvGroupRefocusPolicyT C.lv_group_refocus_policy_t
type LvScrollbarModeT C.lv_scrollbar_mode_t
type LvDirT C.lv_dir_t
type LvScrollSnapT C.lv_scroll_snap_t
type LvThemeT C.lv_theme_t
type LvThemeApplyCbT C.lv_theme_apply_cb_t
type LvIndevT C.lv_indev_t
type LvEventT C.lv_event_t
type LvCoverResT C.lv_cover_res_t
type LvAnimEnableT C.lv_anim_enable_t
type LvFlexFlowT C.lv_flex_flow_t

// type LvMonkeyT C.lv_monkey_t
// type LvDispRotT C.lv_disp_rot_t
// type LvImgCfT C.lv_img_cf_t
// type LvPinyinDictT C.lv_pinyin_dict_t
// type LvImePinyinModeT C.lv_ime_pinyin_mode_t

// type LvFfmpegPlayerCmdT C.lv_ffmpeg_player_cmd_t
// type LvRlottieCtrlT C.lv_rlottie_ctrl_t

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
