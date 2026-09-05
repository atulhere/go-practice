package main 


import "fmt"

func makeCounter() func() int {

	count :=0

	return func() int{

		count++

		return
	}

}


func main (){


// Intitialise function closures

c :=makeCounter()

// calling closures
fmt.Println("First Call to counter: ", c())
fmt.Println("Second Call to counter: ", c())
fmt.Println("Third Call to counter: ", c())
fmt.Println("Fourth Call to counter: ", c())
fmt.Println("Fifth Call to counter: ", c())


}


