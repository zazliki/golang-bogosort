package main

import (
	"fmt"

	"github.com/zazliki/golang-bogosort/bogosort"
)

func main() {
	arr := []int{3, 4, 44, 5, 1, 2, 15}
	bogosort.Bogosort(arr)

	fmt.Println(arr)
	fmt.Printf("Steps: %d\n", bogosort.Steps())
	fmt.Printf("Time: %v\n", bogosort.Time())
}
