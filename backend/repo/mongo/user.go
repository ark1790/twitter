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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type bsonUser struct {
	model.User `bson:"inline"`
	ID         primitive.ObjectID `bson:"_id"`
}

func (u *bsonUser) toModel() *model.User {
	u.User.ID = u.ID.Hex()
	return &u.User
}

type User struct {
	db   *mongo.Database
	name string
}

func (u *User) collection() *mongo.Collection {
	return u.db.Collection(u.name)
}

func NewUser(db *mongo.Database, table string) *User {
	return &User{db, table}
}

func (u *User) EnsureIndices(*model.User) error {
	log.Println("Starting EnsureIndices")
	_, err := u.collection().Indexes().CreateMany(context.Background(),
		[]mongo.IndexModel{
			{
				Keys:    bson.D{{Key: "username", Value: 1}},
				Options: options.Index().SetUnique(true)},
		})
	log.Println("Completed EnsureIndices", err)
	return err
}

func (u *User) Create(user *model.User) error {
	log.Println("Starting Create", user)
	now := time.Now().UTC()
	user.CreatedAt = now
	user.UpdatedAt = now

	result, err := u.collection().InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Completed Create", err)
		if err, ok := err.(mongo.WriteException); ok {
			for _, err := range err.WriteErrors {
				if err.Code == 11000 {
					return errors.ErrDuplicateKey
				}
			}
		}
		return err
	}
	user.ID = result.InsertedID.(primitive.ObjectID).Hex()
	log.Println("Completed Create")
	return nil
}

func (u *User) Fetch(username string) (*model.User, error) {
	log.Println("Starting Fetch", username)

	result := u.collection().FindOne(context.Background(), bson.M{"username": username})
	if err := result.Err(); err != nil {
		log.Println("Completed Fetch", err)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	var user bsonUser
	if err := result.Decode(&user); err != nil {
		log.Println("Completed Fetch", err)
		return nil, err
	}
	log.Println("Completed Fetch")
	return user.toModel(), nil
}
