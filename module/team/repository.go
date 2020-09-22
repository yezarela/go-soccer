package team

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/yezarela/go-soccer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository represents repository of pkg team
type Repository interface {
	ListTeam(ctx context.Context) ([]model.Team, error)
	GetTeam(ctx context.Context, id string) (*model.Team, error)
	CreateTeam(ctx context.Context, data model.Team) (*model.Team, error)
}

type repository struct {
	db *mongo.Database
}

// NewRepository creates a new team repository
func NewRepository(db *mongo.Database) Repository {
	return &repository{db}
}

// ListTeam returns list of teams
func (repo *repository) ListTeam(ctx context.Context) ([]model.Team, error) {
	op := "team.Repository.ListTeam"

	var items []model.Team

	cur, err := repo.db.Collection("teams").Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var data model.Team

		err := cur.Decode(&data)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}

		items = append(items, data)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, op)
	}

	return items, nil
}

// GetTeam returns a team by id
func (repo *repository) GetTeam(ctx context.Context, id string) (*model.Team, error) {
	op := "team.Repository.GetTeam"

	oid, _ := primitive.ObjectIDFromHex(id)

	var data *model.Team

	err := repo.db.Collection("teams").FindOne(ctx, bson.M{"_id": oid}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, op)
	}

	return data, nil
}

// CreateTeam creates a new team
func (repo *repository) CreateTeam(ctx context.Context, data model.Team) (*model.Team, error) {
	op := "team.Repository.CreateTeam"

	body := bson.M{
		"name":        data.Name,
		"description": data.Description,
		"location":    data.Location,
		"players":     data.Players,
		"created_at":  time.Now(),
	}

	res, err := repo.db.Collection("teams").InsertOne(ctx, body)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return repo.GetTeam(ctx, oid.Hex())
	}

	return nil, nil
}
