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

type SetLine set

func (setter SetLine) SetYInvert(en bool) SetLine {
	C.lv_line_set_y_invert(setter.cObj, C.bool(en))

	return setter
}
