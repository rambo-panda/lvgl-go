package create

import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
	"lvgl-go/src/types"
	"unsafe"
)

type tCreate[
	SetT set.Animimg | set.Area | set.Canvas | set.Checkbox | set.Label | set.Line | set.Spangroup | set.Table | set.Theme | set.Arc | set.Bar | set.Chart | set.Img | set.Led | set.Obj | set.Span | set.Style | set.Textarea,
	GetT get.Label | get.Obj] struct {
	Set SetT
	Get GetT
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
//			: set.CreateAnimimg(getParent(o)),
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
func c2goObj(o *C.struct__lv_obj_t) *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(o))
}
func CreateLabel(o *types.LvObjT) tCreate[set.Label, get.Label] {
	p := getParent(o)
	label := C.lv_label_create(p)
	j := c2goObj(label)

	return tCreate[set.Label, get.Label]{
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
