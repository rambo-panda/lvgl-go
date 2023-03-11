package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"

type Bar get[CObjT]

func CreateBar(o CObjT) *Bar {
	_o := Bar{
		CObj: o,
	}
	return &_o
}
