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

type SetArea set

func CreateArea(o *lib.LvObjT) SetArea {
	return SetArea{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetArea) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetArea) SetWidth(area_p *lib.LvAreaT, w lib.LvCoordT) SetArea {
	C.lv_area_set_width((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(w))

	return setter
}
func (setter SetArea) SetHeight(area_p *lib.LvAreaT, h lib.LvCoordT) SetArea {
	C.lv_area_set_height((*C.lv_area_t)(unsafe.Pointer(area_p)), C.lv_coord_t(h))

	return setter
}
