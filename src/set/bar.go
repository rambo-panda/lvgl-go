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

type SetBar set

func CreateBar(o *types.LvObjT) SetBar {
	return SetBar{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetBar) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetBar) SetValue(value int32, anim types.LvAnimEnableT) SetBar {
	C.lv_bar_set_value(setter.CStructLvObjT, C.int(value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter SetBar) SetStartValue(start_value int32, anim types.LvAnimEnableT) SetBar {
	C.lv_bar_set_start_value(setter.CStructLvObjT, C.int(start_value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter SetBar) SetRange(min int32, max int32) SetBar {
	C.lv_bar_set_range(setter.CStructLvObjT, C.int(min), C.int(max))

	return setter
}
func (setter SetBar) SetMode(mode types.LvBarModeT) SetBar {
	C.lv_bar_set_mode(setter.CStructLvObjT, C.lv_bar_mode_t(mode))

	return setter
}
