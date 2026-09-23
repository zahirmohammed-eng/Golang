package main

// script commands:  go get github.com/google/uuid

import {
	"fmt"
	"github.com/google/uuid"
}

func main(){
	id : uuid.New()
	fmt.Println(id)
}