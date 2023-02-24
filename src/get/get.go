package get

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type TgetC *C.struct__lv_obj_t
type get struct {
	CStructLvObjT TgetC
}
