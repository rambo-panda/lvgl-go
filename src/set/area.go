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

type Area set[CObjT]

func CreateArea(o CObjT) Area {
	return Area{
		CObj: o,
	}
}

func (setter Area) Width(area_p *types.LvAreaT, w lib.LV_COLOR_T) Area {
	C.lv_area_set_width((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(w))

	return setter
}
func (setter Area) Height(area_p *types.LvAreaT, h lib.LV_COLOR_T) Area {
	C.lv_area_set_height((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(h))

	return setter
}
