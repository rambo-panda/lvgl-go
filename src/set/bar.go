package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
)

type SetBar set

func (setter SetBar) SetValue(value int32, anim lib.LvAnimEnableT) SetBar {
	C.lv_bar_set_value(setter.cObj, C.int(value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter SetBar) SetStartValue(start_value int32, anim lib.LvAnimEnableT) SetBar {
	C.lv_bar_set_start_value(setter.cObj, C.int(start_value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter SetBar) SetRange(min int32, max int32) SetBar {
	C.lv_bar_set_range(setter.cObj, C.int(min), C.int(max))

	return setter
}
func (setter SetBar) SetMode(mode lib.LvBarModeT) SetBar {
	C.lv_bar_set_mode(setter.cObj, C.lv_bar_mode_t(mode))

	return setter
}
