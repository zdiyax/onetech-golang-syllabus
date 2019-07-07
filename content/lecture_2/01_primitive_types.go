package main

import (
	"fmt"
	_ "reflect"
)

var (
	i int //

	i8   int8
	i16  int16
	i32  int32
	i64  int64
	ui32 uint32 //unsigned integer

	b bool

	f float32 // float64
	s string
	r rune

	ii interface{}
)

func main() {
	//format your code using this shortcut
	//Ctrl + Alt + L
	// Cmd + Alt + L

	i32 = 1<<31 - 1
	ui32 = 1<<32 - 1

	var array = []int{5, 2, 2, 4}

	fmt.Printf("this is my integer array : %v", array)
	fmt.Println("this is my unsigned int:", ui32)

	//slice
	var runes []rune

	for i := 0; i < 100; i++ {
		runes = append(runes, 's')
	}

	s := "Almas"

	fmt.Println([]rune(s))

	var arr1 = []int{1, 3, -31, 5, 0}

	fmt.Println("initial capacity is : ", cap(arr1))

	i := 0
	cap0 := cap(arr1)
	//Ctrl + Shift + V - вытянуть конкретные данные из буфера
	for i < 1000000 {
		i++
		arr1 = append(arr1, 5)

		if cap(arr1) != cap0 {
			fmt.Printf("current capacity is : %1v\n", cap(arr1))
			fmt.Printf("current capacity is : %4.2f\n", float32(cap(arr1))/float32(cap0))
		}

		cap0 = cap(arr1)
	}

	jar := 1233

	//Ctrl + Space (Пробел) - предложение аргументов, функций, пакетов
	fmt.Printf("jar is: %1000v", jar)
}
