package storage

// EventKeeper interface with the storage backend
type EventKeeper interface {
	SaveEvent()
	GetEvent()
}
