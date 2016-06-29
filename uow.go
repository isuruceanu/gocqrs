package gocqrs

type UnitOfWorker interface {
    Commit()
    Load()
}

