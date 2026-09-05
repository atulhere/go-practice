package main

import "fmt"

type sum int

// Define a medthod on named type s of type integer
func (s sum) display(number int) {

	fmt.Printf("Type of s is %T ", s)

	fmt.Println("Inside display and value of the number is  : ", number)

}

type Example struct {
	s string

	n int
}

// Define a medthod on named type ex of type struct

func (ex Example) display() {

	//Print the values

	fmt.Println(ex.s, ex.n)

}

func main() {

	var s sum
	number := 30

	s.display(number)

	// Now  create a instance of Example type

	var ex Example = Example{"Maa", 30}

	ex.display()

}
