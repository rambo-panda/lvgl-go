package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"

type Style get[CStyleT]

func CreateStyle(o CStyleT) *Style {
	_o := Style{
		CObj: o,
	}
	return &_o
}
