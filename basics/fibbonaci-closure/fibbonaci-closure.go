package main

import "fmt"

func fibbonaci() func() int {

	prev := 0

	cursor := 1

	next := 0

	return func() int {

		next = prev + cursor

		prev, cursor = cursor,next

		return next 

	}

}

func main() {

	f := fibbonaci()

	for i := 1; i < 10; i++ {

		fmt.Println(f())

	}

}
