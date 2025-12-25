package endpoint

import (
	"context"
	"ohp/internal/pkg/token"
)

type EndpointService struct {
	repo EndpointRepository
}

func NewEndpointService(
	repo EndpointRepository,
) *EndpointService {
	return &EndpointService{
		repo: repo,
	}
}

func (s *EndpointService) Add(ctx context.Context, serviceName string) error {

	userClaim, err := token.UserFromContext(ctx)
	if err != nil {
		return err
	}

	if err := s.repo.Add(ctx, insertEndpointParams{
		userID:      userClaim.UserID,
		serviceName: serviceName,
		endpoint:    "test",
	}); err != nil {
		return err
	}

	return nil
}
