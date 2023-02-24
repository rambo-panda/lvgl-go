package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type LvObjT *C.struct__lv_obj_t
type get struct {
	CStructLvObjT LvObjT
}
