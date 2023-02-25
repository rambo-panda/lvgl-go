package main

import (
	"lvgl-go/src/create"
	"lvgl-go/src/lib"
)

func main() {
	lib.Ready()
	label := create.CreateLabel(nil)
	label.Set.Text("jjjjj")
	// lvgl_go.Demo()
	lib.TaskHandler(0)
}
