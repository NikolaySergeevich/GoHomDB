package links

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"gitlab.com/robotomize/gb-golang/homework/03-01-umanager/internal/database"
)

const collection = "links"

func New(db *mongo.Database, timeout time.Duration) *Repository {
	return &Repository{db: db, timeout: timeout}
}

type Repository struct {
	db      *mongo.Database
	timeout time.Duration
}

func (r *Repository) Create(ctx context.Context, req CreateReq) (database.Link, error) {

	l := database.Link{
		ID:        req.ID,
		Title:     req.Title,
		URL:       req.URL,
		Images:    req.Images,
		Tags:      req.Tags,
		UserID:    req.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if _, err := r.db.Collection(collection).InsertOne(ctx, l); err != nil {
		return l, fmt.Errorf("mongo InsertOne: %w", err)
	}

	return l, nil
}

func (r *Repository) FindByUserAndURL(ctx context.Context, link, userID string) (database.Link, error) {
	var l database.Link
	// implement me
	result := r.db.Collection(collection).FindOne(ctx, bson.M{"url": link, "userID": userID})
	if err := result.Err(); err != nil {
		return l, fmt.Errorf("mongo FindOne: %w", err)
	}

	if err := result.Decode(&l); err != nil {
		return l, fmt.Errorf("mongo Decode: %w", err)
	}
	return l, nil
}

func (r *Repository) FindByCriteria(ctx context.Context, criteria Criteria) ([]database.Link, error) {
	res := make([]database.Link, 0)
	// result, err := r.db.Collection(collection).Find(ctx, bson.D{{Key: "tags", Value: criteria.Tags}})
	if len(criteria.Tags) == 0 && criteria.UserID == nil {
		return nil, nil
	}
	if len(criteria.Tags) == 1 && criteria.UserID == nil {
		req, err := r.db.Collection(collection).Find(ctx, bson.M{"tags": criteria.Tags[0]})
		if err != nil {
			return res, fmt.Errorf("mongo Find: %w", err)
		}
		if err = req.All(context.TODO(), &res); err != nil {
			return res, fmt.Errorf("mongo All Decod: %w", err)
		}
	}

	if len(criteria.Tags) > 1 && criteria.UserID == nil {
		result, err := r.db.Collection(collection).Find(ctx, bson.M{"tags": criteria.Tags})
		if err != nil {
			return res, fmt.Errorf("mongo Find: %w", err)
		}
		if err = result.All(context.TODO(), &res); err != nil {
			return res, fmt.Errorf("mongo All Decod: %w", err)
		}
	}

	if len(criteria.Tags) == 0 && criteria.UserID != nil {
		result, err := r.db.Collection(collection).Find(ctx, bson.M{"userID": criteria.UserID})
		if err != nil {
			return res, fmt.Errorf("mongo Find: %w", err)
		}
		if err = result.All(context.TODO(), &res); err != nil {
			return res, fmt.Errorf("mongo All Decod: %w", err)
		}
	}

	if len(criteria.Tags) == 1 && criteria.UserID != nil {
		result, err := r.db.Collection(collection).Find(ctx, bson.M{"tags": criteria.Tags[0], "userID": criteria.UserID})
		if err != nil {
			return res, fmt.Errorf("mongo Find: %w", err)
		}
		if err = result.All(context.TODO(), &res); err != nil {
			return res, fmt.Errorf("mongo All Decod: %w", err)
		}
	}

	if len(criteria.Tags) > 1 && criteria.UserID != nil {
		result, err := r.db.Collection(collection).Find(ctx, bson.M{"tags": criteria.Tags, "userID": criteria.UserID})
		if err != nil {
			return res, fmt.Errorf("mongo Find: %w", err)
		}
		if err = result.All(context.TODO(), &res); err != nil {
			return res, fmt.Errorf("mongo All Decod: %w", err)
		}
	}

	return res, nil
}
