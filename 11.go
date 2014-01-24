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
	go chanSender(ch, "cs1")
	go chanSender(ch, "cs2")
	
	for msg := range ch {
		fmt.Println("Message", msg.seqNum, ":", msg.message)
	}
}

func chanSender(out chan<- *myMsg, prefix string) {
	seqNum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(1*time.Second)
		out <- &myMsg{seqNum, fmt.Sprintf("%s: %s", prefix, "moo")}
		seqNum++
	}
	close(out)
}