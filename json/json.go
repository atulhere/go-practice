package main


import "fmt"

import "encoding/json"


type Message struct{

	Name  string
	Body  string
	Time  int32

}

func main(){



	m := Message{"Atul", "All Well", 1234566789}

	fmt.Println(m)

	// Marshal helps to convert the struct into json 
	jsonBytes, err := json.Marshal(m)

  if err != nil{
  	fmt.Println("Error While Marshaling the json")

  }

  	object := string(jsonBytes)

    fmt.Println("Print Json Object ",string(object))

   jsonString := []byte (`{"Name": "Atul","Body":"India is great!", "Time": 123323323}`)

  // Now let's use UnMarshal to store value in n type of Message 

     var n  Message

     error := json.Unmarshal(jsonString, &n)

     if error !=nil{
 		
 		fmt.Println("Issue in Unmarshal", error)
     
     }
     	 fmt.Println("Print Struct Values", n)


    








}



