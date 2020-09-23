package player

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/yezarela/go-soccer/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository represents repository of pkg player
type Repository interface {
	ListPlayer(ctx context.Context) ([]model.Player, error)
	GetPlayer(ctx context.Context, id string) (*model.Player, error)
	CreatePlayer(ctx context.Context, data model.Player) (*model.Player, error)
}

type repository struct {
	db *mongo.Database
}

// NewRepository creates a new player repository
func NewRepository(db *mongo.Database) Repository {
	return &repository{db}
}

// ListPlayer returns list of players
func (repo *repository) ListPlayer(ctx context.Context) ([]model.Player, error) {
	op := "player.Repository.ListPlayer"

	cur, err := repo.db.Collection("players").Find(ctx, bson.D{})
	if err != nil {
		return nil, errors.Wrap(err, op)
	}
	defer cur.Close(ctx)

	var items []model.Player

	for cur.Next(ctx) {
		var data model.Player

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

// GetPlayer returns a player by id
func (repo *repository) GetPlayer(ctx context.Context, id string) (*model.Player, error) {
	op := "player.Repository.GetPlayer"

	oid, _ := primitive.ObjectIDFromHex(id)

	var data *model.Player

	err := repo.db.Collection("players").FindOne(ctx, bson.M{"_id": oid}).Decode(&data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, errors.Wrap(err, op)
	}

	return data, nil
}

// CreatePlayer creates a new player
func (repo *repository) CreatePlayer(ctx context.Context, data model.Player) (*model.Player, error) {
	op := "player.Repository.CreatePlayer"

	data.CreatedAt = time.Now()
	data.ID = primitive.NilObjectID

	res, err := repo.db.Collection("players").InsertOne(ctx, data)
	if err != nil {
		return nil, errors.Wrap(err, op)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		return repo.GetPlayer(ctx, oid.Hex())
	}

	return nil, nil
}
