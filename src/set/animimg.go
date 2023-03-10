package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin/g
#include "lv_init.h"
*/
import "C"

type Animimg set[CObjT]

func CreateAnimimg(o CObjT) Animimg {
	return Animimg{
		CObj: o,
	}
}

func (setter Animimg) Duration(duration uint32) Animimg {
	C.lv_animimg_set_duration(setter.CObj, C.uint(duration))

	return setter
}
func (setter Animimg) RepeatCount(count uint16) Animimg {
	C.lv_animimg_set_repeat_count(setter.CObj, C.ushort(count))

	return setter
}
