package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quotes struct {
	ID             primitive.ObjectID `bson:"_id"`
	Quotes         string             `bson:"quotes"`
	Author         string             `bson:"author"`
	Category       string             `bson:"category"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	UsedAsQOTD     bool               `bson:"used_as_qotd"`
	UsedAsQotdDate string             `bson:"used_as_qotd_date"`
}

type QuotesList struct {
	ID       primitive.ObjectID `bson:"_id"`
	Quotes   string             `bson:"quotes"`
	Author   string             `bson:"author"`
	Category string             `bson:"category"`
}
