package lvgl_go

/*
#cgo CFLAGS: -I./include/
#cgo LDFLAGS: -L./lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/lib"
	"lvgl-go/src/set"
	"unsafe"
)

func getParent(o *lib.LvObjT) *lib.LvObjT {
	if o == nil {
		return (*lib.LvObjT)(unsafe.Pointer(C.lv_scr_act()))
	}

	return o
}
func CreateAnimimg(o *lib.LvObjT) set.SetAnimimg {
	return set.CreateAnimimg(getParent(o))
}
func CreateArc(o *lib.LvObjT) set.SetArc {
	return set.CreateArc(getParent(o))
}
func CreateArea(o *lib.LvObjT) set.SetArea {
	return set.CreateArea(getParent(o))
}
func CreateBar(o *lib.LvObjT) set.SetBar {
	return set.CreateBar(getParent(o))
}
func CreateCanvas(o *lib.LvObjT) set.SetCanvas {
	return set.CreateCanvas(getParent(o))
}
func CreateChart(o *lib.LvObjT) set.SetChart {
	return set.CreateChart(getParent(o))
}
func CreateCheckbox(o *lib.LvObjT) set.SetCheckbox {
	return set.CreateCheckbox(getParent(o))
}
func CreateImg(o *lib.LvObjT) set.SetImg {
	return set.CreateImg(getParent(o))
}
func CreateLabel(o *lib.LvObjT) set.SetLabel {
	return set.CreateLabel(getParent(o))
}
func CreateLed(o *lib.LvObjT) set.SetLed {
	return set.CreateLed(getParent(o))
}
func CreateLine(o *lib.LvObjT) set.SetLine {
	return set.CreateLine(getParent(o))
}
func CreateObj(o *lib.LvObjT) set.SetObj {
	return set.CreateObj(getParent(o))
}
func CreateSpan(o *lib.LvObjT) set.SetSpan {
	return set.CreateSpan(getParent(o))
}
func CreateSpangroup(o *lib.LvObjT) set.SetSpangroup {
	return set.CreateSpangroup(getParent(o))
}
func CreateStyle(o *lib.LvObjT) set.SetStyle {
	return set.CreateStyle(getParent(o))
}
func CreateTable(o *lib.LvObjT) set.SetTable {
	return set.CreateTable(getParent(o))
}
func CreateTextarea(o *lib.LvObjT) set.SetTextarea {
	return set.CreateTextarea(getParent(o))
}
func CreateTheme(o *lib.LvObjT) set.SetTheme {
	return set.CreateTheme(getParent(o))
}
