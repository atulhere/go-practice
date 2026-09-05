package main


import "fmt"
import "net/http"



func handler(w http.ResponseWriter, r *http.Request){

   	w.Write([]byte("India is a great country!"))

	fmt.Println("Maa")
}


func main(){


http.HandleFunc("/",handler)

fmt.Println("Server Starting on port 8080")

error := http.ListenAndServe(":8080",nil)

if error!= nil{

	fmt.Println("Error While connecting the server")
}


}

