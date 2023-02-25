package main

import (
	"lvgl-go/src/create"
	"lvgl-go/src/lib"
	"lvgl-go/src/types"
)

func main() {
	lib.Ready()
	label := create.CreateLabel(nil)
	label.Set.Text("jjjjj")
	create.CreateObj(label).Set.Align(types.LV_ALIGN_CENTER)
	// lvgl_go.Demo()
	lib.TaskHandler(0)
}
