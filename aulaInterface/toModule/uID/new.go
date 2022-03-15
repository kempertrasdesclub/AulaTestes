package uID

import (
	"github.com/google/uuid"
)

type UID struct {
}

func (e *UID) Get() (uID string) {
	id := uuid.New()
	return id.String()
}
