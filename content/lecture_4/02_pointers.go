package main

import "fmt"

func main() {
	a := 5

	//адрес в памяти - & (амперсанд)
	p := &a

	pp := &p

	fmt.Printf("value of pp : %v | type of pp : %T\n", pp, pp)

	fmt.Printf("value of a : %v | type of a : %T\n", a, a)
	fmt.Printf("value of p : %v | type of p : %T\n", p, p)

	//dereference - по адресу памяти получаете значение переменной

	//дай мне значение по адресу p
	b := *p //b = 5

	fmt.Printf("value of b : %v | type of b : %T\n", b, b)

	multiplyIntBy(3, &a)

	fmt.Printf("value of a : %v | type of a : %T\n", a, a)
}

//go is most of the times PASS BY VALUE
func multiplyIntBy(multiplier int, i *int) {
	//i - копия из функции вызова
	//i = i * multiplier
	fmt.Printf("value of a : %v | type of a : %T\n", i, i)

	*i *= multiplier
	//значение по этому адресу = значение по этому адресу * multiplier
}
