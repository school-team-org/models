package models

import "go.mongodb.org/mongo-driver/bson/primitive"

var (
	DISTRICT_MODEL_NAME = "districts"
)

type DistrictModel struct {
	DefaultField `bson:",inline"`
	NameLA       string             `json:"name_la" bson:"name_la" validate:"required"`
	NameEn       string             `json:"name_en" bson:"name_en" validate:"required"`
	ProvinceID   primitive.ObjectID `json:"province_id" bson:"province_id" validate:"required"`
}
