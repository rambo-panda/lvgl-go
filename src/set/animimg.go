package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type SetAnimimg set

func (setter SetAnimimg) SetDuration(duration uint32) SetAnimimg {
	C.lv_animimg_set_duration(setter.cObj, C.uint(duration))

	return setter
}
func (setter SetAnimimg) SetRepeatCount(count uint16) SetAnimimg {
	C.lv_animimg_set_repeat_count(setter.cObj, C.ushort(count))

	return setter
}
