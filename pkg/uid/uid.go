package uid

import (
	"github.com/google/uuid"
)

// Uid : generate a unique id on each call
func Uid() string {
	return uuid.New().String()
}
