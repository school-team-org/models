package models

var (
	VILLAGE_MODEL_NAME = "districts"
)

type VillageModel struct {
	DefaultField `bson:",inline"`
	NameLA       string `json:"name_la" bson:"name_la" validate:"required"`
	NameEn       string `json:"name_en" bson:"name_en" validate:"required"`
}
