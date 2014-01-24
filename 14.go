package main

import (
	"fmt"
)

type responseMsg struct {
	hash int
}

type requestMsg struct {
	nonce int
	responseChan chan *responseMsg
}

func main() {
	fmt.Println("Go channels starting")
	
	requestChan := make(chan *requestMsg)	
	
	go chanWorker(requestChan)

	for i := 0; i < 5; i++ {
		request := &requestMsg{i, make(chan *responseMsg)}
		requestChan <- request
		response := <- request.responseChan
		fmt.Println("Got response: ", response.hash)
	}
}

func chanWorker(requestChan <-chan *requestMsg) {
	for req := range requestChan {
		respVal := req.nonce * 1234567; // Super-secure hash
		response := &responseMsg{respVal}
		req.responseChan <- response
	}
}