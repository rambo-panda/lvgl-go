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

func FlexFlowForStyle(style *lib.LvStyleT, value lib.LvFlexFlowT) {
	C.lv_style_set_flex_flow((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_flow_t(value))

}
func FlexMainPlaceForStyle(style *lib.LvStyleT, value lib.LvFlexAlignT) {
	C.lv_style_set_flex_main_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

}
func FlexCrossPlaceForStyle(style *lib.LvStyleT, value lib.LvFlexAlignT) {
	C.lv_style_set_flex_cross_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

}
func FlexTrackPlaceForStyle(style *lib.LvStyleT, value lib.LvFlexAlignT) {
	C.lv_style_set_flex_track_place((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_flex_align_t(value))

}
func FlexGrowForStyle(style *lib.LvStyleT, value uint8) {
	C.lv_style_set_flex_grow((*C.lv_style_t)(unsafe.Pointer(style)), C.uint8_t(value))

}
func GridRowAlignForStyle(style *lib.LvStyleT, value lib.LvGridAlignT) {
	C.lv_style_set_grid_row_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

}
func GridColumnAlignForStyle(style *lib.LvStyleT, value lib.LvGridAlignT) {
	C.lv_style_set_grid_column_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_grid_align_t(value))

}
func GridCellColumnPosForStyle(style *lib.LvStyleT, value lib.LvCoordT) {
	C.lv_style_set_grid_cell_column_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

}
func GridCellColumnSpanForStyle(style *lib.LvStyleT, value lib.LvCoordT) {
	C.lv_style_set_grid_cell_column_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

}
func GridCellRowPosForStyle(style *lib.LvStyleT, value lib.LvCoordT) {
	C.lv_style_set_grid_cell_row_pos((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

}
func GridCellRowSpanForStyle(style *lib.LvStyleT, value lib.LvCoordT) {
	C.lv_style_set_grid_cell_row_span((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

}
func GridCellXAlignForStyle(style *lib.LvStyleT, value lib.LvCoordT) {
	C.lv_style_set_grid_cell_x_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

}
func GridCellYAlignForStyle(style *lib.LvStyleT, value lib.LvCoordT) {
	C.lv_style_set_grid_cell_y_align((*C.lv_style_t)(unsafe.Pointer(style)), C.lv_coord_t(value))

}
