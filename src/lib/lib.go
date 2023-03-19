package lib

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func Ready() {
	C.lv_ready()
}

func TaskHandler(ms uint) {
	C.lv_task_handler2(C.uint(ms))
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

type DelT uint8

const (
	DEL_ASYNC DelT = 1
	DEL       DelT = 2
)

func Destroy(m CreateI, tag DelT) {
	defer func() {
		// NOTE: 这里偷懒了，再用reflect去判断是否是LV_OBJ_T 这个代价还不如直接recover呢
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()

	_o := (LV_OBJ_T)(m.GetObj())

	if _o == nil {
		return
	}

	switch tag {
	case DEL_ASYNC:
		C.lv_obj_del_async(_o)
	case DEL:
		C.lv_obj_del(_o)
	}
}

func GetChineseStyle() unsafe.Pointer {
	return unsafe.Pointer(C.getChineseStyle())
}
