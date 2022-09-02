package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	USER_MODEL_NAME = "admins"
	POSITION_LIST   = []string{"head", "teacher", "intern"}
	ROLE_LIST       = []string{"student", "admin", "teacher", "parents"}
)

type UserModel struct {
	DefaultField
	Firstname      string             `json:"first_name" bson:"first_name" validate:"required"`
	Lastname       string             `json:"last_name" bson:"last_name" validate:"required"`
	PhoneNumber    string             `json:"phone_number" bson:"phone_number" validate:"required"`
	Avatar         string             `json:"avartar" bson:"avartar"`
	Cover          string             `json:"cover" bson:"cover"`
	Email          string             `json:"email" bson:"email" validate:"required"`
	IsActive       bool               `json:"is_active" bson:"is_active"`
	BirthDate      primitive.DateTime `json:"birth_date" bson:"birth_date" validate:"required"`
	Role           string             `json:"role" bson:"role" validate:"required"`
	Position       string             `json:"position" bson:"position" validate:"required"`
	VillageId      primitive.ObjectID `json:"village_id" bson:"village_id" validate:"required"`
	IdNo           string             `json:"id_no" bson:"id_no" validate:"required"`
	DeviceToken    string             `json:"device_token" bson:"device_token"`
	VillageBirthId primitive.ObjectID `json:"village_birth_id" bson:"village_birth_id" validate:"required"`
	Gender         string             `json:"gender" bson:"gender" validate:"required"`
	Height         float32            `json:"height" bson:"height"`
	Weight         float32            `json:"weight" bson:"weight"`
	National       string             `json:"national" bson:"national"`
	Region         string             `json:"region" bson:"region"`
	Password       string             `json:"password" bson:"password" validate:"required"`
}

func GetUserModel(m *UserModel) *UserModel {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return m
}
