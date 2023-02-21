package lvgl_go

/*
#include "lv_init.h"
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func toShort(a any) C.short {
	return a.(C.short)
}

type s_attr struct {
	t string
}

func (f s_attr) ToMethoName() string {
	return genMethodBuilder("Set", f.t)
}
func (f s_attr) SetW(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_width(o, toShort(args[0]))
}
func (f s_attr) SetX(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_x(o, toShort(args[0]))
}
func (f s_attr) SetH(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_height(o, toShort(args[0]))
}
func (f s_attr) SetY(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_y(o, toShort(args[0]))
}
func (f s_attr) SetSize(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_size(o, toShort(args[0]), toShort(args[1]))
}
func (f s_attr) SetPos(o *C.struct__lv_obj_t, args ...any) {
	// fmt.Println("如需同时设置align & pos 推荐使用 `lv_obj_align(obj, LV_ALIGN_CENTER, 10, 20)`")
	C.lv_obj_set_pos(o, toShort(args[0]), toShort(args[1]))
}
func (f s_attr) SetParent(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_parent(o, args[0].(*C.struct__lv_obj_t))
}
func (f s_attr) SetChild(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_parent(o, args[0].(*C.struct__lv_obj_t))
}
func (f s_attr) SetAlign(o *C.struct__lv_obj_t, args ...any) {
	C.lv_obj_set_align(o, (C.uint8_t)(args[0].(uint8)))
}
func (f s_attr) SetText(o *C.struct__lv_obj_t, args ...any) {
	textObjs := [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
}

// use Map https://stackoverflow.com/questions/46377302/golang-when-iterate-map-how-to-get-key-as-pointer
const textObjs := [...]*C.LV_OBJ_CLASS_ARR{
	&C.lv_label_class,
	&C.lv_checkbox_class,
	&C.lv_dropdown_class,
	&C.lv_spangroup_class,
	&C.lv_textarea_class
};

func Attr(o *LvObj, t string, args ...any) *LvObj {
	if o == nil {
		return o
	}

	defer func() {
		err := recover()

		if err != nil {
			fmt.Println("set attr error", err)
		}
	}()

	_attr := &s_attr{t}

	in := make([]reflect.Value, len(args)+1)
	in[0] = reflect.ValueOf((*C.struct__lv_obj_t)(unsafe.Pointer(o)))

	for i, v := range args {
		in[i] = reflect.ValueOf(v)
	}

	reflect.ValueOf(_attr).MethodByName(_attr.ToMethoName()).Call(in)

	return o
}
