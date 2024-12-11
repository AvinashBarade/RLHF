package main

import (
	"fmt"
	"time"

	"github.com/tchajed/go-actor"
)

type ComputerActor struct {
	actor.Behavior
	id string
}

func (a *ComputerActor) New() actor.Behavior {
	return &ComputerActor{id: a.Self().String()}
}

func (a *ComputerActor) Receive(ctx actor.Context, msg interface{}) {
	switch msg := msg.(type) {
	case string:
		fmt.Printf("Computer %s received message: %s\n", a.id, msg)
		time.Sleep(time.Second) // Simulate processing time
		ctx.Tell(ctx.Sender(), "Processed")
	default:
		fmt.Printf("Computer %s received unknown message: %v\n", a.id, msg)
	}
}

func main() {
	system := actor.NewSystem("ComputerSystem")
	defer system.Stop()

	computer1 := system.Spawn("computer1", &ComputerActor{})
	computer2 := system.Spawn("computer2", &ComputerActor{})

	computer1.Tell(computer2, "Hello from Computer 1!")
	computer2.Tell(computer1, "Hello from Computer 2!")

	select {} // Block main goroutine to keep the system running
}
