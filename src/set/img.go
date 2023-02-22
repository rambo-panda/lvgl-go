package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
  lib "lvgl-go/src/lib"
)
func SrcForImg(obj *lib.LvObjT ,src any)  {
   C.lv_img_set_src((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),unsafe.Pointer(&src))

  
}
func OffsetXForImg(obj *lib.LvObjT ,x lib.LvCoordT)  {
   C.lv_img_set_offset_x((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.lv_coord_t(x))

  
}
func OffsetYForImg(obj *lib.LvObjT ,y lib.LvCoordT)  {
   C.lv_img_set_offset_y((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.lv_coord_t(y))

  
}
func AngleForImg(obj *lib.LvObjT ,angle int16)  {
   C.lv_img_set_angle((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.short(angle))

  
}
func PivotForImg(obj *lib.LvObjT ,x lib.LvCoordT,y lib.LvCoordT)  {
   C.lv_img_set_pivot((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.lv_coord_t(x),C.lv_coord_t(y))

  
}
func ZoomForImg(obj *lib.LvObjT ,zoom uint16)  {
   C.lv_img_set_zoom((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.ushort(zoom))

  
}
func AntialiasForImg(obj *lib.LvObjT ,antialias bool)  {
   C.lv_img_set_antialias((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.bool(antialias))

  
}
func SizeModeForImg(obj *lib.LvObjT ,mode lib.LvImgSizeModeT)  {
   C.lv_img_set_size_mode((*C.struct__lv_obj_t)(unsafe.Pointer(obj)),C.lv_img_size_mode_t(mode))

  
}