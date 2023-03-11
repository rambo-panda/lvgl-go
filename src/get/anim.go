package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"

type Anim get[CAnimT]

func CreateAnim(o CAnimT) *Anim {
	_o := Anim{
		CObj: o,
	}
	return &_o
}
