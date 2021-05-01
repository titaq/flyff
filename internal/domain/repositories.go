package domain

import "github.com/AsynkronIT/protoactor-go/actor"

type ConnectionsRepository interface {
	Get(connId string) (*actor.PID, error)
	Set(connId string, pid *actor.PID) error
	DelPid(pid *actor.PID) error
	DelConn(connId string) error
}
