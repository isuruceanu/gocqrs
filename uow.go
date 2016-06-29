package gocqrs

type UnitOfWorker interface {
    Commit()
    Load()
}

func Test() string {
    
}