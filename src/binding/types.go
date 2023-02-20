package lvgl_go

/*
#include "lv_init.h"
static const int _LV_SIZE_CONTENT = LV_SIZE_CONTENT;
*/
import "C"

type LvObj C.struct__lv_obj_t
type LvSpan C.struct__lv_span_t

const (
	LV_SIZE_CONTENT C.int = C._LV_SIZE_CONTENT

	LV_ALIGN_CENTER C.int = C.LV_ALIGN_CENTER
)
