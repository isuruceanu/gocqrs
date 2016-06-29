package gocqrs

import (
    "errors"
)

var ErrHandlerNotFound = errors.New("no handlers for command")

var ErrHandlerAlreadySet = errors.New("handler is already set")

type CommandHandlerCollection struct {
    handlers map[string]CommandHandler
}

func NewCommandHandlerCollection() *CommandHandlerCollection {
    b := &CommandHandlerCollection{
        handlers: make(map[string]CommandHandler),
    }
    
    return b
}

func (b *CommandHandlerCollection) GetFor(command Command) (CommandHandler, error)  {
    if handler, ok := b.handlers[command.CommandType()]; ok {
        return handler, nil
    }
    return nil, ErrHandlerNotFound
}

func (b *CommandHandlerCollection) Register(handler CommandHandler, command Command) error  {
    if _, ok := b.handlers[command.CommandType()]; ok {
        return ErrHandlerAlreadySet
    }
    
    b.handlers[command.CommandType()] = handler
    return nil
}