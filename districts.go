package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	DISTRICT_MODEL_NAME = "districts"
)

type DistrictModel struct {
	DefaultField `bson:",inline"`
	NameLA       string `json:"name_la" bson:"name_la" validate:"required"`
	NameEn       string `json:"name_en" bson:"name_en" validate:"required"`
}

func BindDistrictModel(m *DistrictModel) *DistrictModel {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return m
}

func BindUpdateDistrictModel(m *DistrictModel) *DistrictModel {
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return m
}
