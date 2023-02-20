package lvgl_go

/*
#include "lv_init.h"
*/
import "C"
import (
	"reflect"
	"strings"
	"unsafe"
)

type lv struct {
	t string
}

func (f lv) ToMethoName() string {
	return strings.Join([]string{"Create", strings.ToUpper(f.t[:1]), f.t[1:]}, "")
}
func (f lv) CreateLabel(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return C.lv_label_create(parent)
}
func (f lv) CreateObj(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return C.lv_obj_create(parent)
}
func (f lv) CreateScreen(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return f.CreateObj(parent)
}

func Create(t string, parent *LvObj) *LvObj {
	if parent == nil && t != "screen" {
		parent = (*LvObj)(unsafe.Pointer(C.lv_scr_act()))
	}

	_lv := &lv{t: strings.ToLower(t)}

	in := make([]reflect.Value, 1)
	in[0] = reflect.ValueOf((*C.struct__lv_obj_t)(unsafe.Pointer(parent)))

	ret := reflect.ValueOf(_lv).MethodByName(_lv.ToMethoName()).Call(in)

	return LvObjToGo(ret[0].Interface().(*C.struct__lv_obj_t))
}

func LvObjToC(o *LvObj) *C.struct__lv_obj_t {
	return (*C.struct__lv_obj_t)(unsafe.Pointer(o))
}

func LvObjToGo(o *C.struct__lv_obj_t) *LvObj {
	return (*LvObj)(o)
}
