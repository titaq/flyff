package interfaces

import (
	"context"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/sirupsen/logrus"
	"github.com/titaq/flyff/internal/application"
	"github.com/titaq/flyff/internal/domain"
	"github.com/titaq/flyff/pkg/flyff"
	"google.golang.org/protobuf/proto"
)

type relay struct {
	subscriber cloudevents.Client
	world      *application.World
}

type packet struct {
	header    byte
	checksum  uint32
	length    uint32
	checksum2 uint32
	cmd       uint32
}

func (r *relay) Receive(ctx context.Context) error {
	return r.subscriber.StartReceiver(ctx, func(e event.Event) {
		ie := new(domain.IncomingEvent)
		if err := proto.Unmarshal(e.Data(), ie); err != nil {
			logrus.Warnf("Cannot unmarshal incoming event : %s", err.Error())
			return
		}

		r.handleIncomingEvent(ctx, ie)
	})
}

func (r *relay) handleIncomingEvent(ctx context.Context, ie *domain.IncomingEvent) {
	switch ie.Type {
	case domain.IncomingEvent_CONNECT:
		r.world.OnConnect(ctx, ie.GetClientId())
	case domain.IncomingEvent_DISCONNECT:
		r.world.OnDisconnect(ctx, ie.GetClientId())
	case domain.IncomingEvent_NONE:
		// TODO : Increase security while reading packet (comparing length etc..)
		reader := flyff.InitializePacketReader(ie.GetData())
		packet := &packet{
			header:    reader.ReadByte(),
			checksum:  reader.ReadUInt32(),
			length:    reader.ReadUInt32(),
			checksum2: reader.ReadUInt32(),
		}

		_ = reader.ReadInt32() // -1 always
		packet.cmd = reader.ReadUInt32()

		switch packet.cmd {
		case 0xff00:
			r.world.OnMessage(ctx, &application.WorldMessage{
				ClientID: ie.GetClientId(),
				Packet:   flyff.NewFFPacketIn_Join(reader),
			})
		default:
			logrus.Infof("Received unknown packet cmd 0x%x of len %d", packet.cmd, packet.length)
		}
	default:
	}
}

func (r *relay) readPacket(data []byte) {

}

func NewRelay(subscriber cloudevents.Client, world *application.World) domain.RelayInterface {
	return &relay{
		subscriber: subscriber,
		world:      world,
	}
}
