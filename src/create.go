package lvgl_go

/*
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: -L./lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"fmt"
	"lvgl-go/src/get"
	"lvgl-go/src/lib"
	"lvgl-go/src/set"
	"unsafe"
)

type tCreate[
	T set.Animimg | set.Area | set.Canvas | set.Checkbox | set.Label | set.Line | set.Spangroup | set.Table | set.Theme | set.Arc | set.Bar | set.Chart | set.Img | set.Led | set.Obj | set.Span | set.Style | set.Textarea,
	T get.Label | get.Obj] struct {
	// _setAnimimg  set.Animimg
	// _setArea     set.Area
	// _setCanvas   set.Canvas
	// _setCheckbox set.Checkbox
	// _setLabel    set.Label
	// _setLine     set.Line
	// _setSpangrou set.Spangroup
	// _setTable    set.Table
	// _setTheme    set.Theme
	// _setArc      set.Arc
	// _setBar      set.Bar
	// _setChart    set.Chart
	// _setImg      set.Img
	// _setLed      set.Led
	// _setObj      set.Obj
	// _setSpan     set.Span
	// _setStyle    set.Style
	// _setTextarea set.Textarea
	Set T
	Get T
}

// func c2goObj[T *C.struct__lv_obj_t | set.TsetC](o T, t string) *lib.LvObjT {
// 	switch t {
// 	case "darwin":
// 		fmt.Println("OS X.")
// 	case "linux":
// 		return (*lib.LvObjT)(unsafe.Pointer(o))
// 	default:
// 		return (*lib.LvObjT)(unsafe.Pointer(o))
// 	}
// }

func go2cObj(o *lib.LvObjT, t string) any {
	_o := unsafe.Pointer(o)

	switch t {
	case "darwin":
		fmt.Println("OS X.")
	case "set":
		return (*set.TsetC)(unsafe.Pointer(o))
	default:
		return (*C.struct__lv_obj_t)(_o)
	}
}
func getParent(o *lib.LvObjT) *C.struct__lv_obj_t {
	if o == nil {
		return C.lv_scr_act()
	}

	return go2cObj(o)
}

//	func CreateAnimimg(o *lib.LvObjT) tCreate {
//		var c tCreate[set.Animimg, get.Obj]
//		return tCreate{
//			: set.CreateAnimimg(getParent(o)),
//		}
//	}
//
//	func CreateArc(o *lib.LvObjT) _create {
//		return set.CreateArc(getParent(o))
//	}
//
//	func CreateArea(o *lib.LvObjT) set.Area {
//		return set.CreateArea(getParent(o))
//	}
//
//	func CreateBar(o *lib.LvObjT) set.Bar {
//		return set.CreateBar(getParent(o))
//	}
//
//	func CreateCanvas(o *lib.LvObjT) set.Canvas {
//		return set.CreateCanvas(getParent(o))
//	}
//
//	func CreateChart(o *lib.LvObjT) set.Chart {
//		return set.CreateChart(getParent(o))
//	}
//
//	func CreateCheckbox(o *lib.LvObjT) set.Checkbox {
//		return set.CreateCheckbox(getParent(o))
//	}
//
//	func CreateImg(o *lib.LvObjT) set.Img {
//		return set.CreateImg(getParent(o))
//	}
func CreateLabel(o *lib.LvObjT) tCreate[set.Label, get.Label] {
	p := getParent(o)
	label := C.lv_label_create(p)
	// j := c2goObj(label)
	j := label

	return tCreate[set.Label, get.Label]{
		: set.CreateLabel(j),
		: get.CreateLable(j),
	}
}

// func CreateLed(o *lib.LvObjT) set.Led {
// 	return set.CreateLed(getParent(o))
// }
// func CreateLine(o *lib.LvObjT) set.Line {
// 	return set.CreateLine(getParent(o))
// }
// func CreateObj(o *lib.LvObjT) set.Obj {
// 	return set.CreateObj(getParent(o))
// }
// func CreateSpan(o *lib.LvObjT) set.Span {
// 	return set.CreateSpan(getParent(o))
// }
// func CreateSpangroup(o *lib.LvObjT) set.Spangroup {
// 	return set.CreateSpangroup(getParent(o))
// }
// func CreateStyle(o *lib.LvObjT) set.Style {
// 	return set.CreateStyle(getParent(o))
// }
// func CreateTable(o *lib.LvObjT) set.Table {
// 	return set.CreateTable(getParent(o))
// }
// func CreateTextarea(o *lib.LvObjT) set.Textarea {
// 	return set.CreateTextarea(getParent(o))
// }
// func CreateTheme(o *lib.LvObjT) set.Theme {
// 	return set.CreateTheme(getParent(o))
// }
