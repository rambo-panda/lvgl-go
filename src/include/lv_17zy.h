#ifndef LVGL_GO
#define LVGL_GO

#include "lvgl.h"
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define COUNT_ARGS(...) (sizeof((int[]){0, ##__VA_ARGS__}) / sizeof(int) - 1)

#define GEN_FN(func, ...) (func(COUNT_ARGS(__VA_ARGS__), __VA_ARGS__))
// #define gen_fn(func, ...) (#func(COUNT_ARGS(__VA_ARGS__), __VA_ARGS__))

#define ARGS(count, __type)                                                    \
  __type args[count];                                                          \
  va_list _args;                                                               \
  va_start(_args, count);                                                      \
  for (int X_I = 0; X_I < count; X_I++) {                                      \
    args[X_I] = va_arg(_args, __type);                                         \
  }                                                                            \
  va_end(_args)

void lv_ready();
void lv_task_handler2(uint32_t ms);
// extern const int lv_disp_def_refr_period;
extern const unsigned lv_size_content;

#include "display/fbdev.h"
/*
 * 用于lvgl初始化时注册的函数接口，由于不需要接受输入，只实现了显示部分。
 * static lv_disp_drv_t disp_drv;
 * lv_disp_drv_init(&disp_drv);
 * disp_drv.flush_cb = lvdrv_flush;
 * ......
 * ......
 * lv_disp_drv_register(&disp_drv);
*/

void lvdrv_init();

void lvdrv_flush(lv_disp_drv_t * drv, const lv_area_t * area, lv_color_t * color_p);

void lvdrv_deinit();

void tslib_read();

#define LV_17_HOR_RES 128
#define LV_17_VER_RES 32
#define LV_17_FLUSH_CB lvdrv_flush

#endif
