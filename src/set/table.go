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

type Table set

func CreateTable(o *types.LvObjT) Table {
	return Table{
		CStructLvObjT: (*C.struct__lv_obj_t)(unsafe.Pointer(o)),
	}
}

func (setter Table) GetObj() *types.LvObjT {
	return (*types.LvObjT)(unsafe.Pointer(setter.CStructLvObjT))
}

func (setter Table) CellValue(row uint16, col uint16, txt string) Table {
	C.lv_table_set_cell_value(setter.CStructLvObjT, C.ushort(row), C.ushort(col), C.CString(txt))

	return setter
}
func (setter Table) RowCnt(row_cnt uint16) Table {
	C.lv_table_set_row_cnt(setter.CStructLvObjT, C.ushort(row_cnt))

	return setter
}
func (setter Table) ColCnt(col_cnt uint16) Table {
	C.lv_table_set_col_cnt(setter.CStructLvObjT, C.ushort(col_cnt))

	return setter
}
func (setter Table) ColWidth(col_id uint16, w types.LvCoordT) Table {
	C.lv_table_set_col_width(setter.CStructLvObjT, C.ushort(col_id), C.lv_coord_t(w))

	return setter
}
