package server

import (
	"coffee-maker-pro/internal/statemachine"
)

type RPCServer struct{}

func (r *RPCServer) ReceiveMessage(event statemachine.EventType, reply *statemachine.StateType) error {
	*reply = statemachine.TransitionState(event)
	return nil
}
