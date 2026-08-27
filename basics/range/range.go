package main


import "fmt"

func main(){
	

	array :=[3]int{3,6,9}

	//Iterate for loop using range 
	for k, v := range array{

		fmt.Println(k,v)

	}

	m :=map[string]string{"fruits":"Apple", "Sweets":"Laddu", "Game":"Cricket"}

	for k, v := range m {

		fmt.Println(k,v)

	}

}