package user

import (
	"context"

	"github.com/google/uuid"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) UpsertUserByEmail(ctx context.Context, email string) (*User, error) {
	return s.repo.UpsertUserByEmail(ctx, email)
}

func (s *UserService) FindByEmail(ctx context.Context, id uuid.UUID) (*User, error) {
	return s.repo.FindByID(ctx, id)
}
