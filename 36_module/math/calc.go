package math

import (
	"github.com/google/uuid"
)

func Add(x1 int, x2 int) int {
	return x1 + x2
}

func GetUuid() uuid.UUID {
	return uuid.New()
}