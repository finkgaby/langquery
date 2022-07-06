package server

import "log"

func ConnectNatsAndSubscribe() {
	log.Printf("Connecting to nats")
	js := ConnectNats()
	log.Printf("Connected to nats")

	SubscribeToStream(js, Deserialize)
}
