package create

import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
)

/*
#include "lv_init.h"
*/
import "C"

type CObjT = *C.struct__lv_obj_t

type CStyleT = *C.lv_style_t

type _m[T CObjT | CStyleT] struct {
	o T
}

func (m _m[T]) GetObj() T {
	return m.o
}

var CREATE_NIL _m[CObjT] = _m[CObjT]{}

type _createI interface {
	*_labelT | *_imgT | *_objT | *_styleT | *_m[CObjT]

	GetObj() CObjT
}

type _labelT struct {
	_m[CObjT]
	Set set.Label
	Get get.Label
}

type _imgT struct {
	_m[CObjT]
	Set set.Img
	Get get.Img
}

type _objT struct {
	_m[CObjT]
	Set set.Obj
	Get get.Obj
}

type _styleT struct {
	_m[CStyleT]
	Set set.Style
	Get get.Style
}
