package models

import (
	"time"

	"github.com/bankonly/goutils/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	VILLAGE_MODEL_NAME = "districts"
)

type VillageModel struct {
	DefaultField `bson:",inline"`
	NameLA       string `json:"name_la" bson:"name_la" validate:"required"`
	NameEn       string `json:"name_en" bson:"name_en" validate:"required"`
}

func BindVillageModel(m *VillageModel) *VillageModel {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.IsActive = true
	return m
}

func BindUpdateVillageModel(m VillageModel) primitive.D {
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	result := utils.BindUpdate(m)
	return result
}
