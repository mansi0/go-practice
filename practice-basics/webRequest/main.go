package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)
const url = "https://lco.dev"

func main() {
	fmt.Println("LCO")
	responce, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("responce is of type %T",responce)

	defer responce.Body.Close()

	dataBytes, err := ioutil.ReadAll(responce.Body)

	if(err != nil ) {
		panic(err)
	} 
	content := string(dataBytes)
	fmt.Println(content)
}