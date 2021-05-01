package domain

import (
	"context"
)

type RelayInterface interface {
	Receive(ctx context.Context) error
}

type RelayPublisher interface {
	SendTo(ctx context.Context, data []byte, clients ...string) error
}
