package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Kb struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Desc  string             `json:"desc,omitempty" validate:"required"`
	Title string             `json:"title,omitempty" validate:"required"`
	Stars int32              `json:"stars,omitempty"`
}
