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

	lib.TaskHandler(0)
}

func a() {
	//   static lv_style_t style_bg;
	// static lv_style_t style_indic;
	styleBg := create.CreateStyle()
	styleIndic := create.CreateStyle()

	styleBg.Set.BorderColor(0xCC00FF).BorderWidth(2).Pad(6).Radius(6).AnimTime(1000)
	styleIndic.Set.BgColor(0x0000FF).Opa(lib.LV_OPA_COVER).Radius(3)

	bar := 

	// lv_obj_t * bar = lv_bar_create(lv_scr_act());
	// lv_obj_remove_style_all(bar);  /*To have a clean start*/
	// lv_obj_add_style(bar, &style_bg, 0);
	// lv_obj_add_style(bar, &style_indic, LV_PART_INDICATOR);

	// lv_obj_set_size(bar, 200, 20);
	// lv_obj_center(bar);
	// lv_bar_set_value(bar, 100, LV_ANIM_ON);

}
