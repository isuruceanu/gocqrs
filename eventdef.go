package gocqrs

import (
	"github.com/twinj/uuid"
)


// Event interface for defining a event
type Event interface {
    Messager
    
    GetSourceId() uuid.UUID
    SetSourceId(uuid.UUID)
    
    GetRootId() uuid.UUID
    SetRootId() uuid.UUID
}

// Events collection of Event
type Events []Event

// EventStore interfor for storing events
type EventStore interface {
    SaveEvent(Events) error
    Load(uuid.UUID) (Events, error)
}

// EventBus is an interface for distributing events
type EventBus interface {
    PublishEvents(Events)  
} 

// EventHandler an interface that all handlers of events should implement
type EventHandler interface {
    HandleEvent(Event)
}

// EventQueue an interface that queue an event
type EventQueue interface {
    Enqueue(Event)
}


