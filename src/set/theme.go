package set

import (
	"lvgl-go/src/lib"
	"unsafe"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Theme set[CObjT]

func CreateTheme(o CObjT) Theme {
	return Theme{
		CObj: o,
	}
}

func (setter Theme) Parent(new_theme *lib.LV_FONT_TLvThemeT, parent *lib.LV_FONT_TLvThemeT) Theme {
	C.lv_theme_set_parent((*C.lv_theme_t)(unsafe.Pointer(new_theme)), (*C.lv_theme_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter Theme) ApplyCb(theme *lib.LV_FONT_TLvThemeT, apply_cb lib.LV_FONT_TLvThemeApplyCbT) Theme {
	C.lv_theme_set_apply_cb((*C.lv_theme_t)(unsafe.Pointer(theme)), C.lv_theme_apply_cb_t(apply_cb))

	return setter
}
