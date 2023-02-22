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

func SetSizeForImg(new_slot_num uint16) {
	C.lv_img_cache_set_size(C.ushort(new_slot_num))

}
func SetPxColorForImg(dsc *lib.LvImgDscT, x lib.LvCoordT, y lib.LvCoordT, c lib.LvColorT) {
	C.lv_img_buf_set_px_color((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

}
func SetPxAlphaForImg(dsc *lib.LvImgDscT, x lib.LvCoordT, y lib.LvCoordT, opa lib.LvOpaT) {
	C.lv_img_buf_set_px_alpha((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

}
func SetPaletteForImg(dsc *lib.LvImgDscT, id uint8, c lib.LvColorT) {
	C.lv_img_buf_set_palette((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.uint8_t(id), C.lv_color_t(c))

}
func SetInfoCbForImg(decoder *lib.LvImgDecoderT, info_cb lib.LvImgDecoderInfoFT) {
	C.lv_img_decoder_set_info_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_info_f_t(info_cb))

}
func SetOpenCbForImg(decoder *lib.LvImgDecoderT, open_cb lib.LvImgDecoderOpenFT) {
	C.lv_img_decoder_set_open_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_open_f_t(open_cb))

}
func SetReadLineCbForImg(decoder *lib.LvImgDecoderT, read_line_cb lib.LvImgDecoderReadLineFT) {
	C.lv_img_decoder_set_read_line_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_read_line_f_t(read_line_cb))

}
func SetCloseCbForImg(decoder *lib.LvImgDecoderT, close_cb lib.LvImgDecoderCloseFT) {
	C.lv_img_decoder_set_close_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_close_f_t(close_cb))

}
func SrcForImg(obj *lib.LvObjT, src any) {
	C.lv_img_set_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), unsafe.Pointer(&src))

}
func OffsetXForImg(obj *lib.LvObjT, x lib.LvCoordT) {
	C.lv_img_set_offset_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x))

}
func OffsetYForImg(obj *lib.LvObjT, y lib.LvCoordT) {
	C.lv_img_set_offset_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(y))

}
func AngleForImg(obj *lib.LvObjT, angle int16) {
	C.lv_img_set_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.short(angle))

}
func PivotForImg(obj *lib.LvObjT, x lib.LvCoordT, y lib.LvCoordT) {
	C.lv_img_set_pivot((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(x), C.lv_coord_t(y))

}
func ZoomForImg(obj *lib.LvObjT, zoom uint16) {
	C.lv_img_set_zoom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.ushort(zoom))

}
func AntialiasForImg(obj *lib.LvObjT, antialias bool) {
	C.lv_img_set_antialias((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.bool(antialias))

}
func SizeModeForImg(obj *lib.LvObjT, mode lib.LvImgSizeModeT) {
	C.lv_img_set_size_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_img_size_mode_t(mode))

}
