package main

import (
	"fmt"
	"golang-week-8/student"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println("test")

	// reflect.ValueOf()
	// reflect.TypeOf()

	var number = 23

	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variable: ", reflectValue.Type())

	// if reflectValue.Kind() == reflect.Int {
	// 	fmt.Println("nilai variable :", reflectValue.Int())
	// }

	// * type assertion
	fmt.Println("nilai variable :", reflectValue.Interface().(int))

	fmt.Println("===============================")

	// * A.29.2
	var s1 = &student.Student{Name: "john wick", Grade: 2}

	// s1.GetPropertyInfo()

	fmt.Println("name:", s1.Name)

	var reflectValue1 = reflect.ValueOf(s1)

	var method = reflectValue1.MethodByName("SetName")

	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("name:", s1.Name)

	fmt.Println("===============================")

	// * A.30
	runtime.GOMAXPROCS(1)

	// go print(5, "halo")
	// print(5, "apa kabar")

	// var input string
	// fmt.Scanln(&input)

	fmt.Println("===============================")

	// * A.31

	var messages = make(chan string)

	// var sayHello = func(who string) {
	// 	var data = fmt.Sprintf("hello %s", who)
	// 	messages <- data
	// }

	// go sayHello("john wick")
	// go sayHello("ethan hunt")
	// go sayHello("jan ethes")

	// var message1 = <-messages
	// fmt.Println("1", message1)
	// var message2 = <-messages
	// fmt.Println("2", message2)
	// var message3 = <-messages
	// fmt.Println("3", message3)

	for _, each := range []string{"wick", "hunt", "ethes"} {
		go func(who string) {
			var data = fmt.Sprintf("hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 3; i++ {
		printMessage(messages)
	}
}

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i + 1), message)
	}
}

func printMessage(what chan string) {
	fmt.Println(<-what)
}
