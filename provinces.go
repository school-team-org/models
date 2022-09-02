package models

var (
	PROVINCE_MODEL_NAME = "provinces"
)

type ProvinceModel struct {
	DefaultField `bson:",inline"`
	NameLA       string `json:"name_la" bson:"name_la" validate:"required"`
	NameEn       string `json:"name_en" bson:"name_en" validate:"required"`
	Code         string `json:"code" bson:"code" validate:"required"`
}
