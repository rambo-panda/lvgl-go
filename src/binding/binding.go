package lvgl_go

/*
#include "lv_init.h"
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: -Llib -llvgl
*/
import "C"
import (
	"unsafe"
)

func SetText(obj *LvObj, str string) {
	C.lv_label_set_text((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(str))
}

func Demo() {
	label := Create("label", nil)
	SetText(label, "hello world")
}
