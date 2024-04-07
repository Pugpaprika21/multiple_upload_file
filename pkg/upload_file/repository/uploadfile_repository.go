package repository

import (
	"time"

	"github.com/Pugpaprika21/pkg/upload_file/dto"
	"github.com/Pugpaprika21/pkg/upload_file/models"
	"gorm.io/gorm"
)

type UploadFileRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *UploadFileRepository {
	return &UploadFileRepository{
		db: db,
	}
}

func (u *UploadFileRepository) Save(body *dto.UploadfileDTO) error {
	now := time.Now()
	file := &models.UploadFiles{
		Filename:  body.Filename,
		Path:      body.Path,
		Size:      body.Size,
		MimeType:  body.MimeType,
		Extension: body.Extension,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := u.db.Model(&models.UploadFiles{}).Create(file).Error; err != nil {
		return err
	}
	return nil
}

func (u *UploadFileRepository) GetByID(id string) *dto.UploadfileRecord {
	var row *dto.UploadfileRecord
	u.db.Model(&models.UploadFiles{}).Where("deleted_at IS NULL AND id = ?", id).First(&row)
	return row
}

func (u *UploadFileRepository) GetAll() []*dto.UploadfileRecord {
	var rows []*dto.UploadfileRecord
	u.db.Model(&models.UploadFiles{}).Where("deleted_at IS NULL").Order("created_at DESC").Find(&rows)
	return rows
}

func (u *UploadFileRepository) UpdateByID(id string, body *dto.UploadfileDTO) error {
	now := time.Now()
	file := &models.UploadFiles{
		Filename:  body.Filename,
		Path:      body.Path,
		Size:      body.Size,
		MimeType:  body.MimeType,
		Extension: body.Extension,
		UpdatedAt: now,
	}

	if err := u.db.Model(&models.UploadFiles{}).Where("id = ?", id).Updates(file).Error; err != nil {
		return err
	}
	return nil
}

func (u *UploadFileRepository) DeleteByID(id string) error {
	return u.db.Delete(&models.UploadFiles{}, id).Error
}
