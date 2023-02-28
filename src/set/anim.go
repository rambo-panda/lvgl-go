package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
extern void TestForAnim(void * obj, int temp);
*/
import "C"
import (
	"lvgl-go/src/lib"
	"unsafe"
)

type Anim set[CAnimT]

func CreateAnim(o CAnimT) Anim {
	return Anim{
		CObj: o,
	}
}

func a() {}

//export TestForAnim
func TestForAnim(obj unsafe.Pointer, temp C.int) {
	C.lv_bar_set_value((CObjT)(obj), temp, C.LV_ANIM_ON)
	a()
}

func (setter Anim) Var(creater any) Anim {
	C.lv_anim_set_var(setter.CObj, getParentObj(creater))

	return setter
}
func (setter Anim) ExecCb() Anim {
	// f := C.TestForAnim

	// 首地址
	// C.lv_anim_set_exec_cb(setter.CObj, (*[0]byte)(f))
	C.lv_anim_set_exec_cb(setter.CObj, C.lv_anim_exec_xcb_t(C.TestForAnim))
	return setter
}

func (setter Anim) Time(duration uint32) Anim {
	C.lv_anim_set_time(setter.CObj, C.uint(duration))
	return setter
}
func (setter Anim) PlaybackTime(time uint32) Anim {
	C.lv_anim_set_playback_time(setter.CObj, C.uint(time))
	return setter
}
func (setter Anim) Values(start, end int32) Anim {
	C.lv_anim_set_values(setter.CObj, C.int(start), C.int(end))

	return setter
}
func (setter Anim) RepeatCount(cnt lib.C_SHORT) Anim {
	C.lv_anim_set_repeat_count(setter.CObj, C.ushort(cnt))
	return setter
}
func (setter Anim) Start() Anim {
	C.lv_anim_start(setter.CObj)

	return setter
}

// func (setter Anim) Delay( uint32_t delay) Anim{
// 	return setter
// }
// func (setter Anim) CustomExecCb( lv_anim_custom_execCb_t execCb) Anim{
// 	return setter
// }
// func (setter Anim) PathCb( lv_anim_pathCb_t pathCb) Anim{
// 	return setter
// }
// func (setter Anim) StartCb( lv_anim_startCb_t startCb) Anim{
// 	return setter
// }
// func (setter Anim) GetValueCb( lv_anim_get_valueCb_t get_valueCb) Anim{
// 	return setter
// }
// func (setter Anim) ReadyCb( lv_anim_readyCb_t readyCb) Anim{
// 	return setter
// }
// func (setter Anim) DeletedCb( lv_anim_deletedCb_t deletedCb) Anim{
// 	return setter
// }
// func (setter Anim) PlaybackDelay( uint32_t delay) Anim{
// 	return setter
// }
// func (setter Anim) RepeatDelay( uint32_t delay) Anim{
// 	return setter
// }
// func (setter Anim) EarlyApply( bool en) Anim{
// 	return setter
// }
// func (setter Anim) UserData( void * user_data) Anim{
// 	return setter
// }
