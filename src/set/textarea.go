package set

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../lib -llvgl
#include "lv_init.h"
*/
import "C"
import (
	lib "lvgl-go/src/lib"
)

type SetTextarea set

func (setter SetTextarea) SetText(txt string) SetTextarea {
	C.lv_textarea_set_text(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter SetTextarea) SetPlaceholderText(txt string) SetTextarea {
	C.lv_textarea_set_placeholder_text(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter SetTextarea) SetCursorPos(pos int32) SetTextarea {
	C.lv_textarea_set_cursor_pos(setter.CStructLvObjT, C.int(pos))

	return setter
}
func (setter SetTextarea) SetCursorClickPos(en bool) SetTextarea {
	C.lv_textarea_set_cursor_click_pos(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter SetTextarea) SetPasswordMode(en bool) SetTextarea {
	C.lv_textarea_set_password_mode(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter SetTextarea) SetPasswordBullet(bullet string) SetTextarea {
	C.lv_textarea_set_password_bullet(setter.CStructLvObjT, C.CString(bullet))

	return setter
}
func (setter SetTextarea) SetOneLine(en bool) SetTextarea {
	C.lv_textarea_set_one_line(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter SetTextarea) SetAcceptedChars(list string) SetTextarea {
	C.lv_textarea_set_accepted_chars(setter.CStructLvObjT, C.CString(list))

	return setter
}
func (setter SetTextarea) SetMaxLength(num uint32) SetTextarea {
	C.lv_textarea_set_max_length(setter.CStructLvObjT, C.uint(num))

	return setter
}
func (setter SetTextarea) SetInsertReplace(txt string) SetTextarea {
	C.lv_textarea_set_insert_replace(setter.CStructLvObjT, C.CString(txt))

	return setter
}
func (setter SetTextarea) SetTextSelection(en bool) SetTextarea {
	C.lv_textarea_set_text_selection(setter.CStructLvObjT, C.bool(en))

	return setter
}
func (setter SetTextarea) SetPasswordShowTime(time uint16) SetTextarea {
	C.lv_textarea_set_password_show_time(setter.CStructLvObjT, C.ushort(time))

	return setter
}
func (setter SetTextarea) SetAlign(align lib.LvTextAlignT) SetTextarea {
	C.lv_textarea_set_align(setter.CStructLvObjT, C.lv_text_align_t(align))

	return setter
}
