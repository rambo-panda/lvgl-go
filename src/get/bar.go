package get

/*
#cgo CFLAGS: -I../include/
#include "lv_init.h"
*/
import "C"

type Bar get[CObjT]

func CreateBar(o CObjT) *Bar {
	_o := Bar{
		CObj: o,
	}
	return &_o
}
