package create

/*
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/get"
	"lvgl-go/src/lib"
	"lvgl-go/src/set"
	"unsafe"
)

type CObjT = *C.struct__lv_obj_t

func getParent(o CObjT) CObjT {
	if o == nil {
		return C.lv_scr_act()
	}

	return o
}

func toSetObj(o CObjT) set.CObjT {
	return (set.CObjT)(unsafe.Pointer(o))
}

func toGetObj(o CObjT) get.CObjT {
	return (get.CObjT)(unsafe.Pointer(o))
}

func convertObj(t any) CObjT {
	if lib.IsNil(t) {
		return C.lv_obj_create(nil)
	}
	if o, ok := t.(tCreate[set.Label, get.Label]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Obj, get.Obj]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Animimg, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Area, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Canvas, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Checkbox, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Line, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Spangroup, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Table, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Theme, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Arc, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Bar, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Chart, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Img, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Led, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Span, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Style, any]); ok {
		return o.getObj()
	}
	if o, ok := t.(tCreate[set.Textarea, any]); ok {
		return o.getObj()
	}

	return nil
}
