package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type CObjT = *C.struct__lv_obj_t
type CStyleT = *C.lv_style_t
type CAnimT = *C.lv_anim_t

type set[T CObjT | CStyleT | CAnimT] struct {
	CObj T
}

// TODO: 看如何直接使用create _createI 现在怕的就是互相依赖
func getParentObj(o any) unsafe.Pointer {
	return (reflect.ValueOf(o).MethodByName("GetObj").Call(nil)[0].UnsafePointer())
}
