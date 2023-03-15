#include "lv_17zy.h"

static void tslib_read(lv_indev_drv_t *drv, lv_indev_data_t *data)
{
}

// TODO: 因为下面static原因，且目前对于多屏(屋里显示器)支持有些多余，暂时不考虑
static void createDisplay()
{
    // lv_disp_set_default // 指定默认display
    // lv_obj_get_disp

    // https://docs.lvgl.io/8.3/porting/display.html
    // XXX: 以下几个static全局作用域的变量，必须设置static属性，因为display不能被回收
    static lv_disp_draw_buf_t disp_buf;

    static lv_color_t buf1[LV_17_HOR_RES * LV_17_VER_RES];
    lv_disp_draw_buf_init(&disp_buf, buf1, NULL, LV_17_HOR_RES * LV_17_VER_RES);

    static lv_disp_drv_t disp_drv;
    lv_disp_drv_init(&disp_drv);
    disp_drv.draw_buf = &disp_buf;
    disp_drv.flush_cb = LV_17_FLUSH_CB;
    disp_drv.hor_res = LV_17_HOR_RES;
    disp_drv.ver_res = LV_17_VER_RES;
    disp_drv.direct_mode = 1;
    // disp_drv.antialiasing = 1;
    disp_drv.rotated = LV_DISP_ROT_NONE;
    disp_drv.sw_rotate = 0;
    disp_drv.full_refresh = 1;
    lv_disp_t *disp = lv_disp_drv_register(&disp_drv);

    // XXX: 我也不清楚，在我们设备上必须加这个设置
    static lv_indev_drv_t indev_drv;
    lv_indev_drv_init(&indev_drv);
    indev_drv.type = LV_INDEV_TYPE_POINTER;
    indev_drv.read_cb = tslib_read;
    lv_indev_drv_register(&indev_drv);


    // // 目前WARN ： lv_draw_sw_letter: lv_draw_letter: glyph dsc. not found for U+4F60/U+597D     (in lv_draw_sw_letter.c line #106)
    // static lv_ft_info_t font_device;
    // font_device.name = "/gb18030_b12.bdf";
    // font_device.weight = 12;
    // font_device.style = FT_FONT_STYLE_NORMAL;
    // lv_ft_font_init(&font_device);
    // lv_theme_t *th = lv_theme_default_init(disp, lv_color_black(), lv_color_white(), 1, font_device.font);

    lv_theme_t *th = lv_theme_default_init(disp, lv_color_black(), lv_color_white(), 1, LV_FONT_DEFAULT);
    lv_disp_set_theme(disp, th);
}

void lv_ready()
{
    if (lv_is_initialized())
    {
        return;
    }

    LV_LOG_INFO("\r\n"
                "LVGL v%d.%d.%d "
                " Benchmark (in csv format)\r\n",
                LVGL_VERSION_MAJOR, LVGL_VERSION_MINOR, LVGL_VERSION_PATCH);

    lv_init();
    lvdrv_init();

    createDisplay();

    lv_fs_stdio_init(); // 盘符目前写死为A: LV_USE_FS_STDIO
    lv_png_init();      // 默认开启png LV_USE_PNG

    lv_freetype_init(64, 1, 0); // 开启freetype 需要使用我们自己的字体 LV_USE_FREETYPE
}

void lv_task_handler2(uint32_t ms)
{
    ms = ms == 0 ? LV_DISP_DEF_REFR_PERIOD : ms;

    while (1)
    {
        lv_timer_handler_run_in_period(ms);
        usleep(ms * 1e3);
    }
}
