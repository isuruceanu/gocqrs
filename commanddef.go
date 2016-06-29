package gocqrs

// Command domain coomand that is sent to a Dispatcher 
type Command interface {
    Messager
    CommandType() string
}

// CommandHandler dispatches command
type CommandHandler interface {
    Execute(Command, CommandContext) error
}

// CommandService service for executing command
type CommandService interface {
    Execute(Command) error
}

// CommandStore store command
type CommandStore interface {
    Save(Command) error
}