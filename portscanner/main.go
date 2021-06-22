package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/sleepycat/eventserver/portscanner/scan"
	"log"
)

func main() {
	natsURL := flag.String("nats", "", "the url of the Nats instance to connect to.")
	publishTopic := flag.String("publish", "", "The Nats topic to publish results to.")
	// subscribeTopic := flag.String("subscribe", "", "The Nats topic that will broadcast domains to scan.")
	jsonOutput := flag.Bool("json", true, "Output scan results as JSON.")
	target := flag.String("target", "", "The domain or IP to scan.")

	flag.Parse()

	nc, err := nats.Connect(*natsURL, nats.Name("portscanner"))
	if err != nil {
		log.Fatalf("😞 %s", err)
	}

	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	if *target == "" {
		panic("You must provide a scan target.")
	}

	// if *subscribeTopic != "" {
	// 	c.Subscribe(*subscribeTopic, func(m *nats.Msg) {
	// 		fmt.Printf("Recieved: %s\n", string(m.Data))
	// 	})
	// }

	openPorts := scan.Scan(*target)

	if *jsonOutput == true {
		byteArray, _ := json.Marshal(openPorts)
		fmt.Println(string(byteArray))
	}

	if *natsURL != "" && *publishTopic != "" {
		c.Publish(*publishTopic, openPorts)
	}

}
