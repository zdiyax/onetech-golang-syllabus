package main

import "fmt"

//todo: доделать (домашка)

func main() {
	st := stack{}

	for i := 0; i < 10; i++ {
		st.push(i)
	}

	fmt.Println(st)
}

type stack struct {
	values []int
}

// Это МЕТОД! Не функция
func (st *stack) push(value int) {
	st.values = append(st.values, value)
}

func (st *stack) pop() int {
	panic("implement me later")
}

func (st *stack) peek() int {
	panic("implement me later")
}
