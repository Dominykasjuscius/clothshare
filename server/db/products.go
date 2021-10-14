package db

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsWrapper struct {
	conn *mongo.Collection
	log  *logrus.Logger
}

type ProductSchema struct {
	ID            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	Description   string             `bson:"description"`
	Condition     string             `bson:"condition"`
	Size          string             `bson:"size"`
	Color         string             `bson:"color"`
	ViewCount     int64              `bson:"viewCount"`
	Brand         string             `bson:"brand"`
	Category      string             `bson:"category"`
	Location      string             `bson:"location"`
	ImageFilePath string             `bson:"imgPath"`
	CreatedAt     time.Time          `bson:"createdAt"`
	LastSeenAt    time.Time          `bson:"lastSeenAt"`
	UpdatedAt     time.Time          `bson:"updatedAt"`
	Price         float64            `bson:"price"`
	Tags          []string           `bson:"tags"`
	Author        primitive.ObjectID `bson:"author"`
}

func (w *ProductsWrapper) InsertOne(db ProductSchema) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := w.conn.InsertOne(ctx, db)
	return err
}

func (w *ProductsWrapper) Find(filter bson.M) ([]ProductSchema, error) {
	schemas := []ProductSchema{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cur, err := w.conn.Find(ctx, filter)
	if err != nil {
		return schemas, err
	}

	err = cur.All(ctx, &schemas)
	if err != nil {
		return schemas, err
	}

	return schemas, nil
}

func (w *ProductsWrapper) UpdateImgPath(id primitive.ObjectID, path string) error {
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
				"imgPath": path,
			},
		},
	}

	_, err := w.conn.Aggregate(ctx, pipeline)
	if err != nil {
		return err
	}

	return nil
}

func (w *ProductsWrapper) Delete(filter bson.M) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := w.conn.DeleteMany(ctx, filter)
	if err != nil {
		return 0, err
	}

	return res.DeletedCount, nil
}
