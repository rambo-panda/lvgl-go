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

type SetArea set

func CreateArea(o *types.LvObjT) SetArea {
	return SetArea{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetArea) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetArea) SetWidth(area_p *types.LvAreaT, w types.LvCoordT) SetArea {
	C.lv_area_set_width((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(w))

	return setter
}
func (setter SetArea) SetHeight(area_p *types.LvAreaT, h types.LvCoordT) SetArea {
	C.lv_area_set_height((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(h))

	return setter
}
