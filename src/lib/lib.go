package lib

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L./ -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

func Ready() {
	C.lv_ready()
}

func TaskHandler(ms uint) {
	C.lv_task_handler2(C.uint(ms))
}

func TaskHandlerAsync(ms uint) {
	go TaskHandler(ms)
}

func Noop(s ...any) {
}

func C2GoString(o unsafe.Pointer) string {
	return C.GoString((*C.char)(o))
}

func Go2CString(o string) unsafe.Pointer {
	cs := C.CString(o)

	//TODO: 何时删除?? lvgl应该会自己删除 C.free(unsafe.Pointer(cs))

	return unsafe.Pointer(cs)
}

func Ternary[T any](c bool, trueVal, falseVal T) T {
	if c {
		return trueVal
	}

	return falseVal
}
