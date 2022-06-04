package main

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	req,err := http.NewRequest("GET","/",nil);

	if(err !=nil){
		t.Fatal(err)
	}
	if(req == nil){
		t.Fatal("No response")
	}

}