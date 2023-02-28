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

func CreateLabel[T _createI](o T) _labelT {
	_o := C.lv_label_create(getParent(o))

	return _labelT{
		_m[CObjT]{_o},
		set.CreateLabel(toSetObj(_o)),
		get.CreateLable(toGetObj(_o)),
	}
}

func CreateImg[T _createI](o T) _imgT {
	_o := C.lv_img_create(getParent(o))
	return _imgT{
		_m[CObjT]{_o},
		set.CreateImg(toSetObj(_o)),
		get.CreateImg(toGetObj(_o)),
	}
}

func CreateObj[T _createI](o T) _objT {
	_o := getParent2(o)

	return _objT{
		_m[CObjT]{_o},
		set.CreateObj(toSetObj(_o)),
		get.CreateObj(toGetObj(_o)),
	}
}

func CreateStyle() _styleT {
	obj := _cStyleT{}
	so := &obj

	C.lv_style_init(so)

	return _styleT{
		_m[CStyleT]{so},
		set.CreateStyle(toSetStyleT(so)),
		get.CreateStyle(toGetStyleT(so)),
	}
}

func CreateBar[T _createI](o T) _barT {
	_o := C.lv_bar_create(getParent(o))

	return _barT{
		_m[CObjT]{_o},
		set.CreateBar(toSetObj(_o)),
		get.CreateBar(toGetObj(_o)),
	}
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
