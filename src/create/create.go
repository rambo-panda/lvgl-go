package create

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"
import (
	"runtime"

	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/get"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/set"
)

func create[T Label | Img | Obj | Bar](o T) T {
	runtime.SetFinalizer(&o, func(z *Label) {
		z.Destroy(DEL)
	})
	return o
}

func CreateLabel[T _createI](o T) Label {
	_o := C.lv_label_create(getParent(o))

	return create(Label{
		_m[LV_OBJ_T]{_o},
		set.CreateLabel(toSetObj(_o)),
		get.CreateLable(toGetObj(_o)),
	})
}

func CreateImg[T _createI](o T) Img {
	_o := C.lv_img_create(getParent(o))

	return create(Img{
		_m[LV_OBJ_T]{_o},
		set.CreateImg(toSetObj(_o)),
		get.CreateImg(toGetObj(_o)),
	})
}

func CreateObj[T _createI](o T) Obj {
	_o := getParent2(o)

	return create(Obj{
		_m[LV_OBJ_T]{_o},
		set.CreateObj(toSetObj(_o)),
		get.CreateObj(toGetObj(_o)),
	})
}

func CreateBar[T _createI](o T) Bar {
	_o := C.lv_bar_create(getParent(o))

	return create(Bar{
		_m[LV_OBJ_T]{_o},
		set.CreateBar(toSetObj(_o)),
		get.CreateBar(toGetObj(_o)),
	})
}

func CreateAnim() AnimT {
	obj := _cAnimT{}
	anim := &obj

	C.lv_anim_init(anim)

	return AnimT{
		_m[CAnimT]{anim},
		set.CreateAnim(toSetAnimT(anim)),
		get.CreateAnim(toGetAnimT(anim)),
	}
}

func CreateStyle() StyleT {
	obj := _cStyleT{}
	so := &obj

	C.lv_style_init(so)

	return StyleT{
		_m[CStyleT]{so},
		set.CreateStyle(toSetStyleT(so)),
		get.CreateStyle(toGetStyleT(so)),
	}
}
