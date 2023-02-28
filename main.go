package main

import (
	"fmt"
	"lvgl-go/src/create"
	"lvgl-go/src/lib"
	"lvgl-go/src/types"
)

func main() {
	lib.Ready()

	label := create.CreateLabel(&create.CREATE_NIL)
	label.Set.Text("jjjjj")
	create.CreateObj(&label).Set.Align(types.LV_ALIGN_CENTER)

	src := "S:/Users/rambo/work_space/lvgl_tutorial/lvgl-go/build/bin/a.png"
	{
		img := create.CreateImg(&create.CREATE_NIL)
		obj := create.CreateObj(&img)
		obj.Set.Width(types.LV_SIZE_CONTENT)
		obj.Set.Height(types.LV_SIZE_CONTENT)

		// img.Set.Src(&src)
		img.Set.Src(src)
		obj.Set.Align(types.LV_ALIGN_CENTER)

		j := img.Get.Src()
		fmt.Println("img src -> ", j)
	}

	a()

	lib.TaskHandler(0)
}

func a() {
	styleBg := create.CreateStyle()
	styleIndic := create.CreateStyle()

	styleBg.Set.BorderColor(0xCC00FF).BorderWidth(2).Pad(6).Radius(6).AnimTime(1000)
	styleIndic.Set.BgColor(0x0000FF).Opa(lib.LV_OPA_COVER).Radius(3)

	bar := create.CreateBar(&create.CREATE_NIL)
	create.CreateObj(&bar).Set.Style(&styleBg, 0).Style(&styleIndic, int(lib.LV_PART_INDICATOR)).Size(200, 20).Center()
	bar.Set.Value(100, types.LV_ANIM_ON)

}

func b() {
	styleIndic := create.CreateStyle()
	styleIndic.Set.BgOpa(lib.LV_OPA_COVER).BgColor(0xFFC0CB).BgGradColor(0x7B68EE).BgGradDir(lib.LV_PART_INDICATOR)

	bar := create.CreateBar(&create.CREATE_NIL)
	create.CreateObj(&bar).Set.Style(&styleIndic, lib.LV_PART_INDICATOR).Size(20, 200).Center()
	bar.Set.Range(-20, 40)

	// lv_anim_t a;
	// lv_anim_init(&a);
	// lv_anim_set_exec_cb(&a, set_temp);
	// lv_anim_set_time(&a, 3000);
	// lv_anim_set_playback_time(&a, 3000);
	// lv_anim_set_var(&a, bar);
	// lv_anim_set_values(&a, -20, 40);
	// lv_anim_set_repeat_count(&a, LV_ANIM_REPEAT_INFINITE);
	// lv_anim_start(&a);
}
