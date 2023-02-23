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

type SetTheme set

func CreateTheme(o *lib.LvObjT) SetTheme {
	return SetTheme{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter SetTheme) GetObj() *lib.LvObjT {
	return (*lib.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter SetTheme) SetParent(new_theme *lib.LvThemeT, parent *lib.LvThemeT) SetTheme {
	C.lv_theme_set_parent((*C.lv_theme_t)(unsafe.Pointer(new_theme)), (*C.lv_theme_t)(unsafe.Pointer(parent)))

	return setter
}
func (setter SetTheme) SetApplyCb(theme *lib.LvThemeT, apply_cb lib.LvThemeApplyCbT) SetTheme {
	C.lv_theme_set_apply_cb((*C.lv_theme_t)(unsafe.Pointer(theme)), C.lv_theme_apply_cb_t(apply_cb))

	return setter
}
