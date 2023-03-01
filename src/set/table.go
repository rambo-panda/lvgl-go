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

type Table set[CObjT]

func CreateTable(o CObjT) Table {
	return Table{
		CObj: o,
	}
}

func (setter Table) CellValue(row uint16, col uint16, txt string) Table {
	C.lv_table_set_cell_value(setter.CObj, C.ushort(row), C.ushort(col), C.CString(txt))

	return setter
}
func (setter Table) RowCnt(row_cnt uint16) Table {
	C.lv_table_set_row_cnt(setter.CObj, C.ushort(row_cnt))

	return setter
}
func (setter Table) ColCnt(col_cnt uint16) Table {
	C.lv_table_set_col_cnt(setter.CObj, C.ushort(col_cnt))

	return setter
}
func (setter Table) ColWidth(col_id uint16, w lib.LV_COLOR_T) Table {
	C.lv_table_set_col_width(setter.CObj, C.ushort(col_id), C.lv_coord_t(w))

	return setter
}
