package set

import types "lvgl-go/src/types"

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	"unsafe"
)

type Textarea set

func CreateTextarea(o *types.LvObjT) Textarea {
	return Textarea{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Textarea) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter Textarea) Text(txt string) Textarea {
	C.lv_textarea_set_text(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter Textarea) PlaceholderText(txt string) Textarea {
	C.lv_textarea_set_placeholder_text(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter Textarea) CursorPos(pos int32) Textarea {
	C.lv_textarea_set_cursor_pos(setter.CStructLvObjT, C.int(pos))

	return setter
}
func (setter Textarea) CursorClickPos(en bool) Textarea {
	C.lv_textarea_set_cursor_click_pos(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter Textarea) PasswordMode(en bool) Textarea {
	C.lv_textarea_set_password_mode(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter Textarea) PasswordBullet(bullet string) Textarea {
	C.lv_textarea_set_password_bullet(setter.CStructLvObjT, C.CString(bullet))

	return setter
}
func (setter Textarea) OneLine(en bool) Textarea {
	C.lv_textarea_set_one_line(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter Textarea) AcceptedChars(list string) Textarea {
	C.lv_textarea_set_accepted_chars(setter.CStructLvObjT, C.CString(list))

	return setter
}
func (setter Textarea) MaxLength(num uint32) Textarea {
	C.lv_textarea_set_max_length(setter.CStructLvObjT, C.uint(num))

	return setter
}
func (setter Textarea) InsertReplace(txt string) Textarea {
	C.lv_textarea_set_insert_replace(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter Textarea) TextSelection(en bool) Textarea {
	C.lv_textarea_set_text_selection(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter Textarea) PasswordShowTime(time uint16) Textarea {
	C.lv_textarea_set_password_show_time(setter.CStructLvObjT, C.ushort(time))

	return setter
}
func (setter Textarea) Align(align types.LvTextAlignT) Textarea {
	C.lv_textarea_set_align(setter.CStructLvObjT, C.lv_text_align_t(align))

	return setter
}
