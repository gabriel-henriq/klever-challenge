package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title   string             `bson:"title,omitempty" json:"title"`
	Author  string             `bson:"author,omitempty" json:"author"`
	Likes   int64              `bson:"likes,omitempty" json:"likes"`
	Unlikes int64              `bson:"unlikes,omitempty" json:"unlikes"`
}
