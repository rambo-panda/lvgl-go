package lvgl_go

/*
#include "lv_init.h"
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: -Llib -llvgl
*/
import "C"
import "unsafe"

type LvObj C.struct__lv_obj_t

func Label(parent *LvObj) *LvObj {
	if parent == nil {
		parent = (*LvObj)(unsafe.Pointer(C.lv_scr_act()))
	}

	_parent := (*C.struct__lv_obj_t)(unsafe.Pointer(parent))

	return (*LvObj)(unsafe.Pointer(C.lv_label_create(_parent)))
}

func SetText(obj *LvObj, str string) {
	C.lv_label_set_text((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(str))
}

func Demo() {
	label := Label(nil)
	SetText(label, "hello world")
}

// /*
// #cgo CFLAGS: -I./include/
// #include "lvgl.h"
// #cgo LDFLAGS: -Llibs -llvgl
// #include <stdlib.h>
// */
// import "C"
// import "unsafe"

// // Lvobj ...
// type Lvobj C.struct__lv_obj_t

// const (
// 	LV_ALIGN_CENTER C.uchar = C.LV_ALIGN_CENTER
// )

// func init() {
// 	C.lv_init()
// 	C.fbdev_init()
// 	C.disp_init()
// }

// // ScrAct ...
// func ScrAct() *Lvobj {
// 	return (*Lvobj)(unsafe.Pointer(C.lv_scr_act()))
// }

// // Label ...
// func Label(par, copy *Lvobj) *Lvobj {
// 	var p1 *C.struct__lv_obj_t
// 	var p2 *C.struct__lv_obj_t
// 	p1 = (*C.struct__lv_obj_t)(unsafe.Pointer(par))
// 	p2 = (*C.struct__lv_obj_t)(unsafe.Pointer(copy))
// 	return (*Lvobj)(unsafe.Pointer(C.lv_label_create(p1, p2)))
// }

// // SetText ...
// func (obj *Lvobj) SetText(str string) {
// 	C.lv_label_set_text((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), C.CString(str))
// }

// // Align ...
// func (obj *Lvobj) Align(base *Lvobj, align C.uchar, x int, y int) {
// 	var ba *C.struct__lv_obj_t
// 	ba = (*C.struct__lv_obj_t)(unsafe.Pointer(base))
// 	C.lv_obj_align((*C.struct__lv_obj_t)(unsafe.Pointer(obj)), ba, align, C.short(x), C.short(y))
// }

// // TickInc ...
// func TickInc(tick int) {
// 	C.lv_tick_inc(C.uint(tick))
// }

// // TaskHandler ...
// func TaskHandler() {
// 	C.lv_task_handler()
// }
