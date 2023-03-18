package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"

type Bar set[CObjT]

func CreateBar(o CObjT) *Bar {
	_o := Bar{
		CObj: o,
	}

	return &_o
}

func (setter *Bar) Value(value int32, anim uint32) *Bar {
	C.lv_bar_set_value(setter.CObj, C.int(value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter *Bar) StartValue(start_value int32, anim uint8) *Bar {
	C.lv_bar_set_start_value(setter.CObj, C.int(start_value), C.lv_anim_enable_t(anim))

	return setter
}
func (setter *Bar) Range(min int32, max int32) *Bar {
	C.lv_bar_set_range(setter.CObj, C.int(min), C.int(max))

	return setter
}
func (setter *Bar) Mode(mode uint8) *Bar {
	C.lv_bar_set_mode(setter.CObj, C.lv_bar_mode_t(mode))

	return setter
}
func (setter *Bar) Center() *Bar {
	C.lv_obj_center(setter.CObj)

	return setter
}
