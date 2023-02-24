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

type Area set

func CreateArea(o *types.LvObjT) Area {
	return Area{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Area) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter Area) Width(area_p *types.LvAreaT, w types.LvCoordT) Area {
	C.lv_area_set_width((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(w))

	return setter
}
func (setter Area) Height(area_p *types.LvAreaT, h types.LvCoordT) Area {
	C.lv_area_set_height((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(h))

	return setter
}
