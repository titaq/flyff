package application

import (
	"context"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/titaq/flyff/internal/domain"
	"github.com/titaq/flyff/pkg/flyff"
)

type WorldMessage struct {
	ClientID string
	Packet   proto.Message
}

type World struct {
	actorSystem *actor.ActorSystem
	rootCtx     *actor.RootContext

	connectionsRepo domain.ConnectionsRepository
	relayPublisher  domain.RelayPublisher
}

func (w *World) OnConnect(ctx context.Context, clientId string) {
	// TODO : spawn actor & set pid
	if err := w.connectionsRepo.Set(clientId, nil); err != nil {
		logrus.Warn(err)
	}
	logrus.Infof("%s connected", clientId)

	w.relayPublisher.SendTo(ctx, (&flyff.FPWelcome{
		ID: 1,
	}).Serialize(), clientId)
}

func (w *World) OnDisconnect(ctx context.Context, clientId string) {
	if err := w.connectionsRepo.DelConn(clientId); err != nil {
		logrus.Warn(err)
	}
	logrus.Infof("%s disconnected", clientId)
}

func (w *World) OnMessage(ctx context.Context, msg *WorldMessage) {
	pid, err := w.connectionsRepo.Get(msg.ClientID)
	if err != nil {
		return
	}

	if join, ok := msg.Packet.(*flyff.FFPacketIn_Join); ok && pid == nil {
		pid = w.rootCtx.Spawn(CharacterProducer(join.GetSlot()))
		if err = w.connectionsRepo.Set(msg.ClientID, pid); err != nil {
			logrus.
				WithField("client_id", msg.ClientID).
				WithField("error", err.Error()).
				Warnf("Cannot insert actor")
			return
		}
	}

	w.rootCtx.Send(pid, msg.Packet)
}

func NewWorld(connectionsRepo domain.ConnectionsRepository, relayPublisher domain.RelayPublisher) *World {
	world := &World{
		connectionsRepo: connectionsRepo,
		relayPublisher:  relayPublisher,
		actorSystem:     actor.NewActorSystem(),
	}
	world.rootCtx = world.actorSystem.Root
	return world
}
