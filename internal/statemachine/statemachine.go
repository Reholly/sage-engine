package statemachine

import (
	"github.com/google/uuid"
	"sage/internal/model"
)

// StateMachine later it will be instantiated one per every configured transaction
type StateMachine struct {
	ID       uuid.UUID
	Pipeline model.Pipeline
}
