package get

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
type GetLabel Get
func (getter GetLabel) GetText() string {
  res :=  C.lv_label_get_text(getter.CStructLvObjT)

  return C.GoString(res)
}
func (getter GetLabel) GetLongMode() lib.LvLabelLongModeT {
  res :=  C.lv_label_get_long_mode(getter.CStructLvObjT)

  return lib.LvLabelLongModeT(res)
}
func (getter GetLabel) GetRecolor() bool {
  res :=  C.lv_label_get_recolor(getter.CStructLvObjT)

  return bool(res)
}
func (getter GetLabel) GetLetterPos(char_id uint32,pos *lib.LvPointT ) GetLabel {
   C.lv_label_get_letter_pos(getter.CStructLvObjT,C.uint(char_id),(*C.lv_point_t)(unsafe.Pointer(pos)))

  return getter
}
func (getter GetLabel) GetLetterOn(pos_in *lib.LvPointT ) uint32 {
  res :=  C.lv_label_get_letter_on(getter.CStructLvObjT,(*C.lv_point_t)(unsafe.Pointer(pos_in)))

  return uint32(res)
}
func (getter GetLabel) GetTextSelectionStart() uint32 {
  res :=  C.lv_label_get_text_selection_start(getter.CStructLvObjT)

  return uint32(res)
}
func (getter GetLabel) GetTextSelectionEnd() uint32 {
  res :=  C.lv_label_get_text_selection_end(getter.CStructLvObjT)

  return uint32(res)
}