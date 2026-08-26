package main

import "fmt"


//Define a new type Employee of struct  

type Employee struct{
	Name string
 	Age int

}


func main(){

var emp Employee = Employee{"Atul", 39}

fmt.Println(emp)


}