package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -llvgl -llv_driver -L../bin
#include "lv_17zy.h"
extern void _execAnimCb(void * obj, int temp);
*/
import "C"
import (
	"gitlab.17zuoye.net/saas-platform/lvgl-go.git/src/lib"
	"unsafe"
)

// type Anim set[CAnimT]
type AnimExecCb func(o unsafe.Pointer, temp int32)
type Anim struct {
	set[CAnimT]
	execCb AnimExecCb
	agent  CObjT
}

func (a Anim) Destroy() {
	var ri int

	for i, v := range anims {
		if v.CObj == a.CObj {
			ri = i
			break
		}
	}
	a.agent = nil
	anims = append(anims[:ri], anims[ri+1:]...)
	// anims = anims[:copy(anims, anims[ri:])]
}

var anims []*Anim

func CreateAnim(o CAnimT) *Anim {
	a := Anim{
		set: set[CAnimT]{
			CObj: o,
		},
	}
	pa := &a
	anims = append(anims, pa)

	return pa
}

//export _execAnimCb
func _execAnimCb(obj unsafe.Pointer, temp C.int) {
	_o := (CObjT)(obj)

	for _, o := range anims {
		if o.agent == _o {
			o.execCb(obj, int32(temp))
		}
	}
}

func (setter *Anim) Var(creater lib.CreateI) *Anim {
	_o := getParentObj(creater)
	(*setter).agent = (CObjT)(_o)
	C.lv_anim_set_var(setter.CObj, _o)

	return setter
}

func (setter *Anim) ExecCb(cb AnimExecCb) *Anim {
	(*setter).execCb = cb
	C.lv_anim_set_exec_cb(setter.CObj, C.lv_anim_exec_xcb_t(C._execAnimCb))
	return setter
}

func (setter *Anim) Time(duration uint32) *Anim {
	C.lv_anim_set_time(setter.CObj, C.uint(duration))
	return setter
}
func (setter *Anim) PlaybackTime(time uint32) *Anim {
	C.lv_anim_set_playback_time(setter.CObj, C.uint(time))
	return setter
}
func (setter *Anim) Values(start, end int32) *Anim {
	C.lv_anim_set_values(setter.CObj, C.int(start), C.int(end))

	return setter
}
func (setter *Anim) RepeatCount(cnt lib.C_SHORT) *Anim {
	C.lv_anim_set_repeat_count(setter.CObj, C.ushort(cnt))
	return setter
}
func (setter *Anim) Start() *Anim {
	C.lv_anim_start(setter.CObj)

	return setter
}

// func (setter *Anim) Delay( uint32_t delay) Anim{
// 	return setter
// }
// func (setter *Anim) CustomExecCb( lv_anim_custom_execCb_t execCb) Anim{
// 	return setter
// }
// func (setter *Anim) PathCb( lv_anim_pathCb_t pathCb) Anim{
// 	return setter
// }
// func (setter *Anim) StartCb( lv_anim_startCb_t startCb) Anim{
// 	return setter
// }
// func (setter *Anim) GetValueCb( lv_anim_get_valueCb_t get_valueCb) Anim{
// 	return setter
// }
// func (setter *Anim) ReadyCb( lv_anim_readyCb_t readyCb) Anim{
// 	return setter
// }
// func (setter *Anim) DeletedCb( lv_anim_deletedCb_t deletedCb) Anim{
// 	return setter
// }
// func (setter *Anim) PlaybackDelay( uint32_t delay) Anim{
// 	return setter
// }
// func (setter *Anim) RepeatDelay( uint32_t delay) Anim{
// 	return setter
// }
// func (setter *Anim) EarlyApply( bool en) Anim{
// 	return setter
// }
// func (setter *Anim) UserData( void * user_data) Anim{
// 	return setter
// }
