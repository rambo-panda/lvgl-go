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
type _cStyleT = C.lv_style_t
type CStyleT = *_cStyleT

type _cAnimT = C.lv_anim_t
type CAnimT = *_cAnimT

type _m[T CObjT | CStyleT | CAnimT] struct {
	o T
}

func (m _m[T]) GetObj() T {
	return m.o
}

var CREATE_NIL _m[CObjT] = _m[CObjT]{}

type _createI interface {
	*_labelT | *_imgT | *_objT | *_styleT | *_barT | *_m[CObjT]

	GetObj() CObjT
}

type _labelT struct {
	_m[CObjT]
	Set *set.Label
	Get *get.Label
}

type _imgT struct {
	_m[CObjT]
	Set *set.Img
	Get *get.Img
}

type _objT struct {
	_m[CObjT]
	Set *set.Obj
	Get *get.Obj
}

type _styleT struct {
	_m[CStyleT]
	Set *set.Style
	Get *get.Style
}

type _barT struct {
	_m[CObjT]
	Set *set.Bar
	Get *get.Bar
}

type _animT struct {
	_m[CAnimT]
	Set *set.Anim
	Get *get.Anim
}
