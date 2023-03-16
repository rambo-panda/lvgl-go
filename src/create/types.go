package create

import (
	"fmt"
	"unsafe"

	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/get"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/set"
)

/*
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
*/
import "C"

type _lvObjT = *C.lv_obj_t
type _cStyleT = C.lv_style_t
type _PcStyleT = *_cStyleT

type _cAnimT = C.lv_anim_t
type CAnimT = *_cAnimT

type _m[T _lvObjT | _PcStyleT | CAnimT] struct {
	o T
}

func (m *_m[T]) GetObj() unsafe.Pointer {
	return unsafe.Pointer(m.o)
}

const (
	DEL_ASYNC uint8 = 1
	DEL       uint8 = 2
)

func (m *_m[T]) Destroy(tag uint8) {
	defer func() {
		// NOTE: 这里偷懒了，再用reflect去判断是否是LV_OBJ_T 这个代价还不如直接recover呢
		if err := recover(); err != nil {
			fmt.Println("捕获异常:", err)
		}
	}()

	_o := (_lvObjT)(m.GetObj())

	switch tag {
	case DEL_ASYNC:
		C.lv_obj_del_async(_o)
	case DEL:
		C.lv_obj_del(_o)
	default:
		C.lv_obj_del(_o)
	}
}

type _createI = lib.CreateI

type Label struct {
	_m[_lvObjT]
	Set *set.Label
	Get *get.Label
}

type Img struct {
	_m[_lvObjT]
	Set *set.Img
	Get *get.Img
}

type Obj struct {
	_m[_lvObjT]
	Set *set.Obj
	Get *get.Obj
}

type StyleT struct {
	_m[_PcStyleT]
	Set *set.Style
	Get *get.Style
}

type Bar struct {
	_m[_lvObjT]
	Set *set.Bar
	Get *get.Bar
}

type AnimT struct {
	_m[CAnimT]
	Set *set.Anim
	Get *get.Anim
}
