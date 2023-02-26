package set

import (
	types "lvgl-go/src/types"
)

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"

type Textarea set

func CreateTextarea(o CObjT) Textarea {
	return Textarea{
		CObj: o,
	}
}

func (setter Textarea) Text(txt string) Textarea {
	C.lv_textarea_set_text(setter.CObj, C.CString(txt))

	return setter
}
func (setter Textarea) PlaceholderText(txt string) Textarea {
	C.lv_textarea_set_placeholder_text(setter.CObj, C.CString(txt))

	return setter
}
func (setter Textarea) CursorPos(pos int32) Textarea {
	C.lv_textarea_set_cursor_pos(setter.CObj, C.int(pos))

	return setter
}
func (setter Textarea) CursorClickPos(en bool) Textarea {
	C.lv_textarea_set_cursor_click_pos(setter.CObj, C.bool(en))

	return setter
}
func (setter Textarea) PasswordMode(en bool) Textarea {
	C.lv_textarea_set_password_mode(setter.CObj, C.bool(en))

	return setter
}
func (setter Textarea) PasswordBullet(bullet string) Textarea {
	C.lv_textarea_set_password_bullet(setter.CObj, C.CString(bullet))

	return setter
}
func (setter Textarea) OneLine(en bool) Textarea {
	C.lv_textarea_set_one_line(setter.CObj, C.bool(en))

	return setter
}
func (setter Textarea) AcceptedChars(list string) Textarea {
	C.lv_textarea_set_accepted_chars(setter.CObj, C.CString(list))

	return setter
}
func (setter Textarea) MaxLength(num uint32) Textarea {
	C.lv_textarea_set_max_length(setter.CObj, C.uint(num))

	return setter
}
func (setter Textarea) InsertReplace(txt string) Textarea {
	C.lv_textarea_set_insert_replace(setter.CObj, C.CString(txt))

	return setter
}
func (setter Textarea) TextSelection(en bool) Textarea {
	C.lv_textarea_set_text_selection(setter.CObj, C.bool(en))

	return setter
}
func (setter Textarea) PasswordShowTime(time uint16) Textarea {
	C.lv_textarea_set_password_show_time(setter.CObj, C.ushort(time))

	return setter
}
func (setter Textarea) Align(align types.LvTextAlignT) Textarea {
	C.lv_textarea_set_align(setter.CObj, C.lv_text_align_t(align))

	return setter
}
