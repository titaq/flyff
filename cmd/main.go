package main

import (
	"context"
	"log"

	"github.com/Azure/go-amqp"
	ceamqp "github.com/cloudevents/sdk-go/protocol/amqp/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
	"github.com/titaq/flyff/internal/application"
	"github.com/titaq/flyff/internal/infrastructure"
	"github.com/titaq/flyff/internal/infrastructure/repositories"
	"github.com/titaq/flyff/internal/interfaces"
	"github.com/titaq/flyff/pkg/models"
)

func main() {
	rootCtx := context.Background()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting titaq flyff-server")

	var cfg models.Configuration
	cleanenv.ReadEnv(&cfg)

	/*
	* Messaging interface & infrastructure setup
	 */
	subscriberProtocol, err := ceamqp.NewProtocol(cfg.AmqpUri, cfg.IncomingTopic, []amqp.ConnOption{}, []amqp.SessionOption{})
	if err != nil {
		log.Fatalf("Failed to create amqp protocol: %v", err)
	}
	defer subscriberProtocol.Close(rootCtx)

	publisherProtocol, err := ceamqp.NewProtocol(cfg.AmqpUri, cfg.OutgoingTopic, []amqp.ConnOption{}, []amqp.SessionOption{})
	if err != nil {
		log.Fatalf("Failed to create amqp protocol: %v", err)
	}
	defer publisherProtocol.Close(rootCtx)

	subscriberClient, err := cloudevents.NewClient(subscriberProtocol)
	if err != nil {
		panic(err)
	}

	publisherClient, err := cloudevents.NewClient(publisherProtocol, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
	if err != nil {
		panic(err)
	}

	world := application.NewWorld(repositories.NewConnections(), infrastructure.NewRelay(publisherClient))
	relay := interfaces.NewRelay(subscriberClient, world)
	logrus.Panic(relay.Receive(rootCtx))
}
