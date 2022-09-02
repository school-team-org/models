package models

import (
	"encoding/json"
	"time"

	"github.com/bankonly/goutils/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultField struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
	UpdatedAt primitive.DateTime `json:"updated_at" bson:"updated_at"`
}

func BindUpdate(m any) primitive.D {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(m)
	json.Unmarshal(inrec, &inInterface)

	inInterface["UpdatedAt"] = primitive.NewDateTimeFromTime(time.Now())
	result := utils.BindUpdate(m)
	return result
}

func BindCreate(m any) {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(m)
	json.Unmarshal(inrec, &inInterface)
	inInterface["ID"] = primitive.NewObjectID()
	inInterface["CreatedAt"] = primitive.NewDateTimeFromTime(time.Now())
	inInterface["UpdatedAt"] = primitive.NewDateTimeFromTime(time.Now())
	inInterface["IsActive"] = true
}
