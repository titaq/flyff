package domain

import "github.com/titaq/relay/pkg/domain"

type IncomingEvent = domain.IncomingEvent
type OutgoingEvent = domain.OutgoingEvent

const (
	IncomingEvent_NONE       = domain.IncomingEvent_NONE
	IncomingEvent_CONNECT    = domain.IncomingEvent_CONNECT
	IncomingEvent_DISCONNECT = domain.IncomingEvent_DISCONNECT
)
