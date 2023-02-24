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

type SetAnimimg set

func CreateAnimimg(o *types.LvObjT) SetAnimimg {
	return SetAnimimg{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetAnimimg) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetAnimimg) SetDuration(duration uint32) SetAnimimg {
	C.lv_animimg_set_duration(setter.CStructLvObjT, C.uint(duration))

	return setter
}
func (setter SetAnimimg) SetRepeatCount(count uint16) SetAnimimg {
	C.lv_animimg_set_repeat_count(setter.CStructLvObjT, C.ushort(count))

	return setter
}
