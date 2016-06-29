package command

import (
    "github.com/isuruceanu/gocqrs"
	//"github.com/twinj/uuid"
)

type CommandService struct {
    handlers gocqrs.CommandHandlerCollection
    commandStore gocqrs.CommandStore
    eventStore gocqrs.EventStore
    eventBus gocqrs.EventBus
}

func MakeCommandService(handlers gocqrs.CommandHandlerCollection,
    commandStore gocqrs.CommandStore,
    eventStore gocqrs.EventStore,
    eventBus gocqrs.EventBus) (CommandService) {
        return CommandService{handlers: handlers, commandStore: commandStore, 
            eventStore: eventStore, eventBus:eventBus}
    }

func (cs *CommandService) Execute(command gocqrs.Command) error  {
    
    //todo: check for an empty id
    if command.GetCorrelationId == nil {
        command.SetCorrelationId(command.GetId())
    }
    
    cs.commandStore.Save(command)
   
     
    context := gocqrs.MakeCommandContext(cs.eventStore, cs.eventBus, command.GetCorrelationId())
    
    handler, err := cs.handlers.GetFor(command)
    
    if err != nil {
        return err
    }
    
    return handler.Execute(command, context)
    
}