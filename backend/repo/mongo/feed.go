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

type bsonFeed struct {
	model.Feed `bson:"inline"`
	ID         primitive.ObjectID `bson:"_id"`
}

func (t *bsonFeed) toModel() *model.Feed {
	t.Feed.ID = t.ID.Hex()
	return &t.Feed
}

type Feed struct {
	db   *mongo.Database
	name string
}

func (t *Feed) collection() *mongo.Collection {
	return t.db.Collection(t.name)
}

func NewFeed(db *mongo.Database, table string) *Feed {
	return &Feed{db, table}
}

func (t *Feed) EnsureIndices(*model.Feed) error {
	log.Println("Starting Feed EnsureIndices")
	_, err := t.collection().Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{{Key: "username", Value: 1}},
			},
		})
	log.Println("Completed Feed EnsureIndices", err)
	return err
}

func (t *Feed) Create(twt *model.Feed) error {
	log.Println("Starting Feed Create", twt)
	now := time.Now().UTC()
	twt.CreatedAt = now
	twt.UpdatedAt = now

	result, err := t.collection().InsertOne(context.Background(), twt)
	if err != nil {
		log.Println("Completed Feed Create", err)
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
	log.Println("Completed Feed Create")
	return nil
}

func (t *Feed) List(uName string, tpe string) ([]model.Feed, error) {
	log.Println("Starting Feed List", uName)
	qry := bson.M{
		"createdAt": bson.M{
			"$gt": time.Now().AddDate(0, 0, -1),
		},
	}

	if tpe == "home" {
		qry["for"] = uName
	} else {
		qry["username"] = uName
	}

	cursor, err := t.collection().Find(context.Background(), qry)

	if err != nil {
		log.Println("Completed Feed List", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	twts := []bsonFeed{}
	if err := cursor.All(context.Background(), &twts); err != nil {
		log.Println("Completed Feed List", err)
		return nil, err
	}

	out := make([]model.Feed, len(twts))
	for i, twt := range twts {
		out[i] = *twt.toModel()
	}
	log.Println("Completed Feed List")
	return out, nil
}
