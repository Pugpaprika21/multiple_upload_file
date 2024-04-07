package models

import (
	"time"

	"gorm.io/gorm"
)

type UploadFiles struct {
	gorm.Model
	Filename  string
	Path      string
	Size      int64
	MimeType  string
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}
