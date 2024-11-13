package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("u", "", "The URL to request")
	method := flag.String("m", "GET", "HTTP method (GET, POST, etc.)")
	body := flag.String("bd", "", "Body to sent with the outgoing request")
	number_of_requests := flag.Int("n", 1, "Number of times to send the request")
	flag.Parse()
	for *number_of_requests > 0 {
		make_request(method, url, body)
		*number_of_requests--
	}
}

func make_request(method, url, body *string) {
	status, statusCode, Response := Decode_And_Fetch_Response(*method, *url, *body)
	fmt.Printf("Status:-%v \n", status)
	fmt.Printf("statusCode:-%v\n", statusCode)
	fmt.Printf("response:-%v\n", Response)
}
