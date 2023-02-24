package main

import (
	lvgl_go "lvgl-go/src"
	"lvgl-go/src/lib"
)

func main() {
	lib.Ready()
	label := lvgl_go.CreateLabel(nil)
	label.Set.Text("jjjjj")
	// lvgl_go.Demo()
	lib.TaskHandler(0)
}
