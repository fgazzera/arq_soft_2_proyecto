package hotel

import (
	"context"
	"ficha_hotel-api/model"
	"ficha_hotel-api/utils/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetHotelById(id string) model.Hotel {
	var hotel model.Hotel
	db := db.MongoDb
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
		return hotel
	}
	err = db.Collection("ficha_hotel").FindOne(context.TODO(), bson.D{{"_id", objID}}).Decode(&hotel)
	if err != nil {
		fmt.Println(err)
		return hotel
	}
	return hotel

}

func InsertHotel(hotel model.Hotel) model.Hotel {
	db := db.MongoDb
	insertHotel := hotel
	insertHotel.ID = primitive.NewObjectID()
	_, err := db.Collection("ficha_hotel").InsertOne(context.TODO(), &insertHotel)

	if err != nil {
		fmt.Println(err)
		return hotel
	}
	hotel.ID = insertHotel.ID
	return hotel
}

func UpdateHotel(hotel model.Hotel) model.Hotel {
	db := db.MongoDb
	filter := bson.M{"_id": hotel.ID}
	update := bson.M{
		"$set": bson.M{
			"nombre":      hotel.Nombre,
			"descripcion": hotel.Descripcion,
			"email":       hotel.Email,
			"ciudad":      hotel.Ciudad,
			"images":      hotel.Images,
			"cant_hab":    hotel.CantHab,
			"amenities":   hotel.Amenities,
		},
	}

	_, err := db.Collection("ficha_hotel").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		fmt.Println(err)
		return hotel
	}

	return hotel
}
