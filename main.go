package main

import (
	"fmt"
	lvgl_go "lvgl-go/src/binding"
	"lvgl-go/src/binding/lib"
	"reflect"
)

type xFoo struct{}

func (f xFoo) HHH() int {
	fmt.Printf("0000000000000")
	return 1
}

func main() {

	foo := &xFoo{}

	ret := reflect.ValueOf(foo).MethodByName("HHH").Call([]reflect.Value{})
	println("ret %v", ret[0].Interface().(int))

	return
	// funcValue := reflect.ValueOf("add")
	// paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// retList := funcValue.Call(paramList)
	// fmt.Println(retList[0].Int())

	lib.Ready()
	lvgl_go.Demo()
	lib.TaskHandler(0)
}

// package main

// import (
//     "fmt"
//     "reflect"
// )

// func main() {

//     // 将函数包装为反射值对象
//     funcValue := reflect.ValueOf(add)

//     // 构造函数参数, 传入两个整型值
//     paramList := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}

//     // 反射调用函数
//     retList := funcValue.Call(paramList)

//     // 获取第一个返回值, 取整数值
//     fmt.Println(retList[0].Int())
// }
