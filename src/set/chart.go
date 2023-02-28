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

type Chart set[CObjT]

func CreateChart(o CObjT) Chart {
	return Chart{
		CObj: o,
	}
}

func (setter Chart) Type(_type types.LvChartTypeT) Chart {
	C.lv_chart_set_type(setter.CObj, C.lv_chart_type_t(_type))

	return setter
}
func (setter Chart) PointCount(cnt uint16) Chart {
	C.lv_chart_set_point_count(setter.CObj, C.ushort(cnt))

	return setter
}
func (setter Chart) Range(axis types.LvChartAxisT, min types.LvCoordT, max types.LvCoordT) Chart {
	C.lv_chart_set_range(setter.CObj, C.lv_chart_axis_t(axis), C.lv_coord_t(min), C.lv_coord_t(max))

	return setter
}
func (setter Chart) UpdateMode(update_mode types.LvChartUpdateModeT) Chart {
	C.lv_chart_set_update_mode(setter.CObj, C.lv_chart_update_mode_t(update_mode))

	return setter
}
func (setter Chart) DivLineCount(hdiv uint8, vdiv uint8) Chart {
	C.lv_chart_set_div_line_count(setter.CObj, C.uint8_t(hdiv), C.uint8_t(vdiv))

	return setter
}
func (setter Chart) ZoomX(zoom_x uint16) Chart {
	C.lv_chart_set_zoom_x(setter.CObj, C.ushort(zoom_x))

	return setter
}
func (setter Chart) ZoomY(zoom_y uint16) Chart {
	C.lv_chart_set_zoom_y(setter.CObj, C.ushort(zoom_y))

	return setter
}
func (setter Chart) AxisTick(axis types.LvChartAxisT, major_len types.LvCoordT, minor_len types.LvCoordT, major_cnt types.LvCoordT, minor_cnt types.LvCoordT, label_en bool, draw_size types.LvCoordT) Chart {
	C.lv_chart_set_axis_tick(setter.CObj, C.lv_chart_axis_t(axis), C.lv_coord_t(major_len), C.lv_coord_t(minor_len), C.lv_coord_t(major_cnt), C.lv_coord_t(minor_cnt), C.bool(label_en), C.lv_coord_t(draw_size))

	return setter
}
func (setter Chart) SeriesColor(series *types.LvChartSeriesT, color types.LvColorT) Chart {
	C.lv_chart_set_series_color(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(series)), C.lv_color_t(color))

	return setter
}
func (setter Chart) XStartPoint(ser *types.LvChartSeriesT, id uint16) Chart {
	C.lv_chart_set_x_start_point(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(id))

	return setter
}
func (setter Chart) CursorPos(cursor *types.LvChartCursorT, pos *types.LvPointT) Chart {
	C.lv_chart_set_cursor_pos(setter.CObj, (*C.lv_chart_cursor_t)(unsafe.Pointer(cursor)), (*C.lv_point_t)(unsafe.Pointer(pos)))

	return setter
}
func (setter Chart) CursorPoint(cursor *types.LvChartCursorT, ser *types.LvChartSeriesT, point_id uint16) Chart {
	C.lv_chart_set_cursor_point(setter.CObj, (*C.lv_chart_cursor_t)(unsafe.Pointer(cursor)), (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(point_id))

	return setter
}
func (setter Chart) AllValue(ser *types.LvChartSeriesT, value types.LvCoordT) Chart {
	C.lv_chart_set_all_value(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.lv_coord_t(value))

	return setter
}
func (setter Chart) NextValue(ser *types.LvChartSeriesT, value types.LvCoordT) Chart {
	C.lv_chart_set_next_value(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.lv_coord_t(value))

	return setter
}
func (setter Chart) NextValue2(ser *types.LvChartSeriesT, x_value types.LvCoordT, y_value types.LvCoordT) Chart {
	C.lv_chart_set_next_value2(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.lv_coord_t(x_value), C.lv_coord_t(y_value))

	return setter
}
func (setter Chart) ValueById(ser *types.LvChartSeriesT, id uint16, value types.LvCoordT) Chart {
	C.lv_chart_set_value_by_id(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(id), C.lv_coord_t(value))

	return setter
}
func (setter Chart) ValueById2(ser *types.LvChartSeriesT, id uint16, x_value types.LvCoordT, y_value types.LvCoordT) Chart {
	C.lv_chart_set_value_by_id2(setter.CObj, (*C.lv_chart_series_t)(unsafe.Pointer(ser)), C.ushort(id), C.lv_coord_t(x_value), C.lv_coord_t(y_value))

	return setter
}
