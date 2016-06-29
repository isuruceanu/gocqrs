package gocqrs

import (
	"github.com/twinj/uuid"
)

// CommandContext context of Command
type CommandContext struct {
    eventStore EventStore
    eventBus EventBus
    correlationID uuid.UUID
}

func MakeCommandContext(eventStore EventStore, eventBus EventBus, correlationId uuid.UUID) (CommandContext)  {
    return CommandContext{eventStore: eventStore, eventBus: eventBus, correlationID: correlationId}
}

func (cc *CommandContext) BeginUnitOfWork() (UnitOfWorker, error) {
    //todo implement
    return nil, nil
}
