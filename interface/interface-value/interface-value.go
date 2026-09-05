package main


import "fmt"



//Declare a interface type 

type Employee interface{
	

}

func main(){


	// store string value in name variable type of Employee 
	var name Employee = "Maa!"


	// store int value in age variable type of Employee 
	var age Employee = 30 

	
	fmt.Printf("Type of name and age respectively %T %T\n", name,age)

 	fmt.Printf("Value of name and age respectively %v %v\n", name,age)



}