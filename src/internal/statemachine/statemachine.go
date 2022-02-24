package statemachine

import (
	"coffee-maker-pro/internal/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var CoffeeMaker = New()

const STATEID = "6213d06043a793edf1471135"

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
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	State     StateType          `bson:"value" json:"value"`
	PrevState StateType          `bson:"prevValue" json:"prevValue"`
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
	log.Printf("Event received: %s", event)
	dbClient := database.Init()
	stateId, _ := primitive.ObjectIDFromHex(STATEID)
	switch event {
	case TurnOn:
		CoffeeMaker.SetState(Idle)
		database.Update(dbClient, database.STATUS, bson.M{"$set": bson.M{"value": Idle}}, bson.M{"_id": stateId})
		log.Printf("State changed to: %s", Idle)
		return Idle
	case TurnOff:
		CoffeeMaker.SetState(Off)
		log.Printf("State changed to: %s", Off)
		return Off
	case ApplyConfig:
		CoffeeMaker.SetState(Applying)
		log.Printf("State changed to: %s", Applying)
		return Applying
	case ChangeConfig:
		CoffeeMaker.SetState(Idle)
		log.Printf("State changed to: %s", Idle)
		return Idle
	case SetReady:
		CoffeeMaker.SetState(Ready)
		log.Printf("State changed to: %s", Ready)
		return Ready
	case StartBrewing:
		CoffeeMaker.SetState(Brewing)
		log.Printf("State changed to: %s", Brewing)
		return Brewing
	case StopBrewing:
		CoffeeMaker.SetState(Idle)
		log.Printf("State changed to: %s", Idle)
		return Idle
	default:
		log.Println("No valid event was sent")
		return CoffeeMaker.State
	}
}
