package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserWrapper struct {
	conn *mongo.Collection
	log  *logrus.Logger
}

type UserSchema struct {
	ID         primitive.ObjectID  `bson:"_id"`
	Name       string              `bson:"name"`
	Password   string              `bson:"pass"`
	Email      string              `bson:"email"`
	Location   string              `bson:"location"`
	Bio        string              `bson:"bio"`
	Rating     int                 `bson:"rating"`
	PhotoPath  string              `bson:"photoPath"`
	CreatedAt  primitive.Timestamp `bson:"createdAt"`
	LastSeenAt primitive.Timestamp `bson:"lastSeenAt"`
	UpdatedAt  primitive.Timestamp `bson:"updatedAt"`

	Followers []UserSchema    `bson:"followers"`
	Products  []ProductSchema `bson:"products"`
}

func (w *UserWrapper) InsertOne(db UserSchema) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := w.conn.InsertOne(ctx, db)
	return err
}

func (w *UserWrapper) Find(filter bson.M) ([]UserSchema, error) {
	schemas := []UserSchema{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res := w.conn.FindOne(ctx, filter)
	if res.Err() != nil {
		return schemas, res.Err()
	}

	if err := res.Decode(&schemas); err != nil {
		return schemas, err
	}
	return schemas, nil
}

func (w *UserWrapper) UpdateImgPath(id primitive.ObjectID, path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"_id": id,
			},
		},
		bson.M{
			"$set": bson.M{
				"imgpath": path,
			},
		},
	}

	_, err := w.conn.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	return nil
}

func (w *UserWrapper) Delete(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := w.conn.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
