package model

import "github.com/google/uuid"

type Step struct {
	ID                 uuid.UUID
	Name               string
	Entity             *Entity
	BehaviourReactions map[Behaviour]Action
}
