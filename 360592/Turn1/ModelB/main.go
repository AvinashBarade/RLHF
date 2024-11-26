package main

import "fmt"

type Message struct {
	Text string
}

var (
	sendCh1 = make(chan Message)
	recvCh1 = make(chan Message)
	sendCh2 = make(chan Message)
	recvCh2 = make(chan Message)
)

func startCommunication() {
	go func() {
		for {
			msg := <-sendCh1
			recvCh2 <- msg
		}
	}()

	go func() {
		for {
			msg := <-sendCh2
			recvCh1 <- msg
		}
	}()
}

func service1() {
	startCommunication()

	for {
		msg := Message{Text: "Hello from Service 1!"}
		sendCh1 <- msg
		fmt.Println("Sent message from Service 1:", msg.Text)

		msg = <-recvCh1
		fmt.Println("Received message from Service 2:", msg.Text)
	}
}

func service2() {
	startCommunication()

	for {
		msg := <-recvCh2
		fmt.Println("Received message from Service 1:", msg.Text)

		msg = Message{Text: "Hello from Service 2!"}
		sendCh2 <- msg
		fmt.Println("Sent message from Service 2:", msg.Text)
	}
}

func main() {
	go service1()
	go service2()

	select {}
}
