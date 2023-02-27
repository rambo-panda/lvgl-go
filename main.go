package main

import (
	"fmt"
	"lvgl-go/src/create"
	"lvgl-go/src/lib"
	"lvgl-go/src/types"
)

func main() {
	lib.Ready()

	label := create.CreateLabel(nil)
	label.Set.Text("jjjjj")
	create.CreateObj(&label).Set.Align(types.LV_ALIGN_CENTER)

	src := "S:/Users/rambo/work_space/lvgl_tutorial/lvgl-go/build/bin/a.png"
	{
		img := create.CreateImg(nil)
		obj := create.CreateObj(img.GetObj())
		obj.Set.Width(types.LV_SIZE_CONTENT)
		obj.Set.Height(types.LV_SIZE_CONTENT)

		// img.Set.Src(&src)
		img.Set.Src(src)
		obj.Set.Align(types.LV_ALIGN_CENTER)

		j := img.Get.Src()
		fmt.Println("img src -> ", j)
	}

	lib.TaskHandler(0)
}
