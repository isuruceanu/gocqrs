package gocqrs

import (
    "github.com/twinj/uuid"
	"time"
)


type Messager interface {
    GetId() uuid.UUID
    
    GetCorrelationId() uuid.UUID
    SetCorrelationId(id uuid.UUID)
    
    Timestamp() time.Time
    
    
}