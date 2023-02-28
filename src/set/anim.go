package set

import (
	"lvgl-go/src/lib"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Anim set[CAnimT]
type lvCoordT = C.lv_coord_t

func CreateAnim(o CAnimT) Anim {
	return Anim{
		CObj: o,
	}
}
