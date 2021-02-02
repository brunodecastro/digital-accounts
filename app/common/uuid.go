package common

import "github.com/google/uuid"

// NewUUID - generates a new uuid
func NewUUID() string {
	return uuid.New().String()
}
