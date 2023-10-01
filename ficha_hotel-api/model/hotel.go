package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hotel struct {
	ID          primitive.ObjectID `bson:"_id"`
	Nombre      string             `bson:"nombre"`
	Descripcion string             `bson:"descripcion"`
	Email       string             `bson:"email"`
	Ciudad      string             `bson:"ciudad"`
	Images      []string           `bson:"images"`
	CantHab     int                `bson:"cant_hab"`
	Amenities   []string           `bson:"amenities"`
}
