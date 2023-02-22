package lvgl_go

/*
#include "lv_init.h"
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: -Llib -llvgl
*/
import "C"
import "unsafe"

func SetText(obj *LvObj, str string) {
	C.lv_label_set_text(LvObjToC(obj), C.CString(str))
}

func noop(x any) {
}

func SetObjWidth(obj *LvObj, w int16) {
	ret := C.lv_obj_set_width((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.lv_coord_t(w))

	noop(ret)

	return
}

func Demo() {
	label := Create("label", nil)
	SetText(label, "hello world")
}
