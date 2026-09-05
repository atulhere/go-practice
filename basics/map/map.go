package main 


import "fmt"

// When we need key-value storage and the keys are not restricted to
// sequential or numerical values, we can use a map.
// A map is similar to a hash map (or hash table).

// Example of Map


func main(){


	// Map Literals

	m :=map[string]string{"fruits":"Apple", "Sweets":"Laddu", "Game":"Cricket"}

	for k, v := range m {

		fmt.Println(k,v)

	}

	// Test that a key is present with a two-value assignment:


	element,ok :=m["fruit"]

	fmt.Println(element,ok)
	

}