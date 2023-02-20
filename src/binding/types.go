package lvgl_go

/*
#include "lv_init.h"
*/
import "C"

type LvObj C.struct__lv_obj_t
type LvSpan C.struct__lv_span_t

const LV_ALIGN_CENTER C.uchar = C.LV_ALIGN_CENTER
