// #define LV_USE_GIF 0
// TODO: 头文件暂时不添加对gif的支持
// package set

// /*
// #cgo CFLAGS: -I../include/
// #cgo LDFLAGS: -L../lib -llvgl
// #include "lv_init.h"
// */
// import "C"
// import (
// 	lib "lvgl-go/src/lib"
// 	"unsafe"
// )

// type SetGif set

// func (setter SetGif) SetSrc(src any) SetGif {
// 	C.lv_gif_set_src(setter.cObj, unsafe.Pointer(&src))

// 	return setter
// }
