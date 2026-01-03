package notifications

import (
	"context"
	db "ohp/internal/infrastructure/db/postgresql"
)

type notiRepository struct {
	queries *db.Queries
}

type NotiRepository interface {
	Create(context.Context, Noti) (Noti, error)
	UpdateStatus(context.Context, Noti) error
}

func NewNotiRepository(queries *db.Queries) NotiRepository {
	return &notiRepository{
		queries: queries,
	}
}

func (r *notiRepository) Create(ctx context.Context, noti Noti) (Noti, error) {
	createdRow, err := r.queries.CreateNotification(ctx, db.CreateNotificationParams{
		EndpointID: noti.EndpointID,
		Body:       noti.Body,
	})
	entity := Noti{
		ID:         createdRow.ID,
		EndpointID: createdRow.EndpointID,
		Body:       createdRow.Body,
	}
	return entity, err
}

func (r *notiRepository) UpdateStatus(ctx context.Context, noti Noti) error {
	status := string(noti.Status)
	err := r.queries.UpdateStatusNotification(ctx, db.UpdateStatusNotificationParams{
		ID:     noti.ID,
		Status: &status,
	})
	return err
}
