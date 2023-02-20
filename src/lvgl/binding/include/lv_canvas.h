#ifndef LV_CANVAS_LIB
#define LV_CANVS_LIB
#define _DEFAULT_SOURCE /* needed for usleep() */

#include "lvgl.h"
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

lv_obj_t *createCanvas(int nowAct);
void lvCanvasReady();
static inline void handlerJob()
{
    while (1)
    {
        lv_timer_handler_run_in_period(LV_DISP_DEF_REFR_PERIOD);
        usleep(LV_DISP_DEF_REFR_PERIOD * 1e3);
    }
};

#if USE_SDL == 1
#include "sdl/sdl.h"
#define LV_17_DISP_INIT sdl_init()
#define LV_17_VER_RES SDL_VER_RES
#define LV_17_HOR_RES SDL_HOR_RES
#define LV_17_FLUSH_CB sdl_display_flush
#elif USE_FBDEV || USE_BSD_FBDEV
#include "display/fbdev.h"
#define LV_17_DISP_INIT fbdev_init()
#define LV_17_HOR_RES LV_HOR_RES
#define LV_17_VER_RES LV_VER_RES
#define LV_17_FLUSH_CB fbdev_flush
#endif

#endif
