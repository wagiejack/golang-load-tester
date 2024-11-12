package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

// Expecting to extract the method and the url from the string that the user enters with all the validations that are done over here
func Decode_And_Fetch_Response(str string) string {
	splitted_string := strings.Split(str, " ")
	//method := splitted_string[0]
	url := splitted_string[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("error getting the response as", err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("error reading response body as", err.Error())
	}
	final_body := string(body)
	return final_body
}
