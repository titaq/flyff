package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/go-amqp"
	ceamqp "github.com/cloudevents/sdk-go/protocol/amqp/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"github.com/titaq/flyff/pkg/models"
	"github.com/titaq/relay/pkg/domain"
	"google.golang.org/protobuf/proto"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting titaq flyff-server")

	var cfg models.Configuration
	cleanenv.ReadEnv(&cfg)

	subscriberProtocol, err := ceamqp.NewProtocol(cfg.AmqpUri, cfg.IncomingTopic, []amqp.ConnOption{}, []amqp.SessionOption{})
	if err != nil {
		log.Fatalf("Failed to create amqp protocol: %v", err)
	}
	defer subscriberProtocol.Close(context.Background())

	publisherProtocol, err := ceamqp.NewProtocol(cfg.AmqpUri, cfg.OutgoingTopic, []amqp.ConnOption{}, []amqp.SessionOption{})
	if err != nil {
		log.Fatalf("Failed to create amqp protocol: %v", err)
	}
	defer publisherProtocol.Close(context.Background())

	subscriberClient, err := cloudevents.NewClient(subscriberProtocol)
	if err != nil {
		panic(err)
	}

	publisherClient, err := cloudevents.NewClient(publisherProtocol, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		panic(err)
	}

	subscriberClient.StartReceiver(context.Background(), func(e event.Event) {
		ie := new(domain.IncomingEvent)
		proto.Unmarshal(e.Data(), ie)
		fmt.Println(ie)

		oe := &domain.OutgoingEvent{
			ClientsList: []string{ie.GetClientId()},
			Data:        ([]byte)("Yo !"),
		}
		data, _ := proto.Marshal(oe)

		event := cloudevents.NewEvent()
		event.SetSource("/titaq/flyff")
		event.SetData(cloudevents.TextPlain, data)
		event.SetType("titaq.outgoingevent")

		res := publisherClient.Send(context.Background(), event)
		fmt.Println(res)
	})
}
