package create

import (
	"lvgl-go/src/get"
	"lvgl-go/src/set"
)

type _m struct {
	o CObjT
}

func (m _m) GetObj() CObjT {
	return m.o
}

var CREATE_NIL _m = _m{}

type _createI interface {
	*_labelT | *_imgT | *_objT | *_m

	GetObj() CObjT
}

type _labelT struct {
	_m
	Set set.Label
	Get get.Label
}

type _imgT struct {
	_m
	Set set.Img
	Get get.Img
}

type _objT struct {
	_m
	Set set.Obj
	Get get.Obj
}
