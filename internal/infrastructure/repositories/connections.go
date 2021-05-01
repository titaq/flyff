package repositories

import (
	"os"

	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/titaq/flyff/internal/domain"
)

type connections struct {
	mapping map[string]*actor.PID
}

func (c *connections) Get(connId string) (*actor.PID, error) {
	pid, ok := c.mapping[connId]
	if !ok {
		return nil, os.ErrNotExist
	}

	return pid, nil
}

func (c *connections) Set(connId string, pid *actor.PID) error {
	c.mapping[connId] = pid
	return nil
}

func (c *connections) DelPid(pid *actor.PID) error {
	for k, v := range c.mapping {
		if v.Equal(pid) {
			delete(c.mapping, k)
			return nil
		}
	}
	return os.ErrNotExist
}

func (c *connections) DelConn(connId string) error {
	delete(c.mapping, connId)
	return nil
}

func NewConnections() domain.ConnectionsRepository {
	return &connections{
		mapping: make(map[string]*actor.PID),
	}
}
