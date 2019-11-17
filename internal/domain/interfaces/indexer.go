package indexer

// Indexer interface to provide IDs for business logic
type Indexer interface {
	ID() string
}
