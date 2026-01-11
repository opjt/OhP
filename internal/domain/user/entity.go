package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Email       string
	TermsAgreed bool
	CreatedAt   time.Time
}
