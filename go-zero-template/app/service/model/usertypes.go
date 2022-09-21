package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID    int64                `bson:"_id,omitempty" json:"id,omitempty"`
	Name  string               `bson:"name,omitempty" json:"name,omitempty"`
	Price primitive.Decimal128 `bson:"price,omitempty" json:"price,omitempty"`
	// TODO: Fill your own fields
	UpdateAt time.Time `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
