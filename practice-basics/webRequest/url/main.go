package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println(myurl)
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme, result.Host, result.Path, result.Port(), result.RawQuery)
	qparams := result.Query()
	fmt.Println(qparams)
}
