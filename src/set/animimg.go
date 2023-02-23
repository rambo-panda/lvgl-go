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

type SetAnimimg set

func CreateAnimimg(o *lib.LvObjT) SetAnimimg {
	return SetAnimimg{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetAnimimg) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetAnimimg) SetDuration(duration uint32) SetAnimimg {
	C.lv_animimg_set_duration(setter.CStructLvObjT, C.uint(duration))

	return setter
}
func (setter SetAnimimg) SetRepeatCount(count uint16) SetAnimimg {
	C.lv_animimg_set_repeat_count(setter.CStructLvObjT, C.ushort(count))

	return setter
}
