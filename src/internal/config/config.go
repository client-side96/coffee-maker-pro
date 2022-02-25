package config

import "go.mongodb.org/mongo-driver/bson/primitive"

type Config struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Temp      float64            `bson:"temp" json:"temp"`
	Pressure  float64            `bson:"pressure" json:"pressure"`
	Grinding  float64            `bson:"grinding" json:"grinding"`
	IsApplied bool               `bson:"isApplied" json:"isApplied"`
}
