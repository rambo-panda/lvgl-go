package lvgl_go

/*
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: -L./lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
	"lvgl-go/src/types"
	"unsafe"
)

type tCreate[
	SetT set.SetAnimimg | set.SetArea | set.SetCanvas | set.SetCheckbox | set.SetLabel | set.SetLine | set.SetSpangroup | set.SetTable | set.SetTheme | set.SetArc | set.SetBar | set.SetChart | set.SetImg | set.SetLed | set.SetObj | set.SetSpan | set.SetStyle | set.SetTextarea,
	GetT get.Get | get.GetLabel | get.GetObj] struct {
	// _setAnimimg  set.SetAnimimg
	// _setArea     set.SetArea
	// _setCanvas   set.SetCanvas
	// _setCheckbox set.SetCheckbox
	// _setLabel    set.SetLabel
	// _setLine     set.SetLine
	// _setSpangrou set.SetSpangroup
	// _setTable    set.SetTable
	// _setTheme    set.SetTheme
	// _setArc      set.SetArc
	// _setBar      set.SetBar
	// _setChart    set.SetChart
	// _setImg      set.SetImg
	// _setLed      set.SetLed
	// _setObj      set.SetObj
	// _setSpan     set.SetSpan
	// _setStyle    set.SetStyle
	// _setTextarea set.SetTextarea
	Set SetT
	Get GetT
}

func c2goObj(o *C.struct__lv_obj_t) *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(o))
}
func getParent(o *types.LvObjT) *C.struct__lv_obj_t {
	if o == nil {
		return C.lv_scr_act()
	}

	return (*C.struct__lv_obj_t)(unsafe.Pointer(o))
}

//	func CreateAnimimg(o *types.LvObjT) tCreate {
//		var c tCreate[set.SetAnimimg, get.GetObj]
//		return tCreate{
//			Set: set.CreateAnimimg(getParent(o)),
//		}
//	}
//
//	func CreateArc(o *types.LvObjT) _create {
//		return set.CreateArc(getParent(o))
//	}
//
//	func CreateArea(o *types.LvObjT) set.SetArea {
//		return set.CreateArea(getParent(o))
//	}
//
//	func CreateBar(o *types.LvObjT) set.SetBar {
//		return set.CreateBar(getParent(o))
//	}
//
//	func CreateCanvas(o *types.LvObjT) set.SetCanvas {
//		return set.CreateCanvas(getParent(o))
//	}
//
//	func CreateChart(o *types.LvObjT) set.SetChart {
//		return set.CreateChart(getParent(o))
//	}
//
//	func CreateCheckbox(o *types.LvObjT) set.SetCheckbox {
//		return set.CreateCheckbox(getParent(o))
//	}
//
//	func CreateImg(o *types.LvObjT) set.SetImg {
//		return set.CreateImg(getParent(o))
//	}
func CreateLabel(o *types.LvObjT) tCreate[set.SetLabel, get.GetLabel] {
	p := getParent(o)
	label := C.lv_label_create(p)
	j := c2goObj(label)

	return tCreate[set.SetLabel, get.GetLabel]{
		Set: set.CreateLabel(j),
		// Get: get.CreateLable(j),
	}
}

// func CreateLed(o *types.LvObjT) set.SetLed {
// 	return set.CreateLed(getParent(o))
// }
// func CreateLine(o *types.LvObjT) set.SetLine {
// 	return set.CreateLine(getParent(o))
// }
// func CreateObj(o *types.LvObjT) set.SetObj {
// 	return set.CreateObj(getParent(o))
// }
// func CreateSpan(o *types.LvObjT) set.SetSpan {
// 	return set.CreateSpan(getParent(o))
// }
// func CreateSpangroup(o *types.LvObjT) set.SetSpangroup {
// 	return set.CreateSpangroup(getParent(o))
// }
// func CreateStyle(o *types.LvObjT) set.SetStyle {
// 	return set.CreateStyle(getParent(o))
// }
// func CreateTable(o *types.LvObjT) set.SetTable {
// 	return set.CreateTable(getParent(o))
// }
// func CreateTextarea(o *types.LvObjT) set.SetTextarea {
// 	return set.CreateTextarea(getParent(o))
// }
// func CreateTheme(o *types.LvObjT) set.SetTheme {
// 	return set.CreateTheme(getParent(o))
// }
