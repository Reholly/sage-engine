package model

import (
	"github.com/google/uuid"
	"net/url"
)

// Entity describes service with unique ID and Address.
type Entity struct {
	ID      uuid.UUID
	Name    string
	Address *url.URL
}
