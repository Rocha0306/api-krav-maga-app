package InterfaceAdapters

import (
	"github.com/google/uuid"
)

func GenerateId() uint32 {
	return uuid.New().ID()
}
