package create

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
)

type tCreate[
	// TODO: any是因为有些set/get未实现
	SetT any | set.Animimg | set.Area | set.Canvas | set.Checkbox | set.Label | set.Line | set.Spangroup | set.Table | set.Theme | set.Arc | set.Bar | set.Chart | set.Img | set.Led | set.Obj | set.Span | set.Style | set.Textarea,
	GetT any | get.Label | get.Obj] struct {
	o   CObjT
	Set SetT
	Get GetT
}

func (t tCreate[SetT, GetT]) getObj() *C.struct__lv_obj_t {
	return t.o
}

//	func CreateAnimimg(o CObjT) tCreate {
//		var c tCreate[set.SetAnimimg, get.GetObj]
//		return tCreate{
//			: set.CreateAnimimg(getParent(o)),
//		}
//	}
//
//	func CreateArc(o CObjT) _create {
//		return set.CreateArc(getParent(o))
//	}
//
//	func CreateArea(o CObjT) set.SetArea {
//		return set.CreateArea(getParent(o))
//	}
//
//	func CreateBar(o CObjT) set.SetBar {
//		return set.CreateBar(getParent(o))
//	}
//
//	func CreateCanvas(o CObjT) set.SetCanvas {
//		return set.CreateCanvas(getParent(o))
//	}
//
//	func CreateChart(o CObjT) set.SetChart {
//		return set.CreateChart(getParent(o))
//	}
//
//	func CreateCheckbox(o CObjT) set.SetCheckbox {
//		return set.CreateCheckbox(getParent(o))
//	}
//
//	func CreateImg(o CObjT) set.SetImg {
//		return set.CreateImg(getParent(o))
//	}
func CreateLabel(o CObjT) tCreate[set.Label, get.Label] {
	goO := C.lv_label_create(getParent(o))

	return tCreate[set.Label, get.Label]{
		o:   goO,
		Set: set.CreateLabel(toSetObj(goO)),
		Get: get.CreateLable(toGetObj(goO)),
	}
}

func CreateObj(o any) tCreate[set.Obj, get.Obj] {
	_o := convertObj(o)

	return tCreate[set.Obj, get.Obj]{
		o:   _o,
		Set: set.CreateObj(_o),
		Get: get.CreateObj(_o),
	}
}

// func CreateLed(o CObjT) set.SetLed {
// 	return set.CreateLed(getParent(o))
// }
// func CreateLine(o CObjT) set.SetLine {
// 	return set.CreateLine(getParent(o))
// }
// func CreateSpan(o CObjT) set.SetSpan {
// 	return set.CreateSpan(getParent(o))
// }
// func CreateSpangroup(o CObjT) set.SetSpangroup {
// 	return set.CreateSpangroup(getParent(o))
// }
// func CreateStyle(o CObjT) set.SetStyle {
// 	return set.CreateStyle(getParent(o))
// }
// func CreateTable(o CObjT) set.SetTable {
// 	return set.CreateTable(getParent(o))
// }
// func CreateTextarea(o CObjT) set.SetTextarea {
// 	return set.CreateTextarea(getParent(o))
// }
// func CreateTheme(o CObjT) set.SetTheme {
// 	return set.CreateTheme(getParent(o))
// }
