package models

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	VILLAGE_MODEL_NAME = "districts"
)

type VillageModel struct {
	DefaultField `bson:",inline" json:"default_field"`
	NameLA       string             `json:"name_la" bson:"name_la" validate:"required"`
	NameEn       string             `json:"name_en" bson:"name_en" validate:"required"`
	DistrictID   primitive.ObjectID `json:"district_id" bson:"district_id" validate:"required"`
	Code         string             `json:"code" bson:"code"`
}
