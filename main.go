package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Golang Application ");
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"About page");
}

func main() {
	/**
	* http is the standard GO library which has many tools built in 
	* HandleFunc register a call back function which is executed when "/"
	* end-point is called from the browser or application
	*/
	http.HandleFunc("/",home);
	http.HandleFunc("/about",about)
	/*
	* Best practice is to try catch the return statements from the ListenAndServer function 
	* We could pass the above callback in this function as well. This is Node.js equivalent 
	* of app.listen(port..) 
	*/
	fmt.Printf(fmt.Sprintf("Server Running on port%s",portNumber));
	_ = http.ListenAndServe(portNumber,nil)
}