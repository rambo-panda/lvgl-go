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

type LV_OBJ_T = *C.lv_obj_t
type _cStyleT = C.lv_style_t
type CStyleT = *_cStyleT

type _cAnimT = C.lv_anim_t
type CAnimT = *_cAnimT

type _m[T LV_OBJ_T | CStyleT | CAnimT] struct {
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

	_o := (LV_OBJ_T)(m.GetObj())

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

type _labelT struct {
	_m[LV_OBJ_T]
	Set *set.Label
	Get *get.Label
}

type _imgT struct {
	_m[LV_OBJ_T]
	Set *set.Img
	Get *get.Img
}

type _objT struct {
	_m[LV_OBJ_T]
	Set *set.Obj
	Get *get.Obj
}

type _styleT struct {
	_m[CStyleT]
	Set *set.Style
	Get *get.Style
}

type _barT struct {
	_m[LV_OBJ_T]
	Set *set.Bar
	Get *get.Bar
}

type _animT struct {
	_m[CAnimT]
	Set *set.Anim
	Get *get.Anim
}
