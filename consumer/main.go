package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)
import "github.com/nats-io/nats.go"

func main() {
	fmt.Printf("Connecting")
	nc, err := nats.Connect("nats://localhost:4222", nats.Name("consumer"))
	if err != nil {
		log.Fatalf("😞 %s", err)
	}

	// Simple Async Subscriber
	fmt.Printf("Subscribing to channels.*")
	nc.Subscribe("channels.*", func(m *nats.Msg) {
		fmt.Printf("Recieved: %s\n", string(m.Data))
	})

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	// Connect to a server

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

}
