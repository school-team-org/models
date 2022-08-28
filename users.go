package models

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	USER_MODEL_NAME = "users"
)

type UserModel struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Firstname      string             `json:"first_name" bson:"first_name"`
	Lastname       string             `json:"last_name" bson:"last_name"`
	PhoneNumber    string             `json:"phone_number" bson:"phone_number"`
	Avatar         string             `json:"avartar" bson:"avartar"`
	Cover          string             `json:"cover" bson:"cover"`
	Email          string             `json:"email" bson:"email"`
	IsActive       bool               `json:"is_active" bson:"is_active"`
	BirthDate      primitive.DateTime `json:"birth_date" bson:"birth_date"`
	Role           string             `json:"role" bson:"role"`
	Position       string             `json:"position" bson:"position"`
	VillageId      primitive.ObjectID `json:"village_id" bson:"village_id"`
	IdNo           string             `json:"id_no" bson:"id_no"`
	DeviceToken    string             `json:"device_token" bson:"device_token"`
	VillageBirthId primitive.ObjectID `json:"village_birth_id" bson:"village_birth_id"`
	Gender         string             `json:"gender" bson:"gender"`
	Height         float32            `json:"height" bson:"height"`
	Weight         float32            `json:"weight" bson:"weight"`
	National       string             `json:"national" bson:"national"`
	Region         string             `json:"region" bson:"region"`
	Password       string             `json:"password" bson:"password"`
}
