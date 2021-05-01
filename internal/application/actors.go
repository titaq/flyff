package application

import (
	"fmt"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/titaq/flyff/pkg/flyff"
)

func CharacterProducer(slot uint32) *actor.Props {
	return actor.PropsFromProducer(func() actor.Actor {
		return &Character{
			Slot: slot,
		}
	})
}

type Character struct {
	Slot uint32
}

func (c *Character) Receive(context actor.Context) {
	switch context.Message().(type) {
	case *flyff.FFPacketIn_Join:
		// TODO : Connect user & send character data etc..
	default:
		fmt.Println("not implemented")
	}
}
