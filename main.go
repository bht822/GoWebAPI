package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	/**
	* http is the standard GO library which has many tools built in 
	* HandleFunc register a call back function which is executed when "/"
	* end-point is called from the browser or application
	*/
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		start := time.Now();
		//Simply send back hello word
		fmt.Fprintf(w,"Hello From Go lang %d",time.Since(start));
		elapsed := time.Since(start);
		fmt.Printf("Hello From Go lang %d \n",elapsed);
	})

	/*
	* Best practice is to try catch the return statements from the ListenAndServer function 
	* We could pass the above callback in this function as well. This is Node.js equivalent 
	* of app.listen(port..) 
	*/
	_ = http.ListenAndServe(":8080",nil)
}