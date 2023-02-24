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

type Bar set

func CreateBar(o *types.LvObjT) Bar {
	return Bar{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Bar) Value(value int32, anim types.LvAnimEnableT) Bar {
	C.lv_bar_set_value(setter.CStructLvObjT, C.int(value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter Bar) StartValue(start_value int32, anim types.LvAnimEnableT) Bar {
	C.lv_bar_set_start_value(setter.CStructLvObjT, C.int(start_value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter Bar) Range(min int32, max int32) Bar {
	C.lv_bar_set_range(setter.CStructLvObjT, C.int(min), C.int(max))

	return setter
}
func (setter Bar) Mode(mode types.LvBarModeT) Bar {
	C.lv_bar_set_mode(setter.CStructLvObjT, C.lv_bar_mode_t(mode))

	return setter
}
