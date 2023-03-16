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

func create[T _labelT | _imgT | _objT | _barT](o T) T {
	runtime.SetFinalizer(&o, func(z *_labelT) {
		z.Destroy(DEL)
	})
	return o
}

func CreateLabel[T _createI](o T) _labelT {
	_o := C.lv_label_create(getParent(o))

	return create(_labelT{
		_m[LV_OBJ_T]{_o},
		set.CreateLabel(toSetObj(_o)),
		get.CreateLable(toGetObj(_o)),
	})
}

func CreateImg[T _createI](o T) _imgT {
	_o := C.lv_img_create(getParent(o))

	return create(_imgT{
		_m[LV_OBJ_T]{_o},
		set.CreateImg(toSetObj(_o)),
		get.CreateImg(toGetObj(_o)),
	})
}

func CreateObj[T _createI](o T) _objT {
	_o := getParent2(o)

	return create(_objT{
		_m[LV_OBJ_T]{_o},
		set.CreateObj(toSetObj(_o)),
		get.CreateObj(toGetObj(_o)),
	})
}

func CreateBar[T _createI](o T) _barT {
	_o := C.lv_bar_create(getParent(o))

	return create(_barT{
		_m[LV_OBJ_T]{_o},
		set.CreateBar(toSetObj(_o)),
		get.CreateBar(toGetObj(_o)),
	})
}

func CreateAnim() _animT {
	obj := _cAnimT{}
	anim := &obj

	C.lv_anim_init(anim)

	return _animT{
		_m[CAnimT]{anim},
		set.CreateAnim(toSetAnimT(anim)),
		get.CreateAnim(toGetAnimT(anim)),
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
