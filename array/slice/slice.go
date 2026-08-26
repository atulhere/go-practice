package main

import "fmt"


// An array has a fixed size. 
// A slice, on the other hand, is a dynamically-sized, 
// flexible view into the elements of an array. In practice, 
// slices are much more common than arrays.
// Example of Slice

func main() {

	var array = []int{1, 2, 3, 4, 5}

     slice := array[1:4]

	fmt.Println("Example Slice", slice)

   //Print length and capcity of a slice 

	fmt.Println("length and capicity of the slice are : ", len(slice), cap(slice))
}
