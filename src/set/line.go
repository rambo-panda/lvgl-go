package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin/g
#include "lv_init.h"
*/
import "C"

type Line set[CObjT]

func CreateLine(o CObjT) Line {
	return Line{
		CObj: o,
	}
}

func (setter Line) YInvert(en bool) Line {
	C.lv_line_set_y_invert(setter.CObj, C.bool(en))

	return setter
}
