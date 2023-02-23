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

type SetChart set

func (setter SetChart) SetType(_type lib.LvChartTypeT) SetChart {
	C.lv_chart_set_type(setter.cObj, C.lv_chart_type_t(_type))

	return setter
}
func (setter SetChart) SetPointCount(cnt uint16) SetChart {
	C.lv_chart_set_point_count(setter.cObj, C.ushort(cnt))

	return setter
}
func (setter SetChart) SetRange(axis lib.LvChartAxisT, min lib.LvCoordT, max lib.LvCoordT) SetChart {
	C.lv_chart_set_range(setter.cObj, C.lv_chart_axis_t(axis), C.lv_coord_t(min), C.lv_coord_t(max))

	return setter
}
func (setter SetChart) SetUpdateMode(update_mode lib.LvChartUpdateModeT) SetChart {
	C.lv_chart_set_update_mode(setter.cObj, C.lv_chart_update_mode_t(update_mode))

	return setter
}
func (setter SetChart) SetDivLineCount(hdiv uint8, vdiv uint8) SetChart {
	C.lv_chart_set_div_line_count(setter.cObj, C.uint8_t(hdiv), C.uint8_t(vdiv))

	return setter
}
func (setter SetChart) SetZoomX(zoom_x uint16) SetChart {
	C.lv_chart_set_zoom_x(setter.cObj, C.ushort(zoom_x))

	return setter
}
func (setter SetChart) SetZoomY(zoom_y uint16) SetChart {
	C.lv_chart_set_zoom_y(setter.cObj, C.ushort(zoom_y))

	return setter
}
func (setter SetChart) SetAxisTick(axis lib.LvChartAxisT, major_len lib.LvCoordT, minor_len lib.LvCoordT, major_cnt lib.LvCoordT, minor_cnt lib.LvCoordT, label_en bool, draw_size lib.LvCoordT) SetChart {
	C.lv_chart_set_axis_tick(setter.cObj, C.lv_chart_axis_t(axis), C.lv_coord_t(major_len), C.lv_coord_t(minor_len), C.lv_coord_t(major_cnt), C.lv_coord_t(minor_cnt), C.bool(label_en), C.lv_coord_t(draw_size))

	return setter
}
func (setter SetChart) SetSeriesColor(series *lib.LvChartSeriesT, color lib.LvColorT) SetChart {
	C.lv_chart_set_series_color(setter.cObj, (*C.lv_chart_series_t)(unsafe.Pointer(series)), C.lv_color_t(color))

	return setter
}
func (setter SetChart) SetXStartPoint(ser *lib.LvChartSeriesT, id uint16) SetChart {
	C.lv_chart_set_x_start_point(setter.cObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(id))

	return setter
}
func (setter SetChart) SetCursorPos(cursor *lib.LvChartCursorT, pos *lib.LvPointT) SetChart {
	C.lv_chart_set_cursor_pos(setter.cObj, (*C.lv_chart_cursor_t)(unsafe.Pointer(cursor)), (*C.lv_point_t)(unsafe.Pointer(pos)))

	return setter
}
func (setter SetChart) SetCursorPoint(cursor *lib.LvChartCursorT, ser *lib.LvChartSeriesT, point_id uint16) SetChart {
	C.lv_chart_set_cursor_point(setter.cObj, (*C.lv_chart_cursor_t)(unsafe.Pointer(cursor)), (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(point_id))

	return setter
}
func (setter SetChart) SetAllValue(ser *lib.LvChartSeriesT, value lib.LvCoordT) SetChart {
	C.lv_chart_set_all_value(setter.cObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.lv_coord_t(value))

	return setter
}
func (setter SetChart) SetNextValue(ser *lib.LvChartSeriesT, value lib.LvCoordT) SetChart {
	C.lv_chart_set_next_value(setter.cObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.lv_coord_t(value))

	return setter
}
func (setter SetChart) SetValueById(ser *lib.LvChartSeriesT, id uint16, value lib.LvCoordT) SetChart {
	C.lv_chart_set_value_by_id(setter.cObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(id), C.lv_coord_t(value))

	return setter
}
