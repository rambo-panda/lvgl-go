package create

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"
import (
	"runtime"
	"unsafe"

	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/get"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/set"
)

func create[T Label | Img | Obj | Bar](o *T) *T {
	runtime.SetFinalizer(o, func(z _createI) {
		lib.Destroy(z, lib.DEL)
	})

	return o
}

func CreateLabel[T _createI](o T) *Label {
	_o := C.lv_label_create(getParent(o, normal))

	return create(&Label{
		_m[_lvObjT]{_o},
		set.CreateLabel(toSetObj(_o)),
		get.CreateLable(toGetObj(_o)),
	})
}

func CreateImg[T _createI](o T) *Img {
	_o := C.lv_img_create(getParent(o, normal))

	return create(&Img{
		_m[_lvObjT]{_o},
		set.CreateImg(toSetObj(_o)),
		get.CreateImg(toGetObj(_o)),
	})
}

func _createObj[T _createI](o T, tag tagUint) *Obj {
	_o := getParent(o, tag)

	return create(&Obj{
		_m[_lvObjT]{_o},
		set.CreateObj(toSetObj(_o)),
		get.CreateObj(toGetObj(_o)),
	})
}

func CreateObj[T _createI](o T) *Obj {
	return _createObj(o, container)
}

func CreateScreen() *Obj {
	screen := _createObj(&lib.CREATE_NIL, screen)
	C.lv_scr_load((_lvObjT)(screen.GetObj()))

	return screen
}

func CreateBar[T _createI](o T) *Bar {
	_o := C.lv_bar_create(getParent(o, normal))

	return create(&Bar{
		_m[_lvObjT]{_o},
		set.CreateBar(toSetObj(_o)),
		get.CreateBar(toGetObj(_o)),
	})
}

func CreateAnim(a unsafe.Pointer) *AnimT {
	var anim _pAnimT
	if a == nil {
		anim = &_cAnimT{}
	} else {
		anim = (_pAnimT)(a)
	}

	C.lv_anim_init(anim)

	ret := &AnimT{
		_m[_pAnimT]{anim},
		set.CreateAnim(toSetAnimT(anim)),
		get.CreateAnim(toGetAnimT(anim)),
	}

	// TODO: 这一块的remove需要remove cb
	// runtime.SetFinalizer(ret, func(z _createI) {
	// 	C.lv_anim_del((_PcStyleT)(z.GetObj()), anim_cb)
	// })

	return ret
}

func CreateStyle(a unsafe.Pointer) *StyleT {
	var so _PcStyleT

	if a == nil {
		so = &_cStyleT{}
		C.lv_style_init(so)
	} else {
		so = (_PcStyleT)(a)
	}

	ret := &StyleT{
		_m[_PcStyleT]{so},
		set.CreateStyle(toSetStyleT(so)),
		get.CreateStyle(toGetStyleT(so)),
	}

	runtime.SetFinalizer(ret, func(z _createI) {
		C.lv_style_reset((_PcStyleT)(z.GetObj()))
	})

	return ret
}
