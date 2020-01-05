package utils

import (
	"math/rand"
)

// TODO: Don't forget to initialize randomizer

// NewID returns new (psevdo random ID). Will be used for Events & Users
func NewID() uint64 {
	return rand.Uint64()
}
