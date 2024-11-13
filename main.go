package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
)

var successful_requests_counter int = 0
var failure_requests_counter int = 0

type Request_Metrics struct {
	time_for_request   uint64
	time_to_first_byte int
	time_to_last_byte  int
	request_body       string
}

var Metrics []Request_Metrics

func main() {
	url := flag.String("u", "", "The URL to request")
	method := flag.String("m", "GET", "HTTP method (GET, POST, etc.)")
	body := flag.String("bd", "", "Body to sent with the outgoing request")
	number_of_requests := flag.Int("n", 1, "Number of times to send the request")
	concurrent_requests := flag.Int("c", 1, "Number of Concurrent requests to be sent")
	flag.Parse()

	ResponseCode_Count := make(map[int]int)

	var wg sync.WaitGroup
	for i := 0; i < *concurrent_requests; i++ {
		wg.Add(1)
		go make_request(method, url, body, number_of_requests, &wg, &ResponseCode_Count)
	}
	wg.Wait()

	fmt.Printf("successful request: %v\n", successful_requests_counter)
	fmt.Printf("failure request: %v\n", failure_requests_counter)
	fmt.Println()
	fmt.Printf("Request count by Status Code:-\n")

	for k, v := range ResponseCode_Count {
		fmt.Printf(" %v:%v\n", k, v)
	}

	//Measuring Metrics
	var min_TTFB int = math.MaxInt
	var max_TTFB int = math.MinInt
	var avg_TTFB int = 0
	var min_TTLB int = math.MaxInt
	var max_TTLB int = math.MinInt
	var avg_TTLB int = 0

	var min_req_time uint64 = math.MaxUint64
	var max_req_time uint64 = 0
	var avg_req_time uint64 = 0

	for _, v := range Metrics {
		min_TTFB = min(&min_TTFB, &v.time_to_first_byte)
		max_TTFB = max(&max_TTFB, &v.time_to_first_byte)
		avg_TTFB += (v.time_to_first_byte)
		min_TTLB = min(&min_TTFB, &v.time_to_last_byte)
		max_TTLB = max(&max_TTFB, &v.time_to_last_byte)
		avg_TTLB += (v.time_to_last_byte)

		min_req_time = min_uint64(&min_req_time, &v.time_for_request)
		max_req_time = max_uint64(&max_req_time, &v.time_for_request)
		avg_req_time += v.time_for_request
	}

	fmt.Printf("Total Request Time (s) (Min, Max, Mean).....: %v, %v, %v\nTime to First Byte (s) (Min, Max, Mean).....: %v, %v, %v\nTime to Last Byte (s) (Min, Max, Mean)......:%v, %v, %v", ms_to_seconds_uint64(&min_req_time), ms_to_seconds_uint64(&max_req_time), ms_to_seconds_uint64(&avg_req_time)/float64(len(Metrics)), ms_to_seconds(&min_TTFB), ms_to_seconds(&max_TTFB), ms_to_seconds(&avg_TTFB)/float64(len(Metrics)), ms_to_seconds(&min_TTLB), ms_to_seconds(&max_TTLB), ms_to_seconds(&avg_TTLB)/float64(len(Metrics)))

}

func make_request(method, url, body *string, number_of_requests *int, wg *sync.WaitGroup, ResponseCode_Count *map[int]int) {
	defer (*wg).Done()
	for *number_of_requests > 0 {
		//Making the Request
		_, statusCode, _ := Decode_And_Fetch_Response(*method, *url, *body)

		//Checking if request is successful or not and recording metrics accordingly
		if statusCode >= 200 && statusCode < 300 {
			successful_requests_counter++
		} else {
			failure_requests_counter++
		}

		//Incrementing StatusCode count on map for metrics
		(*ResponseCode_Count)[statusCode]++
		*number_of_requests--
	}
}

func min(a *int, b *int) int {
	if *a <= *b {
		return *a
	}
	return *b
}

func max(a *int, b *int) int {
	if *a <= *b {
		return *b
	}
	return *a
}

func min_uint64(a *uint64, b *uint64) uint64 {
	if *a <= *b {
		return *a
	}
	return *b
}

func max_uint64(a *uint64, b *uint64) uint64 {
	if *a <= *b {
		return *b
	}
	return *a
}

func ms_to_seconds_uint64(v *uint64) float64 {
	return float64(*v) / 1000
}

func ms_to_seconds(v *int) float64 {
	return float64(*v) / 1000
}
