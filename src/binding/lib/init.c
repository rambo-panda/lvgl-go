#include "lv_init.h"

// TODO: 因为下面static原因，且目前对于多屏(屋里显示器)支持有些多余，暂时不考虑
static lv_disp_t *createDisplay()
{
    // lv_disp_set_default // 指定默认display
    // lv_obj_get_disp

    // https://docs.lvgl.io/8.3/porting/display.html
    // XXX: 以下几个static全局作用域的变量，必须设置static属性，因为display不能被回收
    static lv_disp_draw_buf_t disp_buf;

    static lv_color_t buf_1[10 * LV_17_HOR_RES];
    static lv_color_t buf_2[10 * LV_17_HOR_RES];

    lv_disp_draw_buf_init(&disp_buf, buf_1, buf_2, 10 * LV_17_HOR_RES);

    static lv_disp_drv_t disp_drv;
    lv_disp_drv_init(&disp_drv);
    disp_drv.draw_buf = &disp_buf;
    disp_drv.flush_cb = LV_17_FLUSH_CB;
    disp_drv.hor_res = LV_17_HOR_RES;
    disp_drv.ver_res = LV_17_VER_RES;
    disp_drv.antialiasing = 1;

    lv_disp_t *disp = lv_disp_drv_register(&disp_drv);

    lv_theme_t *th = lv_theme_default_init(
        disp, lv_palette_main(LV_PALETTE_BLUE), lv_palette_main(LV_PALETTE_RED),
        LV_THEME_DEFAULT_DARK, LV_FONT_DEFAULT);
    lv_disp_set_theme(disp, th);

    return disp;
}

void lv_ready()
{
    printf("jjjjj lv_ready");
    return ;
    if (lv_is_initialized())
    {
        return;
    }

    LV_LOG_INFO("\r\n"
                "LVGL v%d.%d.%d "
                " Benchmark (in csv format)\r\n",
                LVGL_VERSION_MAJOR, LVGL_VERSION_MINOR, LVGL_VERSION_PATCH);

    lv_init();
    LV_17_DISP_INIT;

    createDisplay();
}

lv_obj_t *createCanvas(int autoLoad)
{
    lv_obj_t *scr = lv_obj_create(NULL);

    if (autoLoad)
    {
        lv_scr_load(scr);
    }

    // lv_obj_get_screen

    return scr;
}
