package main

import (
	"fmt"
	lvgl_go "lvgl-go/src/binding"
	"lvgl-go/src/binding/lib"
	"os"
)

func add[T int | float64](a T, arg ...T) T {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Fprintln(os.Stderr, "hello world")
			// fmt.Errorf("username %s is already existed", "jjjjjjj")
			// fmt.Println("------------------")
		}
	}()

	var i interface{} = 123
	k := i.(float32)
	fmt.Printf("fdsfsd %d ---", 1+k)
	s := arg[0]
	return a + s
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(123)

	return
	lib.Ready()
	lvgl_go.Demo()
	lib.TaskHandler(0)
}
