package mongorepo

import (
	"context"
	"log"
	"time"

	"github.com/ark1790/alpha/errors"
	"github.com/ark1790/alpha/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type bsonTweet struct {
	model.Tweet `bson:"inline"`
	ID          primitive.ObjectID `bson:"_id"`
}

func (t *bsonTweet) toModel() *model.Tweet {
	t.Tweet.ID = t.ID.Hex()
	return &t.Tweet
}

type Tweet struct {
	db   *mongo.Database
	name string
}

func (t *Tweet) collection() *mongo.Collection {
	return t.db.Collection(t.name)
}

func NewTweet(db *mongo.Database, table string) *Tweet {
	return &Tweet{db, table}
}

func (t *Tweet) EnsureIndices(*model.Tweet) error {
	log.Println("Starting Tweet EnsureIndices")
	_, err := t.collection().Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{{Key: "username", Value: 1}},
			},
		})
	log.Println("Completed Tweet EnsureIndices", err)
	return err
}

func (t *Tweet) Create(twt *model.Tweet) error {
	log.Println("Starting Tweet Create", twt)
	now := time.Now().UTC()
	twt.CreatedAt = now
	twt.UpdatedAt = now

	result, err := t.collection().InsertOne(context.Background(), twt)
	if err != nil {
		log.Println("Completed Tweet Create", err)
		if err, ok := err.(mongo.WriteException); ok {
			for _, err := range err.WriteErrors {
				if err.Code == 11000 {
					return errors.ErrDuplicateKey
				}
			}
		}
		return err
	}
	twt.ID = result.InsertedID.(primitive.ObjectID).Hex()
	log.Println("Completed Tweet Create")
	return nil
}

func (t *Tweet) List(uName string) ([]model.Tweet, error) {
	log.Println("Starting Tweet List", uName)
	cursor, err := t.collection().Find(context.Background(),
		bson.M{
			"username": uName,
		},
	)
	if err != nil {
		log.Println("Completed Tweet List", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	twts := []bsonTweet{}
	if err := cursor.All(context.Background(), &twts); err != nil {
		log.Println("Completed Tweet List", err)
		return nil, err
	}

	out := make([]model.Tweet, len(twts))
	for i, twt := range twts {
		out[i] = *twt.toModel()
	}
	log.Println("Completed Tweet List")
	return out, nil
}
