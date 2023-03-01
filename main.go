package main

import (
	"fmt"
	"lvgl-go/src/create"
	"lvgl-go/src/lib"
	"unsafe"
)

func main() {
	lib.Ready()

	label := create.CreateLabel(&lib.CREATE_NIL)
	label.Set.Text("jjjjj")
	create.CreateObj(&label).Set.Align(lib.LV_ALIGN_CENTER)

	src := "S:/Users/rambo/work_space/lvgl_tutorial/lvgl-go/build/bin/a.png"
	{
		img := create.CreateImg(&lib.CREATE_NIL)
		obj := create.CreateObj(&img)
		obj.Set.Width(lib.LV_SIZE_CONTENT)
		obj.Set.Height(lib.LV_SIZE_CONTENT)

		// img.Set.Src(&src)
		img.Set.Src(src)
		obj.Set.Align(lib.LV_ALIGN_CENTER)

		j := img.Get.Src()
		fmt.Println("img src -> ", j)
	}

	b()

	lib.TaskHandler(0)
}

func a() {
	styleBg := create.CreateStyle()
	styleIndic := create.CreateStyle()

	styleBg.Set.BorderColor(0xCC00FF).BorderWidth(2).Pad(6).Radius(6).AnimTime(1000)
	styleIndic.Set.BgColor(0x0000FF).Opa(lib.LV_OPA_COVER).Radius(3)

	bar := create.CreateBar(&lib.CREATE_NIL)
	create.CreateObj(&bar).Set.Style(&styleBg, 0).Style(&styleIndic, lib.LV_PART_INDICATOR).Size(200, 20).Center() // TODO: Size控制垂直还是水平
	bar.Set.Value(100, lib.LV_ANIM_ON)

}

func b() {
	styleIndic := create.CreateStyle()
	styleIndic.Set.BgOpa(lib.LV_OPA_COVER).BgColor(0xFFC0CB).BgGradColor(0x7B68EE).BgGradDir(lib.LV_PART_INDICATOR)

	bar := create.CreateBar(&lib.CREATE_NIL)
	create.CreateObj(&bar).Set.Style(&styleIndic, lib.LV_PART_INDICATOR).Size(20, 200).Center()
	bar.Set.Range(-20, 40)

	create.CreateAnim().Set.ExecCb(func(_ unsafe.Pointer, temp int32) {
		bar.Set.Value(temp, lib.LV_ANIM_ON)
	}).Time(3e3).PlaybackTime(3e3).Var(bar).Values(-20, 40).RepeatCount(lib.LV_ANIM_REPEAT_INFINITE).Start()
}
