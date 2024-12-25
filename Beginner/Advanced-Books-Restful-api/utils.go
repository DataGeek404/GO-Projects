package main

import (
	"math/rand"
	"time"
)

// GenerateRandomID - Generates a random ID using a local random generator
func GenerateRandomID() int {
	source := rand.NewSource(time.Now().UnixNano()) // Create a new random source
	r := rand.New(source)                          // Create a random generator with the source
	return r.Intn(100000)
}
