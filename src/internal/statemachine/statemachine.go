package statemachine

import "log"

var CoffeeMaker = New()

type StateType string
type EventType string

const (
	Off      StateType = "Off"
	Idle     StateType = "Idle"
	Applying StateType = "Applying"
	Ready    StateType = "Ready"
	Brewing  StateType = "Brewing"

	TurnOn       EventType = "TurnOn"
	TurnOff      EventType = "TurnOff"
	ApplyConfig  EventType = "ApplyConfig"
	ChangeConfig EventType = "ChangeConfig"
	SetReady     EventType = "SetReady"
	StartBrewing EventType = "StartBrewing"
	StopBrewing  EventType = "StopBrewing"
)

type StateMachine struct {
	State     StateType
	PrevState StateType
}

func New() StateMachine {
	return StateMachine{
		State:     Off,
		PrevState: Off,
	}
}

func (s *StateMachine) SetState(newState StateType) {
	s.PrevState = s.State
	s.State = newState
}

func TransitionState(event EventType) StateType {
	switch event {
	case TurnOn:
		CoffeeMaker.SetState(Idle)
		return Idle
	case TurnOff:
		CoffeeMaker.SetState(Off)
		return Off
	case ApplyConfig:
		CoffeeMaker.SetState(Applying)
		return Applying
	case ChangeConfig:
		CoffeeMaker.SetState(Idle)
		return Idle
	case SetReady:
		CoffeeMaker.SetState(Ready)
		return Ready
	case StartBrewing:
		CoffeeMaker.SetState(Brewing)
		return Brewing
	case StopBrewing:
		CoffeeMaker.SetState(Idle)
		return Idle
	default:
		log.Println("No valid event was sent")
		return CoffeeMaker.State
	}
}
