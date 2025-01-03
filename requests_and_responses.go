package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// Expecting to extract the method and the url from the string that the user enters with all the validations that are done over here
func Decode_And_Fetch_Response(method, url, body string) (string, int, string) {
	switch strings.ToLower(method) {
	case "get":
		{
			return GET_Response(url), 200, ""
		}
	case "post":
		{
			return POST_Response(url, body), 200, ""
		}
	case "put":
		{
			return Create_Request_And_Send(method, url, body)
		}
	case "patch":
		{
			return Create_Request_And_Send(method, url, body)
		}
	case "delete":
		{
			return Create_Request_And_Send(method, url, body)
		}
	}
	return "The request method does not match any of the methods [GET,PUT,PATCH,POST,DELETE]", -1, ""
}

func GET_Response(url string) string {
	//Recording time of starting
	startTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("error getting the response as", err.Error())
	}
	//TTFB
	time_to_first_byte := time.Since(startTime)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error reading response body as", err.Error())
	}
	final_body := string(body)

	//TTLB
	time_to_last_byte := time.Since(startTime)

	//Storing the metrics
	localMetric := Request_Metrics{
		time_for_request:   uint64(time.Since(startTime)),
		time_to_first_byte: int(time_to_first_byte),
		time_to_last_byte:  int(time_to_last_byte),
		request_body:       final_body,
	}
	Metrics = append(Metrics, localMetric)
	return final_body
}

func POST_Response(url string, body string) string {

	//Recording time of starting
	startTime := time.Now()

	resp, err := http.Post(url, "application/json", bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatalln("Error sending the post response as ", err.Error())
	}

	//TTFB
	time_to_first_byte := time.Since(startTime)

	defer resp.Body.Close()
	final_body := string(body)

	//TTLB
	time_to_last_byte := time.Since(startTime)

	//Storing the metrics
	localMetric := Request_Metrics{
		time_for_request:   uint64(time.Since(startTime)),
		time_to_first_byte: int(time_to_first_byte),
		time_to_last_byte:  int(time_to_last_byte),
		request_body:       final_body,
	}
	Metrics = append(Metrics, localMetric)
	return final_body
}

func Create_Request_And_Send(method string, url string, body string) (string, int, string) {

	//Recording time of starting
	startTime := time.Now()

	//Creation of Request
	req, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatalln("Error sending the post response as ", err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	//TTFB
	time_to_first_byte := time.Since(startTime)

	//Create client send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending PUT request:", err)
	}
	defer resp.Body.Close()

	//Parsing of necessary information from response
	status := resp.Status
	statusCode := resp.StatusCode
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	responseBody := string(bodyBytes)

	//TTLB
	time_to_last_byte := time.Since(startTime)
	//Storing the metrics
	localMetric := Request_Metrics{
		time_for_request:   uint64(time.Since(startTime)),
		time_to_first_byte: int(time_to_first_byte),
		time_to_last_byte:  int(time_to_last_byte),
		request_body:       responseBody,
	}

	Metrics = append(Metrics, localMetric)
	return status, statusCode, responseBody
}
