package set

import types "lvgl-go/src/types"

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type Img set

func CreateImg(o *types.LvObjT) Img {
	return Img{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Img) Src(src any) Img {
	C.lv_img_set_src(setter.CStructLvObjT, unsafe.Pointer(&src))

	return setter
}
func (setter Img) OffsetX(x types.LvCoordT) Img {
	C.lv_img_set_offset_x(setter.CStructLvObjT, C.lv_coord_t(x))

	return setter
}
func (setter Img) OffsetY(y types.LvCoordT) Img {
	C.lv_img_set_offset_y(setter.CStructLvObjT, C.lv_coord_t(y))

	return setter
}
func (setter Img) Angle(angle int16) Img {
	C.lv_img_set_angle(setter.CStructLvObjT, C.short(angle))

	return setter
}
func (setter Img) Pivot(x types.LvCoordT, y types.LvCoordT) Img {
	C.lv_img_set_pivot(setter.CStructLvObjT, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter Img) Zoom(zoom uint16) Img {
	C.lv_img_set_zoom(setter.CStructLvObjT, C.ushort(zoom))

	return setter
}
func (setter Img) Antialias(antialias bool) Img {
	C.lv_img_set_antialias(setter.CStructLvObjT, C.bool(antialias))

	return setter
}
func (setter Img) SizeMode(mode types.LvImgSizeModeT) Img {
	C.lv_img_set_size_mode(setter.CStructLvObjT, C.lv_img_size_mode_t(mode))

	return setter
}

func (setter Img) CacheSize(new_slot_num uint16) Img {
	C.lv_img_cache_set_size(C.ushort(new_slot_num))

	return setter
}

func (setter Img) BufPxColor(dsc *types.LvImgDscT, x types.LvCoordT, y types.LvCoordT, c types.LvColorT) Img {
	C.lv_img_buf_set_px_color((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter Img) BufPxAlpha(dsc *types.LvImgDscT, x types.LvCoordT, y types.LvCoordT, opa types.LvOpaT) Img {
	C.lv_img_buf_set_px_alpha((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter Img) BufPalette(dsc *types.LvImgDscT, id uint8, c types.LvColorT) Img {
	C.lv_img_buf_set_palette((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.uint8_t(id), C.lv_color_t(c))

	return setter
}

func (setter Img) DecoderInfoCb(decoder *types.LvImgDecoderT, info_cb types.LvImgDecoderInfoFT) Img {
	C.lv_img_decoder_set_info_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_info_f_t(info_cb))

	return setter
}
func (setter Img) DecoderOpenCb(decoder *types.LvImgDecoderT, open_cb types.LvImgDecoderOpenFT) Img {
	C.lv_img_decoder_set_open_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_open_f_t(open_cb))

	return setter
}
func (setter Img) DecoderReadLineCb(decoder *types.LvImgDecoderT, read_line_cb types.LvImgDecoderReadLineFT) Img {
	C.lv_img_decoder_set_read_line_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_read_line_f_t(read_line_cb))

	return setter
}
func (setter Img) DecoderCloseCb(decoder *types.LvImgDecoderT, close_cb types.LvImgDecoderCloseFT) Img {
	C.lv_img_decoder_set_close_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_close_f_t(close_cb))

	return setter
}
