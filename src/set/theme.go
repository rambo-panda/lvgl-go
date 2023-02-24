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

type Theme set

func CreateTheme(o *types.LvObjT) Theme {
	return Theme{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Theme) Parent(new_theme *types.LvThemeT, parent *types.LvThemeT) Theme {
	C.lv_theme_set_parent((*C.lv_theme_t)(unsafe.Pointer(new_theme)), (*C.lv_theme_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter Theme) ApplyCb(theme *types.LvThemeT, apply_cb types.LvThemeApplyCbT) Theme {
	C.lv_theme_set_apply_cb((*C.lv_theme_t)(unsafe.Pointer(theme)), C.lv_theme_apply_cb_t(apply_cb))

	return setter
}
