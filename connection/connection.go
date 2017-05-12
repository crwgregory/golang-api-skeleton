package connection

type ConnectionInterface interface {
	Open() bool
	Close() bool
	GetDB() interface{}
}
