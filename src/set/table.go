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

type SetTable set

func (setter SetTable) SetCellValue(row uint16, col uint16, txt string) SetTable {
	C.lv_table_set_cell_value(setter.CStructLvObjT, C.ushort(row), C.ushort(col), C.CString(txt))

	return setter
}
func (setter SetTable) SetRowCnt(row_cnt uint16) SetTable {
	C.lv_table_set_row_cnt(setter.CStructLvObjT, C.ushort(row_cnt))

	return setter
}
func (setter SetTable) SetColCnt(col_cnt uint16) SetTable {
	C.lv_table_set_col_cnt(setter.CStructLvObjT, C.ushort(col_cnt))

	return setter
}
func (setter SetTable) SetColWidth(col_id uint16, w lib.LvCoordT) SetTable {
	C.lv_table_set_col_width(setter.CStructLvObjT, C.ushort(col_id), C.lv_coord_t(w))

	return setter
}
