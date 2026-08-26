package main



import "fmt"

//Declare a interface type 

type Employee interface{
	

}

func main(){

	
	var name Employee = "Maa!"


	//overide the value as the name is type of Empty Iterface
	name =  30

	fmt.Println(name)




}