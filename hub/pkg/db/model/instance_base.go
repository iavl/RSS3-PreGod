package model

import (
	"github.com/NaturalSelectionLabs/RSS3-PreGod/shared/pkg/constants"
)

// `instance_base` model.
type InstanceBase struct {
	Base BaseModel `gorm:"embedded"`

	RSS3ID            string                   `gorm:"primaryKey;type:text;column:rss3_id"`
	ControllerAddress string                   `gorm:"type:text"`
	InstanceTypeID    constants.InstanceTypeID `gorm:"type:int;column:instance_type_id"`
}
