package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	PROVINCE_MODEL_NAME = "provinces"
)

type ProvinceModel struct {
	DefaultField
	NameLA string `json:"name_la" bson:"name_la" validate:"required"`
	NameEn string `json:"name_en" bson:"name_en" validate:"required"`
	Code   string `json:"code" bson:"code" validate:"required"`
}

func GetProvinceModel(m ProvinceModel) ProvinceModel {
	m.ID = primitive.NewObjectID()
	m.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	m.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	return m
}
