package set

import (
	"lvgl-go/src/lib"
	types "lvgl-go/src/types"
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

func (setter *Img) OffsetX(x types.LvCoordT) *Img {
	C.lv_img_set_offset_x(setter.CObj, C.lv_coord_t(x))

	return setter
}
func (setter *Img) OffsetY(y types.LvCoordT) *Img {
	C.lv_img_set_offset_y(setter.CObj, C.lv_coord_t(y))

	return setter
}
func (setter *Img) Angle(angle int16) *Img {
	C.lv_img_set_angle(setter.CObj, C.short(angle))

	return setter
}
func (setter *Img) Pivot(x types.LvCoordT, y types.LvCoordT) *Img {
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
func (setter *Img) SizeMode(mode types.LvImgSizeModeT) *Img {
	C.lv_img_set_size_mode(setter.CObj, C.lv_img_size_mode_t(mode))

	return setter
}

func (setter *Img) CacheSize(new_slot_num uint16) *Img {
	C.lv_img_cache_set_size(C.ushort(new_slot_num))

	return setter
}

func (setter *Img) BufPxColor(dsc *types.LvImgDscT, x types.LvCoordT, y types.LvCoordT, c types.LvColorT) *Img {
	C.lv_img_buf_set_px_color((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter *Img) BufPxAlpha(dsc *types.LvImgDscT, x types.LvCoordT, y types.LvCoordT, opa types.LvOpaT) *Img {
	C.lv_img_buf_set_px_alpha((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter *Img) BufPalette(dsc *types.LvImgDscT, id uint8, c types.LvColorT) *Img {
	C.lv_img_buf_set_palette((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.uint8_t(id), C.lv_color_t(c))

	return setter
}

func (setter *Img) DecoderInfoCb(decoder *types.LvImgDecoderT, info_cb types.LvImgDecoderInfoFT) *Img {
	C.lv_img_decoder_set_info_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_info_f_t(info_cb))

	return setter
}
func (setter *Img) DecoderOpenCb(decoder *types.LvImgDecoderT, open_cb types.LvImgDecoderOpenFT) *Img {
	C.lv_img_decoder_set_open_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_open_f_t(open_cb))

	return setter
}
func (setter *Img) DecoderReadLineCb(decoder *types.LvImgDecoderT, read_line_cb types.LvImgDecoderReadLineFT) *Img {
	C.lv_img_decoder_set_read_line_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_read_line_f_t(read_line_cb))

	return setter
}
func (setter *Img) DecoderCloseCb(decoder *types.LvImgDecoderT, close_cb types.LvImgDecoderCloseFT) *Img {
	C.lv_img_decoder_set_close_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_close_f_t(close_cb))

	return setter
}
