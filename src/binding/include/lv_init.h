#include "lvgl.h"
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define LV_LVGL_H_INCLUDE_SIMPLE

void lv_ready();
void lv_task_handler2(uint32_t ms);
// extern const int lv_disp_def_refr_period;

#if USE_SDL == 1
#include "./sdl/sdl.h"
#define LV_17_DISP_INIT sdl_init()
#define LV_17_VER_RES SDL_VER_RES
#define LV_17_HOR_RES SDL_HOR_RES
#define LV_17_FLUSH_CB sdl_display_flush
#elif USE_FBDEV || USE_BSD_FBDEV
#include "./display/fbdev.h"
#define LV_17_DISP_INIT fbdev_init()
#define LV_17_HOR_RES LV_HOR_RES
#define LV_17_VER_RES LV_VER_RES
#define LV_17_FLUSH_CB fbdev_flush
#endif