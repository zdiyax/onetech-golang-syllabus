package main

import "fmt"

//maps in go are hashmaps

func main() {
	m := make(map[string]string)
	//декларация #2
	//var m2 = map[string]string{}

	m["vasya"] = "123213"
	m["kolya"] = "321321"
	m["petya"] = "424242"
	m["egor"] = "000000"

	//foreach
	for key, value := range m {
		fmt.Println(key + " " + value)
	}

	//проверка есть ли значение в map
	if val, ok := m["kolya"]; ok {
		fmt.Println("val is : ", val)
	}

}
