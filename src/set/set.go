package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type TsetC *C.struct__lv_obj_t

type set struct {
	CStructLvObjT TsetC
}
