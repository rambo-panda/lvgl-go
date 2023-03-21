package create

import (
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
type _pAnimT = *_cAnimT

type _m[T _lvObjT | _PcStyleT | _pAnimT] struct {
	o      T
	parent any
	child  any
}

func (m *_m[T]) GetObj() unsafe.Pointer {
	return unsafe.Pointer(m.o)
}
func (m *_m[T]) Destroy(tag lib.DelT) {
	lib.Destroy(m, tag)
	m.o = nil
}
func (m *_m[T]) SetChild(child any) {
	m.child = child
}
func (m *_m[T]) SetParent(child any) {
	m.parent = child
}

func createM(o _lvObjT, parent _createI) _m[_lvObjT] {
	j := _m[_lvObjT]{o, parent, nil}

	return j
}

var CREATE_NIL = &_m[_lvObjT]{}

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
	_m[_pAnimT]
	Set *set.Anim
	Get *get.Anim
}
