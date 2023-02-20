/**
 * @file tslib.h
 *
 */

#ifndef TSLIB_DRV_H
#define TSLIB_DRV_H

#ifdef __cplusplus
extern "C" {
#endif

/*********************
 *      INCLUDES
 *********************/
#ifndef LV_DRV_NO_CONF
#ifdef LV_CONF_INCLUDE_SIMPLE
#include "lv_drv_conf.h"
#else
#include "../../lv_drv_conf.h"
#endif
#endif

#if USE_TSLIB

#ifdef LV_LVGL_H_INCLUDE_SIMPLE
#include "lvgl.h"
#else
#include "lvgl/lvgl.h"
#endif

/*********************
 *      DEFINES
 *********************/

/**********************
 *      TYPEDEFS
 **********************/

/**********************
 * GLOBAL PROTOTYPES
 **********************/

/**
 * Initialize the tslib
 */
void tslib_init(void);
/**
 * reconfigure the device file for tslib
 * @param dev_name set the tslib device filename
 * @return true: the device file set complete
 *         false: the device file doesn't exist current system
 */
bool tslib_set_file(char* dev_name);
/**
 * Get the current position and state of the tslib
 * @param data store the tslib data here
 */
void tslib_read(lv_indev_drv_t * drv, lv_indev_data_t * data);


/**********************
 *      MACROS
 **********************/

#endif /* USE_TSLIB */

#ifdef __cplusplus
} /* extern "C" */
#endif

#endif /* TSLIB_DRV_H */
