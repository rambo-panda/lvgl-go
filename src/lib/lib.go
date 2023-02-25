package lib

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L./ -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/types"
	"reflect"
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

func C2GoObj(o any) *types.LvObjT {
	return (*types.LvObjT)(reflect.ValueOf(o).UnsafePointer())
}

func IsNil(o any) bool {
	return o == nil || reflect.ValueOf(o).IsNil()
}

func GetParent(o any) *types.LvObjT {
	if IsNil(o) {
		return C2GoObj(C.lv_scr_act())
	}

	return o.(*types.LvObjT)
}

func Ternary[T any](c bool, trueVal, falseVal T) T {
	if c {
		return trueVal
	}

	return falseVal
}
