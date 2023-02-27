package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/lib"
	"unsafe"
)

type Img get

func CreateImg(o any) Img {
	return Img{
		CObj: ((CObjT)(o.(unsafe.Pointer))),
	}
}

func (getter Img) SrcOrigin() unsafe.Pointer {
	src := C.lv_img_get_src(getter.CObj)

	return src
}

func (getter Img) Src() string {
	return lib.C2GoString(getter.SrcOrigin())
}

// const void * lv_img_get_src(lv_obj_t * obj);
// lv_coord_t lv_img_get_offset_x(lv_obj_t * obj);
// lv_coord_t lv_img_get_offset_y(lv_obj_t * obj);
// uint16_t lv_img_get_angle(lv_obj_t * obj);
// void lv_img_get_pivot(lv_obj_t * obj, lv_point_t * pivot);
// uint16_t lv_img_get_zoom(lv_obj_t * obj);
// bool lv_img_get_antialias(lv_obj_t * obj);
// lv_img_size_mode_t lv_img_get_size_mode(lv_obj_t * obj);
