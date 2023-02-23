package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type set struct {
	CStructLvObjT *C.struct__lv_obj_t
}
