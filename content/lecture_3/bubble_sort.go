package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var sortThis []int

	n := 400

	for i := 0; i < n; i++ {
		sortThis = append(sortThis, rand.Intn(100))
	}

	fmt.Println(sortThis)

	bubbleSort(sortThis)

	fmt.Println(sortThis)
}

func bubbleSort(values []int) {
	startTime := time.Now()

	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values); j++ {
			if values[j] < values[i] {
				values[j], values[i] = values[i], values[j]
			}
		}
	}

	dif := time.Since(startTime)
	fmt.Println(dif)
}
