package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	USER_MODEL_NAME = "users"
	POSITION_LIST   = []string{"head", "teacher", "intern"}
	ROLE_LIST       = []string{"student", "admin", "teacher", "parents"}
)

type UserModel struct {
	DefaultField   `bson:",inline"`
	Firstname      string             `json:"first_name" bson:"first_name" form:"first_name" validate:"required"`
	Lastname       string             `json:"last_name" bson:"last_name" form:"last_name" validate:"required"`
	PhoneNumber    string             `json:"phone_number" bson:"phone_number" form:"phone_number" validate:"required"`
	Avatar         string             `json:"avartar" bson:"avartar" form:"avartar"`
	Cover          string             `json:"cover" bson:"cover" form:"cover"`
	Email          string             `json:"email" bson:"email" form:"email" validate:"required"`
	IsActive       bool               `json:"is_active" bson:"is_active" form:"is_active"`
	BirthDate      string             `json:"birth_date" bson:"birth_date" form:"birth_date" validate:"required"`
	Role           string             `json:"role" bson:"role" form:"role" validate:"required"`
	Position       string             `json:"position" bson:"position" form:"position" validate:"required"`
	VillageId      primitive.ObjectID `json:"village_id" bson:"village_id" form:"village_id" validate:"required"`
	IdNo           string             `json:"id_no" bson:"id_no" form:"id_no" validate:"required"`
	DeviceToken    string             `json:"device_token" bson:"device_token" form:"device_token"`
	VillageBirthId primitive.ObjectID `json:"village_birth_id" bson:"village_birth_id" form:"village_birth_id" validate:"required"`
	Gender         string             `json:"gender" bson:"gender" form:"gender" validate:"required"`
	Height         float32            `json:"height" bson:"height" form:"height"`
	Weight         float32            `json:"weight" bson:"weight" form:"weight"`
	National       string             `json:"national" bson:"national" form:"national"`
	Region         string             `json:"region" bson:"region" form:"region"`
	Password       string             `json:"password" bson:"password" form:"password" validate:"required"`
}
