package interfaces

// Indexer is an interface to abstract the layer of selection/verification of new ID for calendar event
type Indexer interface {
	NewID() string                // Returns a new ID
	Parse(string) (string, error) // Verify the provided ID - does it complains to  specific format
}
