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

const SCREEN string = "SCREEN"

type s_create struct {
	o string
	o int

	Set
}
func (f s_create) GetObj() *lib.LvObjT {
	return f.Set._o;
}

func (f s_create) ToMethoName() string {
	return genMethodBuilder("Create", f.t)
}
func (f s_create) CreateLabel(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return C.lv_label_create(parent)
}
func (f s_create) CreateImg(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return C.lv_img_create(parent)
}
func (f s_create) CreateObj(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return C.lv_obj_create(parent)
}
func (f s_create) CreateScreen(parent *C.struct__lv_obj_t) *C.struct__lv_obj_t {
	return f.CreateObj(parent)
}

func Create(t string, parent *LvObj) *LvObj {
	if parent == nil && t != SCREEN {
		parent = (*LvObj)(unsafe.Pointer(C.lv_scr_act()))
	}

	_lv := &s_create{t: strings.ToLower(t)}

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

func mkAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func genMethodBuilder(pre string, f string) string {
	return strings.Join([]string{pre, strings.ToUpper(f[:1]), f[1:]}, "")
}
