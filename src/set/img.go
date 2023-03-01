package set

import (
	"lvgl-go/src/lib"
	lib.LV_FONT_T"lvgl-go/src/lib.LV_FONT_T
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Img set[CObjT]

func CreateImg(o CObjT) *Img {
	_o := Img{
		CObj: o,
	}

	return &_o
}

func (setter *Img) SrcOrigin(src unsafe.Pointer) *Img {
	C.lv_img_set_src(setter.CObj, src)

	return setter
}
func (setter *Img) Src(src string) *Img {
	return setter.SrcOrigin(lib.Go2CString(src))
}

func (setter *Img) OffsetX(x lib.LV_COLOR_T) *Img {
	C.lv_img_set_offset_x(setter.CObj, C.lv_coord_t(x))

	return setter
}
func (setter *Img) OffsetY(y lib.LV_COLOR_T) *Img {
	C.lv_img_set_offset_y(setter.CObj, C.lv_coord_t(y))

	return setter
}
func (setter *Img) Angle(angle int16) *Img {
	C.lv_img_set_angle(setter.CObj, C.short(angle))

	return setter
}
func (setter *Img) Pivot(x lib.LV_COLOR_T, y lib.LV_COLOR_T) *Img {
	C.lv_img_set_pivot(setter.CObj, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter *Img) Zoom(zoom uint16) *Img {
	C.lv_img_set_zoom(setter.CObj, C.ushort(zoom))

	return setter
}
func (setter *Img) Antialias(antialias bool) *Img {
	C.lv_img_set_antialias(setter.CObj, C.bool(antialias))

	return setter
}
func (setter *Img) SizeMode(mode lib.LV_FONT_TLvImgSizeModeT) *Img {
	C.lv_img_set_size_mode(setter.CObj, C.lv_img_size_mode_t(mode))

	return setter
}

func (setter *Img) CacheSize(new_slot_num uint16) *Img {
	C.lv_img_cache_set_size(C.ushort(new_slot_num))

	return setter
}

func (setter *Img) BufPxColor(dsc *lib.LV_FONT_TLvImgDscT, x lib.LV_COLOR_T, y lib.LV_COLOR_T, c lib.LV_COLOR_T) *Img {
	C.lv_img_buf_set_px_color((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter *Img) BufPxAlpha(dsc *lib.LV_FONT_TLvImgDscT, x lib.LV_COLOR_T, y lib.LV_COLOR_T, opa lib.LV_OPA_T) *Img {
	C.lv_img_buf_set_px_alpha((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter *Img) BufPalette(dsc *lib.LV_FONT_TLvImgDscT, id uint8, c lib.LV_COLOR_T) *Img {
	C.lv_img_buf_set_palette((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.uint8_t(id), C.lv_color_t(c))

	return setter
}

func (setter *Img) DecoderInfoCb(decoder *lib.LV_FONT_TLvImgDecoderT, info_cb lib.LV_FONT_TLvImgDecoderInfoFT) *Img {
	C.lv_img_decoder_set_info_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_info_f_t(info_cb))

	return setter
}
func (setter *Img) DecoderOpenCb(decoder *lib.LV_FONT_TLvImgDecoderT, open_cb lib.LV_FONT_TLvImgDecoderOpenFT) *Img {
	C.lv_img_decoder_set_open_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_open_f_t(open_cb))

	return setter
}
func (setter *Img) DecoderReadLineCb(decoder *lib.LV_FONT_TLvImgDecoderT, read_line_cb lib.LV_FONT_TLvImgDecoderReadLineFT) *Img {
	C.lv_img_decoder_set_read_line_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_read_line_f_t(read_line_cb))

	return setter
}
func (setter *Img) DecoderCloseCb(decoder *lib.LV_FONT_TLvImgDecoderT, close_cb lib.LV_FONT_TLvImgDecoderCloseFT) *Img {
	C.lv_img_decoder_set_close_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_close_f_t(close_cb))

	return setter
}
