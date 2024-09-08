package model

import "github.com/google/uuid"

type Transaction struct {
	ID       uuid.UUID
	Pipeline uuid.UUID
}
