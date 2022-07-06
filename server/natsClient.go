package server

import (
	"github.com/nats-io/nats.go"
	"langquery/commons"
	"log"
	"os"
	"runtime"
)

func ConnectNats() nats.JetStreamContext {
	nc, err := nats.Connect(os.Getenv(commons.NATS_URL))
	commons.CheckErr(err)
	js, err := nc.JetStream()
	commons.CheckErr(err)
	createStream(js)
	return js
}

func createStream(js nats.JetStreamContext) {
	err := js.DeleteStream(commons.STREAM_NAME)
	commons.CheckErr(err)
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     commons.STREAM_NAME,
		Subjects: []string{commons.STREAM_SUBJECTS},
	})
	commons.CheckErr(err)

}

func SubscribeToStream(js nats.JetStreamContext, f func([]byte) []byte) {
	log.Printf("Subscribing to " + commons.SUB_SUBJECT_NAME)
	js.Subscribe(commons.SUB_SUBJECT_NAME, func(msg *nats.Msg) {
		log.Printf("Msg recieved")
		msg.Ack()
		log.Printf("Subscriber fetched msg.Data:%s from subSubjectName:%q", string(msg.Data), msg.Subject)
		res := f(msg.Data)
		log.Printf("Data recieved: ")
		PublishToStream(js, res)
	}, nats.Durable("monitor"), nats.ManualAck())
	runtime.Goexit()
}

func PublishToStream(js nats.JetStreamContext, data []byte) {
	_, err := js.Publish(commons.PUB_SUBJECT_NAME, data)
	commons.CheckErr(err)
	log.Printf("Published queryJSON:%s to subjectName:%q", string(data), commons.PUB_SUBJECT_NAME)

}
