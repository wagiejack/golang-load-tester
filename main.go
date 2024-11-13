package main

import (
	"flag"
	"fmt"
	"sync"
)

var successful_requests_counter int = 0
var failure_requests_counter int = 0

type Request_Metrics struct {
	time_for_request   uint64
	time_to_first_byte int
	time_to_last_byte  int
}

var metrics []Request_Metrics

func main() {
	url := flag.String("u", "", "The URL to request")
	method := flag.String("m", "GET", "HTTP method (GET, POST, etc.)")
	body := flag.String("bd", "", "Body to sent with the outgoing request")
	number_of_requests := flag.Int("n", 1, "Number of times to send the request")
	concurrent_requests := flag.Int("c", 1, "Number of Concurrent requests to be sent")
	flag.Parse()

	var wg sync.WaitGroup
	for i := 0; i < *concurrent_requests; i++ {
		wg.Add(1)
		go make_request(method, url, body, number_of_requests, &wg)
	}
	wg.Wait()

	fmt.Printf("successes: %v\n", successful_requests_counter)
	fmt.Printf("failures: %v\n", failure_requests_counter)
}

func make_request(method, url, body *string, number_of_requests *int, wg *sync.WaitGroup) {
	defer (*wg).Done()
	for *number_of_requests > 0 {
		_, statusCode, _ := Decode_And_Fetch_Response(*method, *url, *body)
		if statusCode >= 200 && statusCode < 300 {
			successful_requests_counter++
		} else {
			failure_requests_counter++
		}
		*number_of_requests--
	}
}
