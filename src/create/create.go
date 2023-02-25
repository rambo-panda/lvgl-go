package create

import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
	types "lvgl-go/src/types"
)

type tCreate[
	// TODO: any是因为有些set/get未实现
	SetT any | set.Animimg | set.Area | set.Canvas | set.Checkbox | set.Label | set.Line | set.Spangroup | set.Table | set.Theme | set.Arc | set.Bar | set.Chart | set.Img | set.Led | set.Obj | set.Span | set.Style | set.Textarea,
	GetT any | get.Label | get.Obj] struct {
	o   *types.LvObjT
	Set SetT
	Get GetT
}

func (c tCreate[SetT, GetT]) getObj() *types.LvObjT {
	return c.o
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
func CreateLabel(o any) tCreate[set.Label, get.Label] {
	_o := convert(o)

	return tCreate[set.Label, get.Label]{
		o:   _o,
		Set: set.CreateLable(_o),
		Get: get.CreateLable(_o),
	}
}
func CreateObj(o any) tCreate[set.Label, get.Label] {
	_o := convert(o)

	return tCreate[set.Label, get.Label]{
		o:   _o,
		Set: set.CreateLable(_o),
		Get: get.CreateLable(_o),
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
