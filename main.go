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
	create.CreateObj(label.GetObj()).Set.Align(types.LV_ALIGN_CENTER)

	{
		img := create.CreateImg(nil)
		obj := create.CreateObj(img.GetObj())
		obj.Set.Width(types.LV_SIZE_CONTENT)
		obj.Set.Height(types.LV_SIZE_CONTENT)

		src := "S:lvgl/examples/libs/png/a.png"
		// img.Set.Src(&src)
		img.Set.Src(src)
		obj.Set.Align(types.LV_ALIGN_CENTER)

		j := img.Get.Src()
		fmt.Println("img src -> ", j)
	}

	// lv_obj_set_width(img, LV_SIZE_CONTENT);
	// lv_obj_set_height(img, LV_SIZE_CONTENT);
	// // lv_obj_set_align(img, LV_ALIGN_BOTTOM_LEFT);
	// // const J = LV_SIZE_CONTENT;
	// // LV_IMG_CF_TRUE_COLOR_ALPHA
	// lv_img_set_src(img, "S:lvgl/examples/libs/png/a.png");
	// // lv_img_set_src(img, "S:xxx.png");

	// lv_obj_center(img);

	lib.TaskHandler(0)
}
