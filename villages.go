package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	VILLAGE_MODEL_NAME = "districts"
)

type VillageModel struct {
	DefaultField
	NameLA string `json:"name_la" bson:"name_la" validate:"required"`
	NameEn string `json:"name_en" bson:"name_en" validate:"required"`
}

func BindVillageModel(m *VillageModel) *VillageModel {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return m
}

func BindUpdateVillageModel(m *VillageModel) *VillageModel {
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return m
}
