package links

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

func New(db *mongo.Client, timeout time.Duration) *Repository {
	return &Repository{db: db, timeout: timeout}
}

type Repository struct {
	db      *mongo.Client
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateReq) (database.Link, error) {
	var l database.Link
	// implement me
	return l, nil
}

func (r *Repository) FindByURL(ctx context.Context, u string) (database.Link, error) {
	var l database.Link
	// implement me
	return l, nil
}

func (r *Repository) FindByCriteria(ctx context.Context, criteria Criteria) ([]database.Link, error) {
	return nil, nil
}
