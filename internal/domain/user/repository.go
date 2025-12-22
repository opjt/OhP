package user

import (
	"context"
	db "ohp/internal/infrastructure/db/postgresql"
)

type userRepository struct {
	queries *db.Queries
}
type UserRepository interface {
	UpsertUserByEmail(context.Context, string) (*User, error)
}

func NewUserRepository(queries *db.Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (r *userRepository) UpsertUserByEmail(ctx context.Context, email string) (*User, error) {

	user, err := r.queries.UpsertUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	entity := &User{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
	return entity, nil
}
