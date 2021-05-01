package infrastructure

import (
	"context"
	"os"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/titaq/flyff/internal/domain"
	"google.golang.org/protobuf/proto"
)

type relay struct {
	publisher cloudevents.Client
}

func (r *relay) SendTo(ctx context.Context, data []byte, clients ...string) error {
	if data == nil || len(clients) == 0 {
		return os.ErrInvalid
	}

	oe := &domain.OutgoingEvent{
		ClientsList: clients,
		Data:        data,
	}

	protoData, _ := proto.Marshal(oe)

	event := cloudevents.NewEvent()
	event.SetSource("/titaq/flyff")
	event.SetData(cloudevents.TextPlain, protoData)
	event.SetType("titaq.outgoingevent")

	return r.publisher.Send(context.Background(), event)
}

func NewRelay(publisher cloudevents.Client) domain.RelayPublisher {
	return &relay{
		publisher: publisher,
	}
}
