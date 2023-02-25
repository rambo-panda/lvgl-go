package set

import (
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Animimg set

func CreateAnimimg(o CObjT) Animimg {
	return Animimg{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Animimg) Duration(duration uint32) Animimg {
	C.lv_animimg_set_duration(setter.CStructLvObjT, C.uint(duration))

	return setter
}
func (setter Animimg) RepeatCount(count uint16) Animimg {
	C.lv_animimg_set_repeat_count(setter.CStructLvObjT, C.ushort(count))

	return setter
}
