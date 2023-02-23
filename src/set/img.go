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

type CacheImg set

func (setter CacheImg) CacheSetSize(new_slot_num uint16) CacheImg {
	C.lv_img_cache_set_size(C.ushort(new_slot_num))

	return setter
}

type BufImg set

func (setter BufImg) BufSetPxColor(dsc *lib.LvImgDscT, x lib.LvCoordT, y lib.LvCoordT, c lib.LvColorT) BufImg {
	C.lv_img_buf_set_px_color((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_color_t(c))

	return setter
}
func (setter BufImg) BufSetPxAlpha(dsc *lib.LvImgDscT, x lib.LvCoordT, y lib.LvCoordT, opa lib.LvOpaT) BufImg {
	C.lv_img_buf_set_px_alpha((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.lv_coord_t(x), C.lv_coord_t(y), C.lv_opa_t(opa))

	return setter
}
func (setter BufImg) BufSetPalette(dsc *lib.LvImgDscT, id uint8, c lib.LvColorT) BufImg {
	C.lv_img_buf_set_palette((*C.lv_img_dsc_t)(unsafe.Pointer(dsc)), C.uint8_t(id), C.lv_color_t(c))

	return setter
}

type DecoderImg set

func (setter DecoderImg) DecoderSetInfoCb(decoder *lib.LvImgDecoderT, info_cb lib.LvImgDecoderInfoFT) DecoderImg {
	C.lv_img_decoder_set_info_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_info_f_t(info_cb))

	return setter
}
func (setter DecoderImg) DecoderSetOpenCb(decoder *lib.LvImgDecoderT, open_cb lib.LvImgDecoderOpenFT) DecoderImg {
	C.lv_img_decoder_set_open_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_open_f_t(open_cb))

	return setter
}
func (setter DecoderImg) DecoderSetReadLineCb(decoder *lib.LvImgDecoderT, read_line_cb lib.LvImgDecoderReadLineFT) DecoderImg {
	C.lv_img_decoder_set_read_line_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_read_line_f_t(read_line_cb))

	return setter
}
func (setter DecoderImg) DecoderSetCloseCb(decoder *lib.LvImgDecoderT, close_cb lib.LvImgDecoderCloseFT) DecoderImg {
	C.lv_img_decoder_set_close_cb((*C.lv_img_decoder_t)(unsafe.Pointer(decoder)), C.lv_img_decoder_close_f_t(close_cb))

	return setter
}

type SetImg set

func (setter SetImg) SetSrc(src any) SetImg {
	C.lv_img_set_src(setter.cObj, unsafe.Pointer(&src))

	return setter
}
func (setter SetImg) SetOffsetX(x lib.LvCoordT) SetImg {
	C.lv_img_set_offset_x(setter.cObj, C.lv_coord_t(x))

	return setter
}
func (setter SetImg) SetOffsetY(y lib.LvCoordT) SetImg {
	C.lv_img_set_offset_y(setter.cObj, C.lv_coord_t(y))

	return setter
}
func (setter SetImg) SetAngle(angle int16) SetImg {
	C.lv_img_set_angle(setter.cObj, C.short(angle))

	return setter
}
func (setter SetImg) SetPivot(x lib.LvCoordT, y lib.LvCoordT) SetImg {
	C.lv_img_set_pivot(setter.cObj, C.lv_coord_t(x), C.lv_coord_t(y))

	return setter
}
func (setter SetImg) SetZoom(zoom uint16) SetImg {
	C.lv_img_set_zoom(setter.cObj, C.ushort(zoom))

	return setter
}
func (setter SetImg) SetAntialias(antialias bool) SetImg {
	C.lv_img_set_antialias(setter.cObj, C.bool(antialias))

	return setter
}
func (setter SetImg) SetSizeMode(mode lib.LvImgSizeModeT) SetImg {
	C.lv_img_set_size_mode(setter.cObj, C.lv_img_size_mode_t(mode))

	return setter
}
