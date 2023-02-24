package main

import (
	lvgl_go "lvgl-go/src"
	"lvgl-go/src/lib"
)

func main() {
	lib.Ready()
	label := lvgl_go.CreateLabel(nil)
	label.Set.SetText("jjjjj")
	// lvgl_go.Demo()
	lib.TaskHandler(0)
}
