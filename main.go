package main

import (
	lvgl_go "lvgl-go/src/binding"
	"lvgl-go/src/binding/lib"
)

func main() {
	lib.Ready()
	lvgl_go.Demo()
	lib.TaskHandler(0)
}
