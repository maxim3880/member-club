package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	Name      string
	Email     string
}
