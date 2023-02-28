// https://docs.lvgl.io/master/overview/color.html?highlight=lv_palette_blue
package lib

/*
#include "lv_init.h"
*/
import "C"

type Color uint8

const (
	HEX  Color = 0
	HEX3 Color = 1

	RGB Color = 2

	//h = 0..359, s = 0..100, v = 0..100
	// lv_color_t c = lv_color_hsv_to_rgb(h, s, v);
	HSV Color = 3
	// lv_color_hsv_t c_hsv = lv_color_rgb_to_hsv(r, g, b);
	TO_HSV Color = 4
	// lv_color_hsv_t c_hsv = lv_color_to_hsv(color);
	TO_HSV2 Color = 5

	// https://docs.lvgl.io/master/overview/color.html?highlight=lv_palette_blue#palette
	PALETTE Color = 6

	WHITE Color = 7
	BLACK Color = 8

	LIGHTEN Color = 9
	DARKEN  Color = 10
)

const (
	LV_PALETTE_PINK        LV_PALETTE_T = C.LV_PALETTE_PINK
	LV_PALETTE_PURPLE      LV_PALETTE_T = C.LV_PALETTE_PURPLE
	LV_PALETTE_DEEP_PURPLE LV_PALETTE_T = C.LV_PALETTE_DEEP_PURPLE
	LV_PALETTE_INDIGO      LV_PALETTE_T = C.LV_PALETTE_INDIGO
	LV_PALETTE_BLUE        LV_PALETTE_T = C.LV_PALETTE_BLUE
	LV_PALETTE_LIGHT_BLUE  LV_PALETTE_T = C.LV_PALETTE_LIGHT_BLUE
	LV_PALETTE_CYAN        LV_PALETTE_T = C.LV_PALETTE_CYAN
	LV_PALETTE_TEAL        LV_PALETTE_T = C.LV_PALETTE_TEAL
	LV_PALETTE_GREEN       LV_PALETTE_T = C.LV_PALETTE_GREEN
	LV_PALETTE_LIGHT_GREEN LV_PALETTE_T = C.LV_PALETTE_LIGHT_GREEN
	LV_PALETTE_LIME        LV_PALETTE_T = C.LV_PALETTE_LIME
	LV_PALETTE_YELLOW      LV_PALETTE_T = C.LV_PALETTE_YELLOW
	LV_PALETTE_AMBER       LV_PALETTE_T = C.LV_PALETTE_AMBER
	LV_PALETTE_ORANGE      LV_PALETTE_T = C.LV_PALETTE_ORANGE
	LV_PALETTE_DEEP_ORANGE LV_PALETTE_T = C.LV_PALETTE_DEEP_ORANGE
	LV_PALETTE_BROWN       LV_PALETTE_T = C.LV_PALETTE_BROWN
	LV_PALETTE_BLUE_GREY   LV_PALETTE_T = C.LV_PALETTE_BLUE_GREY
	LV_PALETTE_GREY        LV_PALETTE_T = C.LV_PALETTE_GREY
)

const (
	LV_OPA_TRANSP LV_OPA_T = C.LV_OPA_TRANSP
	LV_OPA_0      LV_OPA_T = C.LV_OPA_0
	LV_OPA_10     LV_OPA_T = C.LV_OPA_10
	LV_OPA_20     LV_OPA_T = C.LV_OPA_20
	LV_OPA_30     LV_OPA_T = C.LV_OPA_30
	LV_OPA_40     LV_OPA_T = C.LV_OPA_40
	LV_OPA_50     LV_OPA_T = C.LV_OPA_50
	LV_OPA_60     LV_OPA_T = C.LV_OPA_60
	LV_OPA_70     LV_OPA_T = C.LV_OPA_70
	LV_OPA_80     LV_OPA_T = C.LV_OPA_80
	LV_OPA_90     LV_OPA_T = C.LV_OPA_90
	LV_OPA_100    LV_OPA_T = C.LV_OPA_100
	LV_OPA_COVER  LV_OPA_T = C.LV_OPA_COVER
)

func fillSlice(slice *[]any, wishLen int, defaultVal any) []any {
	len := len(*slice)
	argsCopy := make([]any, wishLen, wishLen)
	diff := wishLen - len
	enough := diff <= 0

	for i, v := range (*slice)[0:Ternary[int](enough, wishLen, len)] {
		argsCopy[i] = v
	}

	if enough {
		for i, _ := range make([]any, diff, diff) {
			argsCopy[len+i] = defaultVal
		}
	}

	return argsCopy
}

func mapSlice[T uint8](slice *[]any, mapFn func(v T, i int) T) []T {
	res := make([]T, len(*slice))

	for i, v := range *slice {
		res[i] = mapFn(v.(T), i)
	}

	return res
}

func ToColor(action Color, args ...any) LV_COLOR_T {
	var res LV_COLOR_T

	switch action {
	case RGB:
		argsFill := fillSlice(&args, 3, 0)
		argsCopy := mapSlice(&argsFill, func(v uint8, _ int) uint8 {
			return v
		})

		res = C.lv_color_make(C.uchar(argsCopy[0]), C.uchar(argsCopy[1]), C.uchar(argsCopy[2]))
	case HEX3:
		res = C.lv_color_hex(C.uint(args[0].(uint32)))
	case HEX:
		res = C.lv_color_hex3(C.uint(args[0].(uint32)))
	case PALETTE:
		res = C.lv_palette_main(args[0].(LV_PALETTE_T))
	case WHITE:
		res = C.lv_color_white()
	case BLACK:
		res = C.lv_color_black()
	case LIGHTEN:
		res = C.lv_color_lighten(args[0].(LV_COLOR_T), args[1].(LV_OPA_T))
	case DARKEN:
		res = C.lv_color_darken(args[0].(LV_COLOR_T), args[1].(LV_OPA_T))
	default:
		panic("暂未支持该旅行 " + string(action))
	}

	return res
}
