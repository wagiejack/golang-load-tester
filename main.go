package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "The URL to request")
	method := flag.String("method", "GET", "HTTP method (GET, POST, etc.)")
	body := flag.String("body", "", "Body to sent with the outgoing request")
	flag.Parse()
	status, statusCode, Response := Decode_And_Fetch_Response(*method, *url, *body)
	fmt.Printf("Status:-%v \n", status)
	fmt.Printf("statusCode:-%v\n", statusCode)
	fmt.Printf("response:-%v\n", Response)
}
