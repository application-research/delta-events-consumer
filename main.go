package main

import (
	"encoding/json"
	"github.com/application-research/delta-db/db_models"
	"github.com/application-research/delta-db/messaging"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Instantiate a new consumer that subscribes to the specified topic
	//consumer := messaging.NewDeltaMetricsMessageConsumer()
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(messaging.PrimaryTopic, "default_channel", config)
	if err != nil {
		log.Fatal(err)
	}

	// add a handler to consume messages
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		var content db_models.Content
		json.Unmarshal(message.Body, &content)
		log.Printf("Got a message: %v", content)
		log.Printf("Got a message: %v", content.ID)

		return nil
	}))

	err = consumer.ConnectToNSQD(messaging.MetricsTopicUrl)
	if err != nil {
		log.Fatal(err)
	}

	// wait for signal to exit
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan

	// disconnect
	consumer.Stop()

}
