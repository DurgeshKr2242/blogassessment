package helpers

import "github.com/google/uuid"

func ParseUUID(id string) *uuid.UUID {
	s, err := uuid.Parse(id)
	if err != nil {
		return nil
	}
	return &s
}
