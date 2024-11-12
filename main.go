package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal("there was a error initializing the scanner")
	}
	line := scanner.Text()
	response := Decode_And_Fetch_Response(line)
	fmt.Println(response)
}
