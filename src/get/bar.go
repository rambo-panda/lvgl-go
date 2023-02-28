package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Bar get[CObjT]

func CreateBar(o CObjT) Bar {
	return Bar{
		CObj: o,
	}
}
