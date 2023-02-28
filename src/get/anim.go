package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Anim get[CAnimT]

func CreateAnim(o CAnimT) *Anim {
	_o := Anim{
		CObj: o,
	}
	return &_o
}
