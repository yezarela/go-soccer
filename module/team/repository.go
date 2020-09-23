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

	lookup := bson.D{{"$lookup", bson.D{{"from", "players"}, {"localField", "players"}, {"foreignField", "_id"}, {"as", "players"}}}}

	cur, err := repo.db.Collection("teams").Aggregate(ctx, mongo.Pipeline{lookup})
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	var items []model.Team

	if err = cur.All(ctx, &items); err != nil {
		return nil, errors.Wrap(err, op)
	}

	return items, nil
}

// GetTeam returns a team by id
func (repo *repository) GetTeam(ctx context.Context, id string) (*model.Team, error) {
	op := "team.Repository.GetTeam"

	oid, _ := primitive.ObjectIDFromHex(id)

	lookup1 := bson.D{{"$lookup", bson.D{{"from", "players"}, {"localField", "players"}, {"foreignField", "_id"}, {"as", "players"}}}}
	lookup2 := bson.D{{"$match", bson.D{{"_id", oid}}}}

	cur, err := repo.db.Collection("teams").Aggregate(ctx, mongo.Pipeline{lookup1, lookup2})
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	var items []model.Team

	if err = cur.All(ctx, &items); err != nil {
		return nil, errors.Wrap(err, op)
	}

	if len(items) > 0 {
		return &items[0], nil
	}

	return nil, nil
}

// CreateTeam creates a new team
func (repo *repository) CreateTeam(ctx context.Context, data model.Team) (*model.Team, error) {
	op := "team.Repository.CreateTeam"

	body := bson.M{
		"name":        data.Name,
		"description": data.Description,
		"location":    data.Location,
		"created_at":  time.Now(),
	}

	if len(data.Players) > 0 {
		players := []interface{}{}

		for _, p := range data.Players {
			players = append(players, bson.M{
				"name":       p.Name,
				"nickname":   p.Nickname,
				"position":   p.Position,
				"created_at": time.Now(),
			})
		}

		res, err := repo.db.Collection("players").InsertMany(ctx, players)
		if err != nil {
			return nil, errors.Wrap(err, op)
		}

		body["players"] = res.InsertedIDs
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
