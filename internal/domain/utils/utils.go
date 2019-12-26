package utils

import (
	"math/rand"
)

// TODO: Don't forget to initialize randomizer
// NewID returns new (psevdo random ID).
func NewID() uint64 {
	return rand.Uint64()
}
