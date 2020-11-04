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

type bsonFollow struct {
	model.Follow `bson:"inline"`
	ID           primitive.ObjectID `bson:"_id"`
}

func (u *bsonFollow) toModel() *model.Follow {
	u.Follow.ID = u.ID.Hex()
	return &u.Follow
}

// Follow ...
type Follow struct {
	db   *mongo.Database
	name string
}

func (f *Follow) collection() *mongo.Collection {
	return f.db.Collection(f.name)
}

func NewFollow(db *mongo.Database, table string) *Follow {
	return &Follow{db, table}
}

func (f *Follow) EnsureIndices(*model.Follow) error {
	log.Println("Starting Follow EnsureIndices")
	_, err := f.collection().Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					{Key: "username", Value: 1},
					{Key: "profile", Value: 1},
				},
			},
			{
				Keys: bson.D{{Key: "profile", Value: 1}},
			},
		})
	log.Println("Completed Follow EnsureIndices", err)
	return err
}

func (f *Follow) Create(fl *model.Follow) error {
	log.Println("Starting Follow Create", fl)
	now := time.Now().UTC()
	fl.CreatedAt = now
	fl.UpdatedAt = now

	result, err := f.collection().InsertOne(context.Background(), fl)
	if err != nil {
		log.Println("Completed Follow Create", err)
		if err, ok := err.(mongo.WriteException); ok {
			for _, err := range err.WriteErrors {
				if err.Code == 11000 {
					return errors.ErrDuplicateKey
				}
			}
		}
		return err
	}
	fl.ID = result.InsertedID.(primitive.ObjectID).Hex()
	log.Println("Completed Follow Create")
	return nil
}

func (f *Follow) Delete(id string) error {
	log.Println("Starting Follow Delete", id)
	_id, err := primitive.ObjectIDFromHex(id)

	result, err := f.collection().DeleteOne(context.Background(),
		bson.M{
			"_id": _id,
		},
	)
	if err != nil {
		log.Println("Completed Follow Delete", err)
		return err
	}
	if result.DeletedCount == 0 {
		log.Println("Completed Follow Delete", result)
		return errors.ErrResourceNotFound
	}
	log.Println("Completed Follow Delete")
	return nil
}

func (f *Follow) Toggle(fl *model.Follow) error {
	log.Println("Starting Toggle", fl)
	result := f.collection().FindOne(context.Background(),
		bson.M{
			"username": fl.Username,
			"profile":  fl.Profile,
		},
	)
	errM := result.Err()
	if errM != nil && errM != mongo.ErrNoDocuments {
		return errM
	}

	if errM == mongo.ErrNoDocuments {
		return f.Create(fl)
	}

	var flw bsonFollow
	if err := result.Decode(&flw); err != nil {
		log.Println("Completed Fetch", err)
		return err
	}

	log.Println("Completed Fetch")

	return f.Delete(flw.ID.Hex())
}

func (t *Follow) List(uName string) ([]model.Follow, error) {
	log.Println("Starting Follow List", uName)
	cursor, err := t.collection().Find(context.Background(),
		bson.M{
			"username": uName,
		},
	)
	if err != nil {
		log.Println("Completed Follow List", err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	twts := []bsonFollow{}
	if err := cursor.All(context.Background(), &twts); err != nil {
		log.Println("Completed Follow List", err)
		return nil, err
	}

	out := make([]model.Follow, len(twts))
	for i, twt := range twts {
		out[i] = *twt.toModel()
	}
	log.Println("Completed Follow List")
	return out, nil
}
