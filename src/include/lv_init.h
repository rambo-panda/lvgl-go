#include "lvgl.h"
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

#define LV_LVGL_H_INCLUDE_SIMPLE

#define METHOD_NAME_PRE  "lv_"

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

lv_obj_t *Create(char *t, lv_obj_t *parent);

void lv_ready();
void lv_task_handler2(uint32_t ms);
// extern const int lv_disp_def_refr_period;
extern const unsigned lv_size_content;

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