package models

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	USER_MODEL_NAME = "users"
)

type UserModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
}
