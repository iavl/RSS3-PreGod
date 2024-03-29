package model

import (
	"gorm.io/datatypes"
)

// `signature` model.
type Signature struct {
	Base BaseModel `gorm:"embedded"`

	FileURI string `gorm:"primaryKey;type:text;column:file_uri"`

	Signature string         `gorm:"type:text"`
	Agents    datatypes.JSON `gorm:"type:jsonb"`
}
