package main

import (
	"fmt"
	"time"
)

type myMsg struct {
	seqNum int
	message string
}

func main() {
	fmt.Println("Go channels starting")
	
	ch := make(chan *myMsg) // unbuffered
	go chanSender(ch)
	
	for msg := range ch {
		fmt.Println("Message", msg.seqNum, ":", msg.message)
	}
}

func chanSender(out chan<- *myMsg) {
	seqNum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(1*time.Second)
		out <- &myMsg{seqNum, "moo"}
		seqNum++
	}
	close(out)
}