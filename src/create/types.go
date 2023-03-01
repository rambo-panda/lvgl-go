package create

import (
	"lvgl-go/src/get"
	"lvgl-go/src/lib"
	"lvgl-go/src/set"
	"unsafe"
)

/*
#include "lv_init.h"
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

func (m _m[T]) GetObj() unsafe.Pointer {
	return unsafe.Pointer(m.o)
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
