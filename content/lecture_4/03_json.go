package main

import (
	"encoding/json"
	"fmt"
)

// Где используются json на нашем проекте:
// 1. в http запросах
// 2. в системе очередей RabbitMQ
// 3. в нерелеационной бд Elasticsearch
// 4. в некоторых таблицах CassandraDB также хранятся json-ки
//
func main() {
	car := Car{
		EnginePower: 92,
		Price:       320000.0,
		Wheels: []Wheel{
			{Diameter: 32.0, Material: "steel"},
			{Diameter: 32.0, Material: "steel"},
			{Diameter: 32.0, Material: "steel"},
			{Diameter: 32.0, Material: "steel"},
		},
	}

	fmt.Println(car)

	bytes, err := json.MarshalIndent(car, "", "    ")
	if err != nil {
		_ = fmt.Errorf("i got an error and it's huge : %v", err)
	}

	car1 := Car{}

	json.Unmarshal(bytes, &car1)

	// Shift + F6 - реформаттинг переменных, структур, функций глобальноx

	fmt.Println(car1)
}

type Car struct {
	Wheels      []Wheel
	EnginePower int
	Price       float32
}

type Wheel struct {
	Diameter float32
	Material string
}
